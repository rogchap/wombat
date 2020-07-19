// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/therecipe/qt/core"
)

const (
	RepeatedValueValue = int(core.Qt__UserRole) + 1<<iota
	RepeatedValueMsgValue
)

//go:generate qtmoc
type RepeatedValue struct {
	core.QObject

	_ string   `property:"value"`
	_ *Message `property:msgValue"`
}

//go:generate qtmoc
type RepeatedValues struct {
	core.QAbstractListModel

	ref *desc.MessageDescriptor

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*RepeatedValue         `property:"values"`
	_ int                      `property:"count"`

	_ func()    `slot:"addValue"`
	_ func(int) `slot:"remove"`
}

func (m *RepeatedValues) init() {
	m.SetRoles(map[int]*core.QByteArray{
		RepeatedValueValue:    core.NewQByteArray2("value", -1),
		RepeatedValueMsgValue: core.NewQByteArray2("msgValue", -1),
	})
	m.SetCount(0)

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddValue(m.addValue)
	m.ConnectRemove(m.remove)
}

func (m *RepeatedValues) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Values()) {
		return core.NewQVariant()
	}

	var v = m.Values()[index.Row()]

	switch role {
	case RepeatedValueValue:
		return core.NewQVariant1(v.Value())
	case RepeatedValueMsgValue:
		return core.NewQVariant1(v.MsgValue())

	default:
		return core.NewQVariant()
	}
}

func (m *RepeatedValues) rowCount(parent *core.QModelIndex) int {
	return len(m.Values())
}

func (m *RepeatedValues) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *RepeatedValues) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *RepeatedValues) addValue() {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Values()), len(m.Values()))
	rv := NewRepeatedValue(nil)
	if m.ref != nil {
		rv.SetMsgValue(MapMessage(m.ref))
	}
	m.SetValues(append(m.Values(), rv))
	m.EndInsertRows()
	m.SetCount(m.Count() + 1)
}

func (m *RepeatedValues) remove(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetValues(append(m.Values()[:row], m.Values()[row+1:]...))
	m.EndRemoveRows()
	m.SetCount(m.Count() - 1)
}
