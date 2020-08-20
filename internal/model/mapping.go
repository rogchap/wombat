// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/descriptorpb"
)

func MapMessage(dm *dynamic.Message) *Message {
	md := dm.GetMessageDescriptor()
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
		typeName := strings.ToLower(descriptorpb.FieldDescriptorProto_Type_name[int32(ft)][5:])
		field.SetType(typeName)

		field.SetValue(stringValue(dm, fd))

		switch ft {
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			val, _ := dm.TryGetField(fd)
			var mdm *dynamic.Message
			switch val.(type) {
			case *dynamic.Message:
				mdm = val.(*dynamic.Message)
			case proto.Message:
				mdm, _ = dynamic.AsDynamicMessage(val.(proto.Message))
			}
			if mdm == nil {
				mdm = dynamic.NewMessage(fd.GetMessageType())
			}
			field.SetMessage(MapMessage(mdm))
			field.SetDelegate("message")
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
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
		case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			field.SetDelegate("textArea")
		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
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

func stringValue(dm *dynamic.Message, fd *desc.FieldDescriptor) string {
	v, err := dm.TryGetField(fd)
	if err != nil {
		return ""
	}

	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return string(v.([]byte))
		}
		return ""
	case reflect.Map, reflect.Struct:
		return ""
	default:
		if fd.GetType() == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
			return strconv.Itoa(int(v.(int32)))
		}
		// only return the string if it is not the default value
		if t.Comparable() {
			dv := fd.GetDefaultValue()
			if dv != v {
				return fmt.Sprintf("%v", v)
			}
		}
		return ""
	}
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
