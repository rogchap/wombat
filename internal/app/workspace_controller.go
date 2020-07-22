// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/therecipe/qt/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"rogchap.com/courier/internal/model"
)

//go:generate qtmoc
type workspaceController struct {
	core.QObject

	grpcConn      *grpc.ClientConn
	cancelCtxFunc context.CancelFunc

	_ func() `constructor:"init"`

	_ *inputController  `property:"inputCtrl"`
	_ *outputController `property:"outputCtrl"`

	_ *model.WorkspaceOptions `property:"options"`
	_ string                  `property:"connState"`

	_ func(path string)            `slot:"findProtoFiles"`
	_ func(path string)            `slot:"addImport"`
	_ func() error                 `slot:"processProtos"`
	_ func(addr string) error      `slot:"connect"`
	_ func(service, method string) `slot:"send"`
}

func (c *workspaceController) init() {
	c.SetInputCtrl(NewInputController(nil))
	c.SetOutputCtrl(NewOutputController(nil))

	c.SetOptions(model.NewWorkspaceOptions(nil))

	c.ConnectFindProtoFiles(c.findProtoFiles)
	c.ConnectAddImport(c.addImport)
	c.ConnectProcessProtos(c.processProtos)
	c.ConnectConnect(c.connect)
	c.ConnectSend(c.send)
}

func (c *workspaceController) findProtoFiles(path string) {
	var protoFiles []string

	// TODO [RC] We should do the search async and show a loading/searching icon to the user
	filepath.Walk(path[7:], func(path string, info os.FileInfo, err error) error {
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
	path = path[7:]
	lm := c.Options().ImportListModel()
	for _, p := range lm.StringList() {
		if p == path {
			return
		}
	}
	lm.SetStringList(append(lm.StringList(), path))
}

func (c *workspaceController) processProtos() error {
	imports := c.Options().ImportListModel().StringList()
	protos := c.Options().ProtoListModel().StringList()
	return c.InputCtrl().processProtos(imports, protos)
}

func (c *workspaceController) connect(addr string) error {
	if addr == "" {
		return errors.New("no address to connect")
	}

	if c.Options().Addr() == addr {
		return nil
	}

	if c.grpcConn != nil {
		c.grpcConn.Close()
		c.cancelCtxFunc()
	}

	var err error
	c.grpcConn, err = BlockDial(addr, c.Options())
	if err != nil {
		// TODO [RC] hndle error back to user
		println(err.Error())
		return err
	}

	var ctx context.Context
	ctx, c.cancelCtxFunc = context.WithCancel(context.Background())

	go func() {
		for {
			state := c.grpcConn.GetState()
			MainThread.Run(func() {
				c.SetConnState(state.String())
			})
			if ok := c.grpcConn.WaitForStateChange(ctx, state); !ok {
				break
			}
		}
	}()

	c.Options().SetAddr(addr)
	return nil
}

func (c *workspaceController) send(service, method string) {
	if c.grpcConn == nil {
		return
	}

	inputCtrl := c.InputCtrl()
	outputCtrl := c.OutputCtrl()

	outputCtrl.clear()

	go func() {
		md := inputCtrl.pbSource.GetMethodDesc(service, method)

		req := processMessage(c.InputCtrl().RequestModel())
		stub := grpcdynamic.NewStub(c.grpcConn)

		if md.IsClientStreaming() && md.IsServerStreaming() {
			println("Bidirectional streaming not supported yet")
			return
		}

		if md.IsClientStreaming() {
			println("Client streaming not supported yet")
		}

		if md.IsServerStreaming() {
			stream, err := stub.InvokeRpcServerStream(context.Background(), md, req)
			if err != nil {
				println(err.Error())
				return
			}
			for {
				resp, err := stream.RecvMsg()
				_ = resp
				if err == io.EOF {

					MainThread.Run(func() {
						outputCtrl.SetStatus(0)
						outputCtrl.SetOutput(fmt.Sprintf("%sEOF\n", outputCtrl.Output()))
					})
					break
				}
				if err != nil {
					println(err.Error())
					MainThread.Run(func() {
						outputCtrl.SetStatus(int(status.Code(err)))
					})
					return
				}
				dm, _ := dynamic.AsDynamicMessage(resp)
				b, _ := dm.MarshalTextIndent()

				MainThread.Run(func() {
					outputCtrl.SetOutput(fmt.Sprintf("%s%+v\n", outputCtrl.Output(), string(b)))
				})
			}
			return
		}
		var header, trailer metadata.MD
		resp, err := stub.InvokeRpc(context.Background(), md, req, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			println(err.Error())
			st := status.Convert(err)
			dm, _ := dynamic.AsDynamicMessage(st.Proto())
			b, _ := dm.MarshalTextIndent()
			MainThread.Run(func() {
				outputCtrl.SetStatus(int(status.Code(err)))
				outputCtrl.SetOutput(string(b))
			})
			return
		}

		dm, _ := dynamic.AsDynamicMessage(resp)
		b, _ := dm.MarshalTextIndent()

		fmt.Printf("%+v\n", header)
		fmt.Printf("%+v\n", trailer)
		MainThread.Run(func() {
			outputCtrl.SetStatus(0)
			outputCtrl.SetOutput(fmt.Sprintf("%+v\n", string(b)))
		})
	}()
}

func processMessage(msg *model.Message) *dynamic.Message {
	dm := dynamic.NewMessage(msg.Ref)
	for _, f := range msg.Fields() {
		switch f.FdType {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if f.IsRepeated {
				for idx, v := range f.ValueListModel().Values() {
					dm.SetRepeatedFieldByNumber(f.Tag(), idx, processMessage(v.MsgValue()))
				}
				break
			}
			dm.SetFieldByNumber(f.Tag(), processMessage(f.Message()))

		default:
			if f.IsRepeated {
				for idx, v := range f.ValueListModel().Values() {
					dm.SetRepeatedFieldByNumber(f.Tag(), idx, parseStringValue(f.FdType, v.Value()))
				}
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
