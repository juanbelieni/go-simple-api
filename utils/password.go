package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHashedPassword(ps string) string {
	h := sha256.New()
	h.Write([]byte(ps))
	return hex.EncodeToString(h.Sum(nil))
}
