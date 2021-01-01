package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/cmd"

	"wombat/internal/server"
)

var (
	appname = "Wombat"
	semver  = "0.0.0-dev"
)

// Run is the main function to run the application
func Run() int {
	appData, err := appDataLocation(appname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open add data directory: %v\n", err)
		return 1
	}
	defer crashlog(appData)

	if wails.BuildMode != cmd.BuildModeProd {
		go server.Serve()
	}

	extra := mewn.String("./frontend/public/build/extra.css")
	css := mewn.String("./frontend/public/build/bundle.css")
	js := mewn.String("./frontend/public/build/bundle.js")

	cfg := &wails.AppConfig{
		Width:     1200,
		Height:    820,
		Resizable: true,
		Title:     appname,
		JS:        js,
		CSS:       extra + css,
		Colour:    "#2e3440",
	}

	app := wails.CreateApp(cfg)
	app.Bind(&api{appData: appData})

	if err := app.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "app: error running app: %v\n", err)
		return 1
	}
	return 0
}

func crashlog(appData string) {
	if wails.BuildMode != cmd.BuildModeProd {
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
