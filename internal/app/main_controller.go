// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/therecipe/qt/core"
	"google.golang.org/grpc"

	"rogchap.com/courier/internal/model"
	"rogchap.com/courier/internal/pb"
)

//go:generate qtmoc
type mainController struct {
	core.QObject

	pbSource pb.Source

	_ func() `constructor:"init"`

	_ *model.StringList `property:protoFilesList"`
	_ *model.StringList `property:protoImportsList"`
	_ *model.StringList `property:"serviceList"`
	_ *model.StringList `property:"methodList"`
	_ *model.Message    `property:"input"`

	_ string `property:"addr"`
	_ string `property:"output"`

	_ func(addr string)            `slot:"updateAddr"`
	_ func(path string)            `slot:"findProtoFiles"`
	_ func(path string)            `slot:"addImport"`
	_ func(imports, path string)   `slot:"processProtos"`
	_ func(service string)         `slot:"serviceChanged"`
	_ func(service, method string) `slot:"methodChanged"`
	_ func(service, method string) `slot:"send"`
}

func (c *mainController) init() {
	c.SetProtoFilesList(model.NewStringList(nil))
	c.SetProtoImportsList(model.NewStringList(nil))
	c.SetServiceList(model.NewStringList(nil))
	c.SetMethodList(model.NewStringList(nil))
	c.SetInput(model.NewMessage(nil))

	c.ConnectUpdateAddr(c.updateAddr)
	c.ConnectFindProtoFiles(c.findProtoFiles)
	c.ConnectAddImport(c.addImport)
	c.ConnectProcessProtos(c.processProtos)
	c.ConnectServiceChanged(c.serviceChanged)
	c.ConnectMethodChanged(c.methodChanged)
	c.ConnectSend(c.send)
}

func (c *mainController) updateAddr(addr string) {
	if c.Addr() == addr {
		return
	}
	c.SetAddr(addr)
}

func (c *mainController) findProtoFiles(path string) {
	var protoFiles []string
	filepath.Walk(path[7:], func(path string, info os.FileInfo, err error) error {

		//TODO: only add if we haven't got it already
		if filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if len(protoFiles) == 0 {
		return
		// TODO: Show error to user that there is no proto files found
	}

	// TODO: Shoud we be replacing or adding?
	c.ProtoFilesList().SetStringList(protoFiles)
}

func (c *mainController) addImport(path string) {
	path = path[7:]
	lm := c.ProtoImportsList()
	for _, p := range lm.StringList() {
		if p == path {
			return
		}
	}
	lm.SetStringList(append(lm.StringList(), path))
}

func (c *mainController) processProtos(imports, path string) {

	var err error
	c.pbSource, err = pb.GetSourceFromProtoFiles(c.ProtoImportsList().StringList(), c.ProtoFilesList().StringList())
	if err != nil {
		println(err.Error())
		return
	}

	services := c.pbSource.Services()
	if len(services) == 0 {
		// TODO: Show error that there are no servcies found
		return
	}
	c.ServiceList().SetStringList(services)
	c.serviceChanged(services[0])
}

func (c *mainController) serviceChanged(service string) {
	methods := c.pbSource.Methods()

	srvMethods, ok := methods[service]
	if !ok {
		return
	}
	var methodStrs []string
	for _, m := range srvMethods {
		methodStrs = append(methodStrs, m.GetName())
	}

	c.MethodList().SetStringList(methodStrs)
	c.methodChanged(service, methodStrs[0])
}

func (c *mainController) methodChanged(service, method string) {
	md := c.pbSource.GetMethodDesc(service, method)
	if md == nil {
		return
	}

	input := md.GetInputType()
	c.Input().BeginResetModel()
	c.Input().SetLabel(input.GetFullyQualifiedName())
	c.Input().SetFields(getMessageFields(input))
	c.Input().EndResetModel()
}

func getMessageFields(msg *desc.MessageDescriptor) []*model.Field {
	var fields []*model.Field
	for _, f := range msg.GetFields() {
		field := model.NewField(nil)
		field.SetLabel(f.GetName())
		field.SetTag(int(f.GetNumber()))
		field.SetFullname(f.GetFullyQualifiedName())

		ft := f.GetType()
		field.SetType(descriptor.FieldDescriptorProto_Type_name[int32(ft)])

		switch ft {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			m := f.GetMessageType()
			msg := model.NewMessage(nil)
			msg.SetLabel(m.GetFullyQualifiedName())
			msg.SetFields(getMessageFields(m))
			field.SetMessage(msg)
		}
		fields = append(fields, field)
	}
	return fields
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
	}
}

func (c *mainController) send(service, method string) {
	c.SetOutput("")

	cc, err := grpc.Dial(c.Addr(), grpc.WithInsecure())
	if err != nil {
		//TODO: handle error
		println(err.Error())
		return
	}
	defer cc.Close()

	md := c.pbSource.GetMethodDesc(service, method)
	req := dynamic.NewMessage(md.GetInputType())

	processFields(req, c.Input().Fields())

	stub := grpcdynamic.NewStub(cc)

	if md.IsServerStreaming() {
		stream, err := stub.InvokeRpcServerStream(context.Background(), md, req)
		if err != nil {
			println(err.Error())
			return
		}
		for {
			resp, err := stream.RecvMsg()
			if err == io.EOF {
				c.SetOutput(fmt.Sprintf("%sEOF\n", c.Output()))
				break
			}
			if err != nil {
				println(err.Error())
				return
			}
			c.SetOutput(fmt.Sprintf("%s%+v\n", c.Output(), resp))
		}
		return
	}

	resp, err := stub.InvokeRpc(context.Background(), md, req)
	if err != nil {
		println(err.Error())
	}

	c.SetOutput(fmt.Sprintf("%+v\n", resp))
}
