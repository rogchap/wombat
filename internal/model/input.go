// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/therecipe/qt/core"
)

const (
	FieldType = int(core.Qt__UserRole) + 1<<iota
	FieldDelegate
	FieldLabel
	FieldFullname
	FieldVal
	FieldValues
	FieldMsg
	FieldEnum
)

//go:generate qtmoc
type Field struct {
	core.QObject

	FdType     descriptor.FieldDescriptorProto_Type
	IsRepeated bool

	_ string          `property:"type"`
	_ string          `property:"delegate"`
	_ string          `property:"label"`
	_ string          `property:"fullname"`
	_ int             `property:"tag"`
	_ string          `property:"value"`
	_ *RepeatedValues `property:"valueListModel"`
	_ *Message        `property:"message"`
	_ *KeyvalList     `property:"enumListModel"`
}

//go:generate qtmoc
type Message struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	Ref *desc.MessageDescriptor

	_ map[int]*core.QByteArray `property:"roles"`
	_ string                   `property:"label"`
	_ bool                     `property:"enabled"`
	_ []*Field                 `property:"fields"`

	_ func(row int, val string) `slot:"updateFieldValue"`
}

func (i *Message) init() {
	i.SetRoles(map[int]*core.QByteArray{
		FieldType:     core.NewQByteArray2("type", -1),
		FieldDelegate: core.NewQByteArray2("delegate", -1),
		FieldLabel:    core.NewQByteArray2("label", -1),
		FieldFullname: core.NewQByteArray2("fullname", -1),
		FieldVal:      core.NewQByteArray2("val", -1),
		FieldValues:   core.NewQByteArray2("valueListModel", -1),
		FieldMsg:      core.NewQByteArray2("message", -1),
		FieldEnum:     core.NewQByteArray2("enumListModel", -1),
	})

	// TODO: [RC] Should we default to enabled for nested meeasges?
	// Maybe this should be a setting?
	i.SetEnabled(true)

	i.ConnectData(i.data)
	i.ConnectRowCount(i.rowCount)
	i.ConnectColumnCount(i.columnCount)
	i.ConnectRoleNames(i.roleNames)

	i.ConnectUpdateFieldValue(i.updateFieldValue)
}

func (i *Message) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(i.Fields()) {
		return core.NewQVariant()
	}

	var f = i.Fields()[index.Row()]

	switch role {
	case FieldType:
		return core.NewQVariant1(f.Type())
	case FieldDelegate:
		return core.NewQVariant1(f.Delegate())
	case FieldLabel:
		return core.NewQVariant1(f.Label())
	case FieldFullname:
		return core.NewQVariant1(f.Fullname())
	case FieldVal:
		return core.NewQVariant1(f.Value())
	case FieldValues:
		return core.NewQVariant1(f.ValueListModel())
	case FieldMsg:
		return core.NewQVariant1(f.Message())
	case FieldEnum:
		return core.NewQVariant1(f.EnumListModel())

	default:
		return core.NewQVariant()
	}
}

func (i *Message) rowCount(parent *core.QModelIndex) int {
	return len(i.Fields())
}

func (i *Message) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (i *Message) roleNames() map[int]*core.QByteArray {
	return i.Roles()
}

func (i *Message) updateFieldValue(row int, val string) {
	f := i.Fields()[row]
	if f.Value() == val {
		return
	}
	f.SetValue(val)
	var pIndex = i.Index(row, 0, core.NewQModelIndex())
	i.DataChanged(pIndex, pIndex, []int{FieldVal})
}
