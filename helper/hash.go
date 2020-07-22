package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

// CreateToken generate token
func CreateToken() string {
	return RandString(16)
}

// Hash val
func Hash(val string) string {
	h := sha256.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}
