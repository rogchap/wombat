package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/cmd"

	"wombat/internal/server"
)

// Run is the main function to run the application
func Run() int {

	if wails.BuildMode != cmd.BuildModeProd {
		go server.Serve()
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Resizable: true,
		Title:     "Wombat",
		JS:        mewn.String("./frontend/public/build/bundle.js"),
		CSS:       mewn.String("./frontend/public/build/bundle.css"),
		Colour:    "#2e3440",
	})

	app.Bind(&api{})

	if err := app.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "app: error running app: %v\n", err)
		return 1
	}
	return 0
}
