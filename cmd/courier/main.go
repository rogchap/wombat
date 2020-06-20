// Copyright 2020 Rogchap. All Rights Reserved.

package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"rogchap.com/courier/internal/component"
)

var PRODUCTION = false

func main() {
	core.QCoreApplication_SetApplicationName("Courier")
	core.QCoreApplication_SetOrganizationName("Rogchap")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	view := "qrc:/qml/main.qml"
	if !PRODUCTION {
		view = filepath.Join(".", "qml", "main.qml")
	}

	engine := qml.NewQQmlApplicationEngine(nil)
	eCtx := component.NewEngineCtx(nil)
	eCtx.Engine = engine
	eCtx.MainView = view
	engine.RootContext().SetContextProperty("eCtx", eCtx)

	engine.Load(core.NewQUrl3(view, 0))

	os.Exit(app.Exec())
}
