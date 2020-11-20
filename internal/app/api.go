package app

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/gofrs/uuid"
	protoV1 "github.com/golang/protobuf/proto"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/cmd"
	"github.com/wailsapp/wails/lib/logger"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

const (
	defaultStateKey          = "state_default"
	defaultWorkspaceKey      = "wksp_default"
	workspacePrefix          = "wksp_"
	metadataKeyPrefix        = "md_"
	reflectMetadataKeyPrefix = "rmd_"
	messageKeyPrefix         = "msg_"
)

type api struct {
	runtime          *wails.Runtime
	logger           *logger.CustomLogger
	client           *client
	store            *store
	protofiles       *protoregistry.Files
	streamReq        chan proto.Message
	cancelMonitoring context.CancelFunc
	cancelInFlight   context.CancelFunc
	mu               sync.Mutex
	inFlight         bool
	appData          string
	state            *workspaceState
}

type statsHandler struct {
	*api
}

type storeLogger struct {
	*logger.CustomLogger
}

func (s storeLogger) Warningf(message string, args ...interface{}) {
	s.Warnf(message, args...)
}

// WailsInit is the init fuction for the wails runtime
func (a *api) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime
	a.logger = runtime.Log.New("API")

	var err error

	a.store, err = newStore(a.appData, storeLogger{runtime.Log.New("DB")})
	if err != nil {
		return fmt.Errorf("app: failed to create database: %v", err)
	}
	a.state = a.getCurrentState()

	ready := "wails:ready"
	if wails.BuildMode == cmd.BuildModeBridge {
		ready = "wails:loaded"
	}

	a.runtime.Events.On(ready, a.wailsReady)

	return nil
}

func (a *api) wailsReady(data ...interface{}) {
	a.runtime.Events.Emit(eventInit, initData{semver, wails.BuildMode})

	opts, err := a.GetWorkspaceOptions()
	if err != nil {
		a.logger.Errorf("%v", err)
		return
	}
	hds, err := a.GetReflectMetadata(opts.Addr)
	if err != nil {
		a.logger.Errorf("%v", err)
		return
	}

	if err := a.Connect(opts, hds, false); err != nil {
		a.logger.Errorf("%v", err)
	}

	go a.checkForUpdate()
}

func (a *api) checkForUpdate() {
	r, err := checkForUpdate()
	if err != nil {
		if err == noUpdate {
			a.logger.Info(err.Error())
			return
		}
		a.logger.Warnf("failed to check for updates: %v", err)
		return
	}
	a.runtime.Events.Emit(eventUpdateAvailable, r)
}

// WailsShutdown is the shutdown function that is called when wails shuts down
func (a *api) WailsShutdown() {
	a.store.close()
	if a.cancelMonitoring != nil {
		a.cancelMonitoring()
	}
	if a.cancelInFlight != nil {
		a.cancelInFlight()
	}
	if a.client != nil {
		a.client.close()
	}
}

func (a *api) emitError(title, msg string) {
	a.runtime.Events.Emit(eventError, errorMsg{title, msg})
}

func (a *api) getCurrentState() *workspaceState {
	rtn := &workspaceState{
		CurrentID: defaultWorkspaceKey,
	}
	val, err := a.store.get([]byte(defaultStateKey))
	if err != nil && err != errKeyNotFound {
		a.logger.Errorf("failed to get current state from store: %v", err)
	}
	if len(val) == 0 {
		return rtn
	}
	dec := gob.NewDecoder(bytes.NewBuffer(val))
	if err := dec.Decode(rtn); err != nil {
		a.logger.Errorf("failed to decode state: %v", err)
	}
	return rtn
}

// GetWorkspaceOptions gets the workspace options from the store
func (a *api) GetWorkspaceOptions() (*options, error) {
	val, err := a.store.get([]byte(a.state.CurrentID))
	if err != nil {
		return nil, err
	}

	var opts *options
	dec := gob.NewDecoder(bytes.NewBuffer(val))
	err = dec.Decode(&opts)

	if opts.ID == "" {
		opts.ID = defaultWorkspaceKey
	}

	return opts, err
}

