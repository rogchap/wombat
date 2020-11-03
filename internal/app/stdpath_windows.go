package app

import (
	"fmt"
	"os"
	"os/user"
)

func appDataLocation(name string) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	p := fmt.Sprintf("C:/Users/%s/AppData/Roaming/%s", u.Username, name)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0700); err != nil {
			return "", err
		}
	}
	return p, nil
}
