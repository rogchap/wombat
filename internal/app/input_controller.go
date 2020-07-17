// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/therecipe/qt/core"
	"rogchap.com/courier/internal/model"
	"rogchap.com/courier/internal/pb"
)

//go:generate qtmoc
type inputController struct {
	core.QObject

	pbSource pb.Source

	_ func() `constructor:"init"`

	_ *model.StringList `property:"serviceListModel"`
	_ *model.StringList `property:"methodListModel"`
	_ *model.Message    `property:"requestModel"`

	_ func(service string)         `slot:"serviceChanged"`
	_ func(service, method string) `slot:"methodChanged"`
}

func (c *inputController) init() {
	c.SetServiceListModel(model.NewStringList(nil))
	c.SetMethodListModel(model.NewStringList(nil))
	c.SetRequestModel(model.NewMessage(nil))

	c.ConnectServiceChanged(c.serviceChanged)
	c.ConnectMethodChanged(c.methodChanged)
}

func (c *inputController) processProtos(imports, protos []string) error {
	if len(protos) == 0 {
		return errors.New("no *.proto files to process")
	}
	if len(imports) == 0 {
		// optomistacally try and use a import path
		imports = append(imports, filepath.Dir(protos[0]))
	}

	var err error
	c.pbSource, err = pb.GetSourceFromProtoFiles(imports, protos)
	if err != nil {
		return err
	}

	services := c.pbSource.Services()
	if len(services) == 0 {
		return errors.New("no gRPC services found in proto files")
	}

	c.ServiceListModel().SetStringList(services)
	c.serviceChanged(services[0])
	return nil
}

func (c *inputController) serviceChanged(service string) {
	methods := c.pbSource.Methods()

	srvMethods, ok := methods[service]
	if !ok {
		return
	}
	var methodStrs []string
	for _, m := range srvMethods {
		methodStrs = append(methodStrs, m.GetName())
	}

	c.MethodListModel().SetStringList(methodStrs)
	c.methodChanged(service, methodStrs[0])
}

func (c *inputController) methodChanged(service, method string) {
	md := c.pbSource.GetMethodDesc(service, method)
	if md == nil {
		return
	}

	input := md.GetInputType()
	reqModel := c.RequestModel()

	reqModel.BeginResetModel()
	reqModel.SetLabel(input.GetFullyQualifiedName())
	reqModel.SetFields(getMessageFields(input))
	reqModel.EndResetModel()
}

func getMessageFields(msg *desc.MessageDescriptor) []*model.Field {
	var fields []*model.Field
	for _, f := range msg.GetFields() {
		field := model.NewField(nil)
		field.SetLabel(f.GetName())
		field.SetTag(int(f.GetNumber()))
		field.SetFullname(f.GetFullyQualifiedName())

		ft := f.GetType()
		typeName := strings.ToLower(descriptor.FieldDescriptorProto_Type_name[int32(ft)][5:])
		field.SetType(typeName)

		switch ft {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			m := f.GetMessageType()
			msg := model.NewMessage(nil)
			msg.SetLabel(m.GetFullyQualifiedName())
			msg.SetFields(getMessageFields(m))
			field.SetMessage(msg)
			field.SetDelegate("message")
		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			e := f.GetEnumType()
			var enumValues []string
			for _, enum := range e.GetValues() {
				enumValues = append(enumValues, enum.GetName())
			}
			enumListModel := model.NewStringList(nil)
			enumListModel.SetStringList(enumValues)
			field.SetEnumListModel(enumListModel)
			field.SetDelegate("enum")
		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			field.SetDelegate("textArea")
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			field.SetDelegate("bool")
		default:
			field.SetDelegate("text")
		}

		if f.IsRepeated() {
			field.SetDelegate(field.Delegate() + "_repeated")
			// TODO If field is a message we need to use a list of messages not strings
			vl := model.NewRepeatedValues(nil)
			field.SetValueListModel(vl)
		}

		fields = append(fields, field)
	}
	return fields
}