// GetReflectMetadata gets the reflection metadata from the store by addr
func (a *api) GetReflectMetadata(addr string) (headers, error) {
	val, err := a.store.get([]byte(reflectMetadataKeyPrefix + hash(addr)))
	if err != nil {
		return nil, err
	}
	var hds headers
	dec := gob.NewDecoder(bytes.NewBuffer(val))
	err = dec.Decode(&hds)

	return hds, err
}

// GetMetadata gets the metadata from the store by addr
func (a *api) GetMetadata(addr string) (headers, error) {
	val, err := a.store.get([]byte(metadataKeyPrefix + hash(addr)))
	if err != nil {
		return nil, err
	}
	var hds headers
	dec := gob.NewDecoder(bytes.NewBuffer(val))
	err = dec.Decode(&hds)

	return hds, err
}

// ListWorkspaces returns a list of workspaces as their options
func (a *api) ListWorkspaces() ([]options, error) {
	items, err := a.store.list([]byte(workspacePrefix))
	if err != nil {
		return nil, err
	}
	var opts []options
	for _, val := range items {
		opt := options{}
		dec := gob.NewDecoder(bytes.NewBuffer(val))
		if err := dec.Decode(&opt); err != nil {
			return opts, err
		}
		// opts.ID was added in v0.3.0 so need to double check
		if opt.ID == defaultWorkspaceKey || opt.ID == "" {
			opt.ID = defaultWorkspaceKey
			opts = append([]options{opt}, opts...)
			continue
		}
		opts = append(opts, opt)
	}
	return opts, nil
}

// SelectWorkspace changes the current workspace by ID
func (a *api) SelectWorkspace(id string) (rerr error) {
	if a.state.CurrentID == id {
		return nil
	}

	defer func() {
		if rerr != nil {
			a.logger.Errorf(rerr.Error())
			a.emitError("Workspace Error", rerr.Error())

		}
	}()

	a.changeWorkspace(id)
	opts, err := a.GetWorkspaceOptions()
	if err != nil {
		return err
	}

	hds, err := a.GetReflectMetadata(opts.Addr)
	if err != nil {
		a.logger.Warnf("failed to get reflection metadata: %v", err)
	}

	// Ignoring error as Connect will already emit errors to the frontend
	a.Connect(opts, hds, false)

	return nil
}

// DeleteWorkspace will remove a workspace from the store and switch to
// the default workspace, if the deleted workspace is current.
func (a *api) DeleteWorkspace(id string) error {
	a.store.del([]byte(id))
	if a.state.CurrentID == id {
		a.SelectWorkspace(defaultWorkspaceKey)
	}
	// TODO: should we inform the user of deletion?
	return nil
}

// GetRawMessageState gets the message state by method full name
func (a *api) GetRawMessageState(method string) (string, error) {
	opts, err := a.GetWorkspaceOptions()
	if err != nil {
		return "", fmt.Errorf("failed to get message state, no workspace options: %v", err)
	}

	val, err := a.store.get([]byte(messageKeyPrefix + hash(opts.Addr, method)))
	return string(val), err
}

//FindProtoFiles opens a directory dialog to search for proto files
func (a *api) FindProtoFiles() (files []string, rerr error) {
	defer func() {
		if rerr != nil {
			const errTitle = "Not found"
			a.logger.Errorf(rerr.Error())
			a.emitError(errTitle, rerr.Error())
		}
	}()

	dir := a.SelectDirectory()

	// TODO(rogchap): we need to add a circuit breaker to not walk the whole file system
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".proto" {
			files = append(files, path)
		}
		return nil
	})

	if len(files) == 0 {
		return nil, errors.New("no *.proto files found")
	}

	return files, nil
}

//SelectDirectory opens a directory dialog and returns the path of the selected directory
func (a *api) SelectDirectory() string {
	if wails.BuildMode == cmd.BuildModeBridge {
		f, _ := filepath.Abs(filepath.Join(".", "internal", "server"))
		return f
	}
	return a.runtime.Dialog.SelectDirectory()
}

