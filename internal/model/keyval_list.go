// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/therecipe/qt/core"
)

//go:generate qtmoc
type Keyval struct {
	core.QObject

	_ string `property:"key"`
	_ string `property:"val"`
}

//go:generate qtmoc
type KeyvalList struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ []*Keyval `property:"list"`

	_ func(int) string `slot:"valAt"`
}

func (m *KeyvalList) init() {
	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectValAt(m.valAt)

}

func (m *KeyvalList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	kv := m.List()[index.Row()]

	switch role {
	case int(core.Qt__DisplayRole):
		return core.NewQVariant1(kv.Key())
	default:
		return core.NewQVariant1(kv)
	}
}

func (m *KeyvalList) rowCount(parent *core.QModelIndex) int {
	return len(m.List())
}

func (m *KeyvalList) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *KeyvalList) valAt(idx int) string {
	return m.List()[idx].Val()
}
