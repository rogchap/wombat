// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/therecipe/qt/core"
)

type serviceList struct {
	core.QStringListModel

	_ func() `constructor:"init"`
}

func (s *serviceList) init() {
	s.ConnectData(s.data)
	s.ConnectRowCount(s.rowCount)
	// We set an initial value so that when the data is bound we don't have a ghost first value
	s.SetStringList([]string{" "})
}

func (s *serviceList) data(idx *core.QModelIndex, role int) *core.QVariant {
	if !idx.IsValid() {
		return core.NewQVariant()
	}

	if idx.Row() >= len(s.StringList()) {
		return core.NewQVariant()
	}

	return core.NewQVariant1(s.StringList()[idx.Row()])
}

func (s *serviceList) rowCount(parent *core.QModelIndex) int {
	return len(s.StringList())
}
