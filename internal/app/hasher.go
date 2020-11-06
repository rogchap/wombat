package app

import (
	"crypto/sha1"
	"encoding/hex"
)

func hash(vals ...string) string {
	h := sha1.New()
	for _, v := range vals {
		h.Write([]byte(v))
	}
	return hex.EncodeToString(h.Sum(nil))
}
