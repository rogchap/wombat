package app

import (
	"os"
	"path/filepath"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/cmd"
)

func appDataLocation(name string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	p := filepath.Join(homeDir, "Library", "Application Support", name)
	if wails.BuildMode == cmd.BuildModeBridge {
		p = ".data"
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0700); err != nil {
			return "", err
		}
	}
	return p, nil
}