// Connect will attempt to connect a grpc server and parse any proto files
func (a *api) Connect(data, rawHeaders interface{}, save bool) (rerr error) {
	defer func() {
		if rerr != nil {
			const errTitle = "Connection error"
			a.logger.Errorf(rerr.Error())
			a.runtime.Events.Emit(eventClientStateChanged, connectivity.Shutdown.String())
			a.emitError(errTitle, rerr.Error())
		}
	}()

	var opts options
	if err := mapstructure.Decode(data, &opts); err != nil {
		return err
	}

	// reset all things
	a.runtime.Events.Emit(eventClientConnectStarted, opts.Addr)
	a.runtime.Events.Emit(eventServicesSelectChanged)
	a.runtime.Events.Emit(eventMethodInputChanged)

	if a.client != nil {
		if err := a.client.close(); err != nil {
			return fmt.Errorf("failed to close previous connection: %v", err)
		}
	}
	a.client = &client{}

	if a.cancelMonitoring != nil {
		a.cancelMonitoring()
	}
	ctx := context.Background()
	ctx, a.cancelMonitoring = context.WithCancel(ctx)
	go a.monitorStateChanges(ctx)

	var hds headers
	if err := mapstructure.Decode(rawHeaders, &hds); err != nil {
		a.logger.Errorf("unable to decode reflection metadata headers: %v", err)
	}

	if err := a.client.connect(opts, statsHandler{a}); err != nil {
		// Still try to parse proto definitions. Will fail silently
		// if using reflection services as there is no connection
		// to a valid server.
		a.cancelMonitoring()
		a.client = nil
		go a.loadProtoFiles(opts, hds, true)

		return fmt.Errorf("failed to connect to server: %v", err)
	}

	a.runtime.Events.Emit(eventClientConnected, opts.Addr)

	go a.loadProtoFiles(opts, hds, false)

	if !save {
		return nil
	}

	if opts.ID == "" {
		id := uuid.Must(uuid.NewV4())
		opts.ID = workspacePrefix + id.String()
		a.changeWorkspace(opts.ID)
	}

	go a.setWorkspaceOptions(opts)
	go a.setMetadata(reflectMetadataKeyPrefix+hash(opts.Addr), hds)

	return nil
}

func (a *api) changeWorkspace(id string) {
	a.state.CurrentID = id
	var val bytes.Buffer
	enc := gob.NewEncoder(&val)
	enc.Encode(a.state)

	a.store.set([]byte(defaultStateKey), val.Bytes())
}

func (a *api) loadProtoFiles(opts options, reflectHeaders headers, silent bool) (rerr error) {
	defer func() {
		if rerr != nil {
			const errTitle = "Failed to load RPC schema"
			a.logger.Errorf(rerr.Error())
			if !silent {
				a.emitError(errTitle, rerr.Error())
			}
		}
	}()

	a.protofiles = nil

	var err error
	if opts.Reflect {
		if a.client == nil {
			return errors.New("unable to load proto files via reflection: client is <nil>")
		}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(nil))
		for _, h := range reflectHeaders {
			if h.Key == "" {
				continue
			}
			ctx = metadata.AppendToOutgoingContext(ctx, h.Key, h.Val)
			fmt.Printf("h.Val = %+v\n", h.Val)
		}

		ctx = context.WithValue(ctx, ctxInternalKey{}, struct{}{})
		if a.protofiles, err = protoFilesFromReflectionAPI(ctx, a.client.conn); err != nil {
			return fmt.Errorf("error getting proto files from reflection API: %v", err)
		}
	}
	if !opts.Reflect && len(opts.Protos.Files) > 0 {
		if a.protofiles, err = protoFilesFromDisk(opts.Protos.Roots, opts.Protos.Files); err != nil {
			return fmt.Errorf("error parsing proto files from disk: %v", err)
		}
	}

	a.emitServicesSelect()
	return nil
}

func (a *api) emitServicesSelect() {
	if a.protofiles == nil {
		return
	}

	var ss servicesSelect
	a.protofiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		sds := fd.Services()
		for i := 0; i < sds.Len(); i++ {
			var s serviceSelect
			sd := sds.Get(i)
			s.FullName = string(sd.FullName())

			mds := sd.Methods()
			for j := 0; j < mds.Len(); j++ {
				md := mds.Get(j)
				fname := fmt.Sprintf("/%s/%s", sd.FullName(), md.Name())
				s.Methods = append(s.Methods, methodSelect{
					Name:     string(md.Name()),
					FullName: fname,
				})
			}
			sort.SliceStable(s.Methods, func(i, j int) bool {
				return s.Methods[i].Name < s.Methods[j].Name
			})
			ss = append(ss, s)
		}
		return true
	})

	if len(ss) == 0 {
		return
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].FullName < ss[j].FullName
	})

	a.runtime.Events.Emit(eventServicesSelectChanged, ss)
}

