// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/therecipe/qt/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"rogchap.com/wombat/internal/db"
	"rogchap.com/wombat/internal/model"
)

const defaultWorkspaceKey = "wksp_default"

//go:generate qtmoc
type workspaceController struct {
	core.QObject

	grpcConn      *grpc.ClientConn
	cancelCtxFunc context.CancelFunc
	store         *db.Store
	workspace     *db.Workspace

	_ func() `constructor:"init"`

	_ *inputController  `property:"inputCtrl"`
	_ *outputController `property:"outputCtrl"`

	_ *model.WorkspaceOptions `property:"options"`
	_ string                  `property:"connState"`

	_ func(path string)                  `slot:"findProtoFiles"`
	_ func(path string)                  `slot:"addImport"`
	_ func() error                       `slot:"processProtos"`
	_ func(addr string) error            `slot:"connect"`
	_ func(service, method string) error `slot:"send"`
}

func (c *workspaceController) init() {

	c.ConnectFindProtoFiles(c.findProtoFiles)
	c.ConnectAddImport(c.addImport)
	c.ConnectProcessProtos(c.processProtos)
	c.ConnectConnect(c.connect)
	c.ConnectSend(c.send)

	dbPath := core.QStandardPaths_WritableLocation(core.QStandardPaths__AppDataLocation)
	if isDebug {
		dbPath = filepath.Join(".", ".data")
	}

	var err error
	c.store, err = db.NewStore(dbPath)
	if err != nil {
		println(err.Error())
	}

	c.workspace, err = c.store.GetWorkspace(defaultWorkspaceKey)
	if err != nil {
		println(err.Error())
	}

	c.SetInputCtrl(NewInputController(nil).with(c.store, c.workspace))
	c.SetOutputCtrl(NewOutputController(nil))

	c.SetOptions(model.NewWorkspaceOptions(nil))
	if o := c.workspace.GetOptions(); o != nil {
		opts := c.Options()
		opts.SetReflect(o.Reflect)
		opts.SetInsecure(o.Insecure)
		opts.SetPlaintext(o.Plaintext)
		opts.SetRootca(o.Rootca)
		opts.SetClientcert(o.Clientcert)
		opts.SetClientkey(o.Clientkey)
		opts.ProtoListModel().SetStringList(o.ProtoFiles)
		opts.ImportListModel().SetStringList(o.ImportFiles)
	}
	c.connect(c.workspace.Addr)
	c.processProtos()
}

func (c *workspaceController) findProtoFiles(path string) {
	path = core.NewQUrl3(path, core.QUrl__StrictMode).ToLocalFile()
	var protoFiles []string

	// TODO [RC] We should do the search async and show a loading/searching icon to the user
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if len(protoFiles) == 0 {
		// TODO [RC] Show error to user that there is no proto files found
		return
	}

	// TODO [RC] Shoud we be replacing or adding?
	c.Options().ProtoListModel().SetStringList(protoFiles)
}

func (c *workspaceController) addImport(path string) {
	path = core.NewQUrl3(path, core.QUrl__StrictMode).ToLocalFile()
	lm := c.Options().ImportListModel()
	for _, p := range lm.StringList() {
		if p == path {
			return
		}
	}
	lm.SetStringList(append(lm.StringList(), path))
}

func (c *workspaceController) processProtos() error {
	if c.Options().IsReflect() {
		return c.InputCtrl().processReflectionAPI(c.grpcConn)
	}

	imports := c.Options().ImportListModel().StringList()
	protos := c.Options().ProtoListModel().StringList()
	return c.InputCtrl().processProtos(imports, protos)
}

