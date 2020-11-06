package app

import (
	"os"
	"path/filepath"
)

func appDataLocation(name string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(homeDir, "AppData", "Roaming", name)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0700); err != nil {
			return "", err
		}
	}
	return p, nil
}
