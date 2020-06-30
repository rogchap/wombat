// Copyright 2020 Rogchap. All Rights Reserved.

package model

import "github.com/therecipe/qt/core"

const (
	FieldType = int(core.Qt__UserRole) + 1<<iota
	FieldLabel
	FieldVal
)

//go:generate qtmoc
type Field struct {
	core.QObject

	_ string `property:"type"`
	_ string `property:"label"`
	_ int    `property:"tag"`
	_ string `property:"value"`
}

//go:generate qtmoc
type Input struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ string                   `property:"label"`
	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Field                 `property:"fields"`

	_ func(row int, val string) `slot:"updateFieldValue"`
}

func (i *Input) init() {
	i.SetRoles(map[int]*core.QByteArray{
		FieldType:  core.NewQByteArray2("type", -1),
		FieldLabel: core.NewQByteArray2("label", -1),
		FieldVal:   core.NewQByteArray2("val", -1),
	})

	i.ConnectData(i.data)
	i.ConnectRowCount(i.rowCount)
	i.ConnectColumnCount(i.columnCount)
	i.ConnectRoleNames(i.roleNames)

	i.ConnectUpdateFieldValue(i.updateFieldValue)
}

func (i *Input) data(index *core.QModelIndex, role int) *core.QVariant {
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
	case FieldLabel:
		return core.NewQVariant1(f.Label())
	case FieldVal:
		return core.NewQVariant1(f.Value())

	default:
		return core.NewQVariant()
	}
}

func (i *Input) rowCount(parent *core.QModelIndex) int {
	return len(i.Fields())
}

func (i *Input) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (i *Input) roleNames() map[int]*core.QByteArray {
	return i.Roles()
}

func (i *Input) updateFieldValue(row int, val string) {
	f := i.Fields()[row]
	if f.Value() == val {
		return
	}
	f.SetValue(val)
	var pIndex = i.Index(row, 0, core.NewQModelIndex())
	i.DataChanged(pIndex, pIndex, []int{FieldVal})
}
