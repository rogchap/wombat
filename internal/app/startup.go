// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"os"
	"path/filepath"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"rogchap.com/courier/internal/debug"
	"rogchap.com/courier/internal/model"
	"rogchap.com/courier/internal/pb"
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
	}

	sl := model.NewServiceList(nil)

	println("main thread:", core.QThread_CurrentThread().Pointer())

	// TODO [RC] This is just sample code to test, needs moving
	go func() {
		time.Sleep(1 * time.Second)
		println("some other thread:", core.QThread_CurrentThread().Pointer())
		source, err := pb.GetSourceFromProtoFiles(nil, []string{"./example/route_guide.proto"})
		if err != nil {
			println(err.Error())
			return
		}
		services := source.ListServices()

		MainThread.Run(func() {
			// fmt.Printf("%+v\n", t)
			sl.SetStringList(services)
			println("again on main thread:", core.QThread_CurrentThread().Pointer())
		})

	}()

	// go func() {
	// 	source, _ := pb.GetSourceFromProtoFiles(nil, []string{"../../example/route_guide.proto"})
	// 	go mainThread.RunOnMain(func() {
	//
	// 		sl.SetServices(source.ListServices())
	// 	})
	// }()

	engine.RootContext().SetContextProperty("serviceList", sl)
	engine.Load(core.NewQUrl3(entry, 0))

	return app.Exec()
}
