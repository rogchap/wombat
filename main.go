package main

import (
	"embed"
	_ "embed"
	"os"

	"wombat/internal/app"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	os.Exit(app.Run(assets))
}
