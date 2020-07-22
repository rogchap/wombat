// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"strconv"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
)

func MapMessage(md *desc.MessageDescriptor) *Message {
	msg := NewMessage(nil)
	msg.Ref = md
	msg.SetLabel(md.GetFullyQualifiedName())

	var fields []*Field
	for _, fd := range md.GetFields() {
		field := NewField(nil)
		field.SetLabel(fd.GetName())
		field.SetTag(int(fd.GetNumber()))
		field.SetFullname(fd.GetFullyQualifiedName())

		ft := fd.GetType()
		field.FdType = ft
		typeName := strings.ToLower(descriptor.FieldDescriptorProto_Type_name[int32(ft)][5:])
		field.SetType(typeName)

		switch ft {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			field.SetMessage(MapMessage(fd.GetMessageType()))
			field.SetDelegate("message")
		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			e := fd.GetEnumType()
			var enumValues []keyval
			for _, enum := range e.GetValues() {
				enumValues = append(enumValues, keyval{enum.GetName(), strconv.Itoa(int(enum.GetNumber()))})
			}
			enumListModel := NewKeyvalList(nil)
			enumListModel.list = enumValues
			field.SetEnumListModel(enumListModel)
			field.SetDelegate("enum")
		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			field.SetDelegate("textArea")
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			field.SetDelegate("bool")
		default:
			field.SetDelegate("text")
		}

		if fd.IsRepeated() {
			field.IsRepeated = true
			field.SetDelegate(field.Delegate() + "_repeated")
			vl := NewRepeatedValues(nil)
			vl.ref = fd.GetMessageType()
			field.SetValueListModel(vl)
		}

		fields = append(fields, field)
	}

	msg.SetFields(fields)
	return msg
}
