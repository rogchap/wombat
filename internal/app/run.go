package app

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v2"
	woptions "github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var (
	appname = "Wombat"
	semver  = "0.0.0-dev"
)

// Run is the main function to run the application
func Run(assets embed.FS) int {
	appData, err := appDataLocation(appname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open add data directory: %v\n", err)
		return 1
	}
	defer crashlog(appData)

	appApi := &api{appData: appData}

	cfg := &woptions.App{
		Title:  appname,
		Width:  1200,
		Height: 820,
		BackgroundColour: &woptions.RGBA{
			R: 0x2e,
			G: 0x34,
			B: 0x40,
			A: 0xff,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			appApi,
		},
		OnStartup:  appApi.startup,
		OnDomReady: appApi.wailsReady,
	}

	if err := wails.Run(cfg); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "app: error running app: %v\n", err)
		return 1
	}
	return 0
}

func crashlog(appData string) {
	//if wails.BuildMode != cmd.BuildModeProd {
	//	return
	//}
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
