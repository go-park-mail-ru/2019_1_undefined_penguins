package helpers

import (
	"testing"
)

func TestPasswords(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}
	result := CheckPasswordHash(password, hash)
	if !result {
		t.Error("hashing is not equal")
	}

}

func TestLogs(t *testing.T) {
	LogMsg("hello world")
}
