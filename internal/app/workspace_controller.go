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

	_ string            `property:"addr"`
	_ string            `property:"connState"`
	_ *model.StringList `property:protoListModel"`
	_ *model.StringList `property:importListModel"`

	_ func(path string)            `slot:"findProtoFiles"`
	_ func(path string)            `slot:"addImport"`
	_ func() error                 `slot:"processProtos"`
	_ func(addr string) error      `slot:"connect"`
	_ func(service, method string) `slot:"send"`
}

func (c *workspaceController) init() {
	c.SetInputCtrl(NewInputController(nil))
	c.SetOutputCtrl(NewOutputController(nil))

	c.SetProtoListModel(model.NewStringList(nil))
	c.SetImportListModel(model.NewStringList(nil))

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
	c.ProtoListModel().SetStringList(protoFiles)
}

func (c *workspaceController) addImport(path string) {
	path = path[7:]
	lm := c.ImportListModel()
	for _, p := range lm.StringList() {
		if p == path {
			return
		}
	}
	lm.SetStringList(append(lm.StringList(), path))
}

func (c *workspaceController) processProtos() error {
	imports := c.ImportListModel().StringList()
	protos := c.ProtoListModel().StringList()
	return c.InputCtrl().processProtos(imports, protos)
}

func (c *workspaceController) connect(addr string) error {
	if addr == "" {
		return errors.New("no address to connect")
	}

	if c.Addr() == addr {
		return nil
	}

	if c.grpcConn != nil {
		c.grpcConn.Close()
		c.cancelCtxFunc()
	}
	// TODO [RC] setup grpc options
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.FailOnNonTempDialError(true),
	}

	var err error
	c.grpcConn, err = grpc.Dial(addr, opts...)
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

	// TODO [RC] monitor connection status

	c.SetAddr(addr)
	return nil
}

func (c *workspaceController) send(service, method string) {
	inputCtrl := c.InputCtrl()
	outputCtrl := c.OutputCtrl()

	outputCtrl.clear()

	go func() {

		md := inputCtrl.pbSource.GetMethodDesc(service, method)
		req := dynamic.NewMessage(md.GetInputType())
		processFields(req, c.InputCtrl().RequestModel().Fields())

		stub := grpcdynamic.NewStub(c.grpcConn)

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
				MainThread.Run(func() {
					outputCtrl.SetOutput(fmt.Sprintf("%s%+v\n", outputCtrl.Output(), resp))
				})
			}
			return
		}
		var header, trailer metadata.MD
		resp, err := stub.InvokeRpc(context.Background(), md, req, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			println(err.Error())
			MainThread.Run(func() {
				outputCtrl.SetStatus(int(status.Code(err)))
			})
			return
		}

		fmt.Printf("%+v\n", resp)
		fmt.Printf("%+v\n", header)
		fmt.Printf("%+v\n", trailer)
		MainThread.Run(func() {
			outputCtrl.SetStatus(0)
			outputCtrl.SetOutput(fmt.Sprintf("%+v\n", resp))
		})
	}()
}

func processFields(msg *dynamic.Message, fields []*model.Field) {
	for _, f := range fields {
		switch descriptor.FieldDescriptorProto_Type(descriptor.FieldDescriptorProto_Type_value[f.Type()]) {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			fd := msg.FindFieldDescriptor(int32(f.Tag()))
			m := dynamic.NewMessage(fd.GetMessageType())
			processFields(m, f.Message().Fields())
			msg.SetFieldByNumber(f.Tag(), m)

		case descriptor.FieldDescriptorProto_TYPE_INT32:
			v, _ := strconv.Atoi(f.Value())
			msg.SetFieldByNumber(f.Tag(), int32(v))

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			msg.SetFieldByNumber(f.Tag(), f.Value())
		}

		msg.AddRepeatedFieldByNumber
	}
}
