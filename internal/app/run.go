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

	css := mewn.String("./frontend/public/build/bundle.css")
	js := mewn.String("./frontend/public/build/bundle.js")

	cfg := &wails.AppConfig{
		Width:     1200,
		Height:    820,
		Resizable: true,
		Title:     "Wombat",
		JS:        js,
		CSS:       css,
		Colour:    "#2e3440",
	}
	app := wails.CreateApp(cfg)

	app.Bind(&api{})

	if err := app.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "app: error running app: %v\n", err)
		return 1
	}
	return 0
}
