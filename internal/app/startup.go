// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"rogchap.com/wombat/internal/debug"
)

// The following variables are set via LDFlags at build time
var (
	appname = "Wombat"
	semver  = "0.1.0-alpha.1"
	isDebug = false
)

// Startup is the main startup of the application
func Startup() int {
	core.QCoreApplication_SetApplicationName(appname)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetWindowIcon(gui.NewQIcon5(":/qml/img/icon_128x128@2x.png"))

	engine := qml.NewQQmlApplicationEngine(nil)

	entry := "qrc:/qml/main.qml"
	appData := core.QStandardPaths_WritableLocation(core.QStandardPaths__AppDataLocation)
	if isDebug {
		entry = filepath.Join(".", "qml", "main.qml")
		appData = filepath.Join(".", ".data")

		debug.HotReloader(engine)
		app.SetQuitOnLastWindowClosed(false)
	}
	defer crashlog(appData)

	mc := NewMainController(nil)

	engine.RootContext().SetContextProperty("mc", mc)
	engine.Load(core.NewQUrl3(entry, 0))

	return app.Exec()
}

func crashlog(appData string) {
	if isDebug {
		return
	}
	if r := recover(); r != nil {
		if _, err := os.Stat(appData); os.IsNotExist(err) {
			os.MkdirAll(appData, 0700)
		}
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("%+v\n\n", r))
		buf := make([]byte, 1<<20)
		s := runtime.Stack(buf, true)
		b.Write(buf[0:s])
		ioutil.WriteFile(filepath.Join(appData, "crash.log"), b.Bytes(), 0644)
	}
}
