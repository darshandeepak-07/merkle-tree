package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(data string) string {
	byteData := []byte(data)
	hash := sha256.Sum256(byteData)
	return hex.EncodeToString(reverse(hash[:]))
}

func reverse(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