func (a *api) setWorkspaceOptions(opts options) {
	if opts.ID == "" {
		opts.ID = defaultWorkspaceKey
	}

	var val bytes.Buffer
	enc := gob.NewEncoder(&val)
	enc.Encode(opts)
	a.store.set([]byte(opts.ID), val.Bytes())
}

func (a *api) setMetadata(key string, hds headers) {
	var toSet headers
	for _, h := range hds {
		if h.Key == "" {
			continue
		}
		toSet = append(toSet, h)
	}
	var val bytes.Buffer
	enc := gob.NewEncoder(&val)
	enc.Encode(toSet)
	a.store.set([]byte(key), val.Bytes())
}

func (a *api) setMessage(method string, rawJSON []byte) {
	opts, err := a.GetWorkspaceOptions()
	if err != nil {
		a.logger.Errorf("failed to set message, no workspace options: %v", err)
		return
	}

	a.store.set([]byte(messageKeyPrefix+hash(opts.Addr, method)), rawJSON)
}

func (a *api) monitorStateChanges(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			// this will panic if we are waiting for a state change and the client (and it's connection)
			// get GC'd without this context being canceled
			a.logger.Errorf("panic monitoring state changes: %v", r)
		}
	}()
	for {
		if a.client == nil || a.client.conn == nil {
			continue
		}
		state := a.client.conn.GetState()
		a.runtime.Events.Emit(eventClientStateChanged, state.String())
		if ok := a.client.conn.WaitForStateChange(ctx, state); !ok {
			a.logger.Debug("ending monitoring of state changes")
			return
		}
	}
}

func (a *api) getMethodDesc(fullname string) (protoreflect.MethodDescriptor, error) {
	name := strings.Replace(fullname[1:], "/", ".", 1)
	desc, err := a.protofiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, fmt.Errorf("app: failed to find descriptor: %v", err)
	}

	methodDesc, ok := desc.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, fmt.Errorf("app: descriptor was not a method: %T", desc)
	}

	return methodDesc, nil
}

// SelectMethod is called when the user selects a new method by the given name
func (a *api) SelectMethod(fullname string) (rerr error) {
	defer func() {
		if rerr != nil {
			const errTitle = "Failed to select method"
			a.logger.Errorf(rerr.Error())
			a.emitError(errTitle, rerr.Error())
		}
	}()

	methodDesc, err := a.getMethodDesc(fullname)
	if err != nil {
		return err
	}

	in := messageViewFromDesc(methodDesc.Input())
	m := methodInput{
		FullName: fullname,
		Message:  in,
	}
	a.runtime.Events.Emit(eventMethodInputChanged, m)

	return nil
}

func messageViewFromDesc(md protoreflect.MessageDescriptor) *messageDesc {
	var rtn messageDesc
	rtn.Name = string(md.Name())
	rtn.FullName = string(md.FullName())

	fds := md.Fields()
	rtn.Fields = fieldViewsFromDesc(fds, false)

	return &rtn
}

func setFieldDescBasics(fdesc *fieldDesc, fd protoreflect.FieldDescriptor) {
	fdesc.Name = string(fd.Name())
	fdesc.Kind = fd.Kind().String()
	fdesc.FullName = string(fd.FullName())
	fdesc.Repeated = fd.IsList()

	if emd := fd.Enum(); emd != nil {
		evals := emd.Values()
		for i := 0; i < evals.Len(); i++ {
			eval := evals.Get(i)
			fdesc.Enum = append(fdesc.Enum, string(eval.Name()))
		}
	}
}

