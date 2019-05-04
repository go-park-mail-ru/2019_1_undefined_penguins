package helpers

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string {
	hasher := sha256.New224()
	hasher.Write([]byte(password))
	result := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return string(result)
}

func CheckPasswordHash(password, hash string) bool {
	return HashPassword(password) == hash
}
