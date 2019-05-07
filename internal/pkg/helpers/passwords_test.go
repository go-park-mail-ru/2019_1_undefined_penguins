package helpers

import (
	"testing"
)

func TestPasswords(t *testing.T) {
	password := "password"
	hash := HashPassword(password)
	result := CheckPasswordHash(password, hash)
	if !result {
		t.Error("hashing is not equal")
	}

}

func TestLogs(t *testing.T) {
	LogMsg("hello world")
}
