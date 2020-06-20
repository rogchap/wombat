// Copyright 2020 Roger Chapman. All Rights Reserved.

package component

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type EngineCtx struct {
	core.QObject

	_ func() `signal:"reload,auto"`

	Engine   *qml.QQmlApplicationEngine
	MainView string
}

func (e *EngineCtx) reload() {
	e.Engine.ClearComponentCache()
	e.Engine.Load(core.NewQUrl3(e.MainView, 0))
}
