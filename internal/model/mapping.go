// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
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
			var enumValues []*Keyval
			for _, enum := range e.GetValues() {
				kv := NewKeyval(nil)
				kv.SetKey(enum.GetName())
				kv.SetVal(strconv.Itoa(int(enum.GetNumber())))
				enumValues = append(enumValues, kv)
			}
			enumListModel := NewKeyvalList(nil)
			enumListModel.SetList(enumValues)
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

func MapMetadata(md metadata.MD) []*Keyval {
	keys := make([]string, 0, len(md))
	for k := range md {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var kvs []*Keyval
	for _, k := range keys {
		kv := NewKeyval(nil)
		kv.SetKey(k)
		kv.SetVal(strings.Join(md.Get(k), ", "))
		kvs = append(kvs, kv)
	}
	return kvs
}
