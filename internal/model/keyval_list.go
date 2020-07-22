// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/therecipe/qt/core"
)

type keyval struct {
	key string
	val string
}

//go:generate qtmoc
type KeyvalList struct {
	core.QAbstractListModel

	list []keyval

	_ func() `constructor:"init"`

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

	kv := m.list[index.Row()]

	switch role {
	case int(core.Qt__DisplayRole):
		return core.NewQVariant1(kv.key)
	default:
		return core.NewQVariant1(kv.val)
	}
}

func (m *KeyvalList) rowCount(parent *core.QModelIndex) int {
	return len(m.list)
}

func (m *KeyvalList) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *KeyvalList) valAt(idx int) string {
	return m.list[idx].val
}
