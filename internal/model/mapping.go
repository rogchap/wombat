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

		v, _ := dm.TryGetField(fd)
		field.SetValue(stringValue(v, fd))

		switch ft {
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			enabled := true
			var mdm *dynamic.Message
			switch v.(type) {
			case *dynamic.Message:
				mdm = v.(*dynamic.Message)
			case proto.Message:
				mdm, _ = dynamic.AsDynamicMessage(v.(proto.Message))
			}
			if mdm == nil {
				enabled = false
				mdm = dynamic.NewMessage(fd.GetMessageType())

			}
			msg := MapMessage(mdm)
			msg.SetEnabled(enabled)
			field.SetMessage(msg)
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
			// TODO: populate any repeated fields that in the store
			var rVals []*RepeatedValue
			lg, _ := dm.TryFieldLength(fd)
			for i := 0; i < lg; i++ {
				rval := NewRepeatedValue(nil)
				val, _ := dm.TryGetRepeatedField(fd, i)
				rval.SetValue(stringValue(val, fd))

				if ft == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {

					enabled := true
					var mdm *dynamic.Message
					switch val.(type) {
					case *dynamic.Message:
						mdm = val.(*dynamic.Message)
					case proto.Message:
						mdm, _ = dynamic.AsDynamicMessage(val.(proto.Message))
					}
					if mdm == nil {
						enabled = false
						mdm = dynamic.NewMessage(fd.GetMessageType())

					}
					msg := MapMessage(mdm)
					msg.SetEnabled(enabled)
					rval.SetMsgValue(msg)

				}

				rVals = append(rVals, rval)
			}

			vl.SetValues(rVals)
			vl.SetCount(lg)
			field.SetValueListModel(vl)
		}

		fields = append(fields, field)
	}

	msg.SetFields(fields)
	return msg
}

func stringValue(v interface{}, fd *desc.FieldDescriptor) string {
	if v == nil {
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
