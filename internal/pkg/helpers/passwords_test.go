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
	// fmt.Println(hash)
	result := CheckPasswordHash(password, hash)
	if !result {
		t.Error("hashing is not equal")
	}

}
