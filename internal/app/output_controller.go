// Copyright 2020 Rogchap. All Rights Reserved.

package app

import "github.com/therecipe/qt/core"

//go:generate qtmoc
type outputController struct {
	core.QObject

	_ int32  `property:"status"`
	_ string `property:"output"` // Temp property, should be a list model

	_ func() `constructor:"init"`
}

func (c *outputController) init() {
	c.SetStatus(-1)
}

func (c *outputController) clear() {
	c.SetOutput("")
	c.SetStatus(-1)
}
