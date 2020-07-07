// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"rogchap.com/courier/internal/debug"
)

// The following variables are set via LDFlags at build time
var (
	semver  = "master"
	isDebug = true
)

// Startup is the main startup of the application
func Startup() int {
	core.QCoreApplication_SetApplicationName("Courier")
	core.QCoreApplication_SetOrganizationName("Rogchap")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	engine := qml.NewQQmlApplicationEngine(nil)

	entry := "qrc:/qml/main.qml"
	if isDebug {
		entry = filepath.Join(".", "qml", "main.qml")
		debug.HotReloader(engine)
		app.SetQuitOnLastWindowClosed(false)
	}

	mc := NewMainController(nil)

	engine.RootContext().SetContextProperty("mc", mc)
	engine.Load(core.NewQUrl3(entry, 0))

	return app.Exec()
}