func fieldViewsFromDesc(fds protoreflect.FieldDescriptors, isOneof bool) []fieldDesc {
	var fields []fieldDesc

	seenOneof := make(map[protoreflect.Name]struct{})
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		fdesc := fieldDesc{}
		setFieldDescBasics(&fdesc, fd)

		if fd.IsMap() {
			fdesc.Kind = "map"
			fdesc.MapKey = &fieldDesc{}
			setFieldDescBasics(fdesc.MapKey, fd.MapKey())

			fdesc.MapValue = &fieldDesc{}
			mapVal := fd.MapValue()
			setFieldDescBasics(fdesc.MapValue, mapVal)
			if fmd := mapVal.Message(); fmd != nil {
				fdesc.MapValue.Message = messageViewFromDesc(fmd)
			}
			goto appendField
		}

		if !isOneof {
			if oneof := fd.ContainingOneof(); oneof != nil {
				if _, ok := seenOneof[oneof.Name()]; ok {
					continue
				}
				fdesc.Name = string(oneof.Name())
				fdesc.Kind = "oneof"
				fdesc.Oneof = fieldViewsFromDesc(oneof.Fields(), true)

				seenOneof[oneof.Name()] = struct{}{}
				goto appendField
			}
		}

		if fmd := fd.Message(); fmd != nil {
			fdesc.Message = messageViewFromDesc(fmd)
		}

	appendField:
		fields = append(fields, fdesc)
	}
	return fields
}

func (a *api) Send(method string, rawJSON []byte, rawHeaders interface{}) (rerr error) {
	defer func() {
		if rerr != nil {
			const errTitle = "Unable to send request"
			a.logger.Errorf(rerr.Error())
			a.emitError(errTitle, rerr.Error())
		}
	}()

	md, err := a.getMethodDesc(method)
	if err != nil {
		return err
	}

	fmt.Printf("string(rawJSON = %+v\n", string(rawJSON))
	req := dynamicpb.NewMessage(md.Input())
	if err := protojson.Unmarshal(rawJSON, req); err != nil {
		return err
	}

	go a.setMessage(method, rawJSON)

	if a.inFlight && md.IsStreamingClient() {
		a.streamReq <- req
		return nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	a.inFlight = true
	defer func() {
		a.inFlight = false
	}()

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(nil))

	var hs headers
	if err := mapstructure.Decode(rawHeaders, &hs); err != nil {
		return err
	}

	opts, err := a.GetWorkspaceOptions()
	if err != nil {
		return err
	}
	go a.setMetadata(metadataKeyPrefix+hash(opts.Addr), hs)

	for _, h := range hs {
		if h.Key == "" {
			continue
		}
		ctx = metadata.AppendToOutgoingContext(ctx, h.Key, h.Val)
	}

	ctx, a.cancelInFlight = context.WithCancel(ctx)

	a.runtime.Events.Emit(eventRPCStarted, rpcStart{
		ClientStream: md.IsStreamingClient(),
		ServerStream: md.IsStreamingServer(),
	})

	if md.IsStreamingClient() && md.IsStreamingServer() {
		stream, err := a.client.invokeBidiStream(ctx, method)
		if err != nil {
			return err
		}

		a.streamReq = make(chan proto.Message)
		go func() {
			for r := range a.streamReq {
				if err := stream.SendMsg(r); err != nil {
					close(a.streamReq)
					a.streamReq = nil
				}
			}
			stream.CloseSend()
		}()
		a.streamReq <- req

		for {
			resp := dynamicpb.NewMessage(md.Output())
			if err := stream.RecvMsg(resp); err != nil {
				break
			}
		}

		return nil
	}

	if md.IsStreamingClient() {
		stream, err := a.client.invokeClientStream(ctx, method)
		if err != nil {
			return err
		}
		a.streamReq = make(chan proto.Message, 1)
		a.streamReq <- req
		for r := range a.streamReq {
			if err := stream.SendMsg(r); err != nil {
				close(a.streamReq)
				a.streamReq = nil
			}
		}
		stream.CloseSend()
		resp := dynamicpb.NewMessage(md.Output())
		if err := stream.RecvMsg(resp); err != nil {
			return err
		}
		if err := stream.RecvMsg(nil); err != io.EOF {
			return err
		}

		return nil
	}

	if md.IsStreamingServer() {
		stream, err := a.client.invokeServerStream(ctx, method, req)
		if err != nil {
			return err
		}
		for {
			resp := dynamicpb.NewMessage(md.Output())
			if err := stream.RecvMsg(resp); err != nil {
				break
			}
		}

		return nil
	}

	resp := dynamicpb.NewMessage(md.Output())
	a.client.invoke(ctx, method, req, resp)
	return nil
}

