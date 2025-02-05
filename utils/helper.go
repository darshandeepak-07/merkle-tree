package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(data string) string {
	byteData := []byte(data)
	hash := sha256.Sum256(byteData)
	return hex.EncodeToString(hash[:])
}
