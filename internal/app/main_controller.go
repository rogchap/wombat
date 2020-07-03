// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

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
	_ *model.Input      `property:"input"`

	_ string `property:"output"`

	_ func(path string)                  `slot:"findProtoFiles"`
	_ func(path string)                  `slot:"addImport"`
	_ func(imports, path string)         `slot:"processProtos"`
	_ func(service string)               `slot:"serviceChanged"`
	_ func(host, service, method string) `slot:"send"`
}

func (c *mainController) init() {
	c.SetProtoFilesList(model.NewStringList(nil))
	c.SetProtoImportsList(model.NewStringList(nil))
	c.SetServiceList(model.NewStringList(nil))
	c.SetMethodList(model.NewStringList(nil))
	c.SetInput(model.NewInput(nil))

	c.ConnectFindProtoFiles(c.findProtoFiles)
	c.ConnectAddImport(c.addImport)
	c.ConnectProcessProtos(c.processProtos)
	c.ConnectServiceChanged(c.serviceChanged)
	c.ConnectSend(c.send)
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

	input := srvMethods[0].GetInputType()
	c.Input().SetLabel(input.GetFullyQualifiedName())

	var fields []*model.Field
	for _, f := range input.GetFields() {
		field := model.NewField(nil)
		field.SetLabel(f.GetName())
		field.SetTag(int(f.GetNumber()))
		fields = append(fields, field)
	}
	c.Input().BeginResetModel()
	c.Input().SetFields(fields)
	c.Input().EndResetModel()
}

func (c *mainController) send(host, service, method string) {
	c.SetOutput("")

	cc, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		//TODO: handle error
		println(err.Error())
		return
	}
	defer cc.Close()

	md := c.pbSource.GetMethodDesc(service, method)
	req := dynamic.NewMessage(md.GetInputType())

	// req.SetFieldByNumber(1, int32(1))
	// req.SetFieldByNumber(2, int32(2))
	for _, f := range c.Input().Fields() {
		v, _ := strconv.Atoi(f.Value())
		req.SetFieldByNumber(f.Tag(), int32(v))
	}

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
