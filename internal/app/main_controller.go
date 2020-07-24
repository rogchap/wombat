// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"github.com/therecipe/qt/core"

	"rogchap.com/wombat/internal/pb"
)

//go:generate qtmoc
type mainController struct {
	core.QObject

	pbSource pb.Source

	_ func() `constructor:"init"`

	_ *workspaceController `property:"workspaceCtrl"`

	_ string `property:"output"`
}

func (c *mainController) init() {
	c.SetWorkspaceCtrl(NewWorkspaceController(nil))
}
