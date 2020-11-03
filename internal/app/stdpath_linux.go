package app

import (
	"fmt"
	"os"
)

func appDataLocation(name string) (string, error) {
	p := fmt.Sprintf("~/.local/share/%s", name)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0700); err != nil {
			return "", err
		}
	}
	return p, nil
}
