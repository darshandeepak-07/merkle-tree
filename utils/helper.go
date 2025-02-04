package utils

import "crypto/sha256"

func GenerateHash(data string) string {
	byteData := []byte(data)
	hash := sha256.Sum256(byteData)
	return string(hash[:])
}