// TagConn implements the stats.Handler interface
func (statsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	// noop
	return ctx
}

// HandleConn implements the stats.Handler interface
func (statsHandler) HandleConn(context.Context, stats.ConnStats) {
	// noop
}

// TagRPC implements the stats.Handler interface
func (statsHandler) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context {
	// noop
	return ctx
}

// HandleRPC implements the stats.Handler interface
func (a statsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	if internal := ctx.Value(ctxInternalKey{}); internal != nil {
		return
	}

	switch s := stat.(type) {
	case *stats.Begin:
		a.runtime.Events.Emit(eventStatBegin, s)
	case *stats.OutHeader:
		a.runtime.Events.Emit(eventStatOutHeader, rpcStatOutHeader{s, fmt.Sprintf("%+v", s.Header)})
	case *stats.OutPayload:
		if p, err := formatPayload(s.Payload); err == nil {
			s.Payload = p
		}
		a.runtime.Events.Emit(eventStatOutPayload, rpcStatOutPayload{s, fmt.Sprintf("%+v", s.Data)})
	case *stats.OutTrailer:
		a.runtime.Events.Emit(eventStatOutTrailer, rpcStatOutTrailer{s, fmt.Sprintf("%+v", s.Trailer)})
	case *stats.InHeader:
		a.runtime.Events.Emit(eventStatInHeader, rpcStatInHeader{s, fmt.Sprintf("%+v", s.Header)})
		a.runtime.Events.Emit(eventInHeaderReceived, s.Header)
	case *stats.InPayload:
		txt, err := formatPayload(s.Payload)
		if err != nil {
			a.logger.Errorf("failed to marshal in payload to proto text: %v", err)
			return
		}
		s.Payload = txt
		a.runtime.Events.Emit(eventStatInPayload, rpcStatInPayload{s, fmt.Sprintf("%+v", s.Data)})
		a.runtime.Events.Emit(eventInPayloadReceived, txt)
	case *stats.InTrailer:
		a.runtime.Events.Emit(eventStatInTrailer, rpcStatInTrailer{s, fmt.Sprintf("%+v", s.Trailer)})
		a.runtime.Events.Emit(eventInTrailerReceived, s.Trailer)
	case *stats.End:

		errProtoStr := ""
		stus := status.Convert(s.Error)
		if stus != nil {
			var err error
			errProtoStr, err = formatPayload(stus.Proto())
			if err != nil {
				a.logger.Errorf("failed to marshal status error to proto text: %v", err)
			}
			if errProtoStr != "" {
				a.runtime.Events.Emit(eventInPayloadReceived, errProtoStr)
			}
		}
		a.runtime.Events.Emit(eventStatEnd, rpcStatEnd{s, errProtoStr})

		var end rpcEnd
		end.StatusCode = int32(stus.Code())
		end.Status = stus.Code().String()
		end.Duration = s.EndTime.Sub(s.BeginTime).String()
		a.runtime.Events.Emit(eventRPCEnded, end)
	}
}

func formatPayload(payload interface{}) (string, error) {
	msg, ok := payload.(proto.Message)
	if !ok {
		// check to see if we are dealing with a APIv1 message
		msgV1, ok := payload.(protoV1.Message)
		if !ok {
			return "", fmt.Errorf("payload is not a proto message: %T", payload)
		}
		msg = protoV1.MessageV2(msgV1)
	}

	marshaler := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}
	b, err := marshaler.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// CloseSend will stop streaming client messages
func (a *api) CloseSend() {
	if a.streamReq != nil {
		close(a.streamReq)
		a.streamReq = nil
	}
}

// Cancel will attempt to cancel the current inflight request
func (a *api) Cancel() {
	if a.cancelInFlight != nil {
		a.cancelInFlight()
	}
}
