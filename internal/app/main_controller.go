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

	_ string               `property:"version"`
	_ *workspaceController `property:"workspaceCtrl"`
}

func (c *mainController) init() {
	c.SetWorkspaceCtrl(NewWorkspaceController(nil))
	c.SetVersion(semver)
}
