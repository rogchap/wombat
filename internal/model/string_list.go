// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/therecipe/qt/core"
)

//go:generate qtmoc
type StringList struct {
	core.QStringListModel

	_ func() `constructor:"init"`

	_ func(string) `slot:"add"`
}

func (s *StringList) init() {
	s.ConnectData(s.data)
	s.ConnectRowCount(s.rowCount)
	s.ConnectAdd(s.add)
}

func (s *StringList) data(idx *core.QModelIndex, role int) *core.QVariant {
	if !idx.IsValid() {
		return core.NewQVariant()
	}

	if idx.Row() >= len(s.StringList()) {
		return core.NewQVariant()
	}

	return core.NewQVariant1(s.StringList()[idx.Row()])
}

func (s *StringList) rowCount(parent *core.QModelIndex) int {
	return len(s.StringList())
}

func (s *StringList) add(val string) {
	s.BeginInsertRows(core.NewQModelIndex(), len(s.StringList()), len(s.StringList()))
	s.SetStringList(append(s.StringList(), ""))
	s.EndInsertRows()
}