func (c *workspaceController) connect(addr string) error {
	if addr == "" {
		return errors.New("no address to connect")
	}

	if c.grpcConn != nil {
		c.grpcConn.Close()
		c.cancelCtxFunc()
		c.grpcConn = nil
	}

	var err error
	c.grpcConn, err = BlockDial(addr, c.Options(), c.OutputCtrl())
	if err != nil {
		return err
	}

	var ctx context.Context
	ctx, c.cancelCtxFunc = context.WithCancel(context.Background())

	go func() {
		defer func() {
			// TODO(rogchap) Should be a better way than swallowing this panic?
			recover()
		}()

		for {
			if c.grpcConn == nil {
				c.SetConnState(connectivity.Shutdown.String())
				break

			}
			state := c.grpcConn.GetState()
			c.SetConnState(state.String())
			if ok := c.grpcConn.WaitForStateChange(ctx, state); !ok {
				break
			}
		}
	}()

	c.Options().SetAddr(addr)

	go func() {
		opts := c.Options()
		c.workspace.Options = &db.Workspace_Options{
			Reflect:     opts.IsReflect(),
			Insecure:    opts.IsInsecure(),
			Plaintext:   opts.IsPlaintext(),
			Rootca:      opts.Rootca(),
			Clientcert:  opts.Clientcert(),
			Clientkey:   opts.Clientkey(),
			ProtoFiles:  opts.ProtoListModel().StringList(),
			ImportFiles: opts.ImportListModel().StringList(),
		}
		c.workspace.Addr = addr

		c.store.SetWorkspace(defaultWorkspaceKey, c.workspace)
	}()
	return nil
}

func (c *workspaceController) send(service, method string) error {
	if c.grpcConn == nil {
		return nil
	}

	md := c.InputCtrl().pbSource.GetMethodDesc(service, method)
	req := processMessage(c.InputCtrl().RequestModel())

	if data, err := req.Marshal(); err == nil {
		c.store.Set([]byte(md.GetFullyQualifiedName()), data)
	}

	meta := make(map[string]string)
	for _, kv := range c.InputCtrl().MetadataListModel().List() {
		if kv.Key() == "" {
			continue
		}
		meta[kv.Key()] = kv.Val()
	}
	c.workspace.Metadata = meta
	c.store.SetWorkspace(defaultWorkspaceKey, c.workspace)

	return c.OutputCtrl().invokeMethod(c.grpcConn, md, req, meta)
}

func processMessage(msg *model.Message) *dynamic.Message {
	if !msg.IsEnabled() {
		return nil
	}
	dm := dynamic.NewMessage(msg.Ref)
	for _, f := range msg.Fields() {
		switch f.FdType {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if f.IsRepeated {
				var fields []interface{}
				for _, v := range f.ValueListModel().Values() {
					m := processMessage(v.MsgValue())
					if m != nil {
						fields = append(fields, m)
					}
				}
				dm.SetFieldByNumber(f.Tag(), fields)
				break
			}
			m := processMessage(f.Message())
			if m != nil {
				dm.SetFieldByNumber(f.Tag(), m)
			}
		default:
			if f.IsRepeated {
				var fields []interface{}
				for _, v := range f.ValueListModel().Values() {
					fields = append(fields, parseStringValue(f.FdType, v.Value()))
				}
				dm.SetFieldByNumber(f.Tag(), fields)
				break
			}
			dm.SetFieldByNumber(f.Tag(), parseStringValue(f.FdType, f.Value()))
		}
	}

	return dm
}

func parseStringValue(fdType descriptor.FieldDescriptorProto_Type, val string) interface{} {
	switch fdType {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		v, _ := strconv.ParseFloat(val, 64)
		return v
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		v, _ := strconv.ParseFloat(val, 32)
		return float32(v)
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		v, _ := strconv.ParseInt(val, 10, 32)
		return int32(v)
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		v, _ := strconv.ParseInt(val, 10, 64)
		return v
	case descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32:
		v, _ := strconv.ParseUint(val, 10, 32)
		return uint32(v)
	case descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		v, _ := strconv.ParseUint(val, 10, 64)
		return v
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		v, _ := strconv.ParseBool(val)
		return v
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return []byte(val)
	default:
		return val
	}
}
