package app

import (
	"fmt"
	"os"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/cmd"
)

func appDataLocation(name string) (string, error) {
	p := fmt.Sprintf("~/Library/Application Support/%s", name)
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
