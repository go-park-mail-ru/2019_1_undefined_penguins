package helpers

import "testing"

func TestPasswords(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	result := CheckPasswordHash(password, hash)
	if !result {
		t.Fatal("hashing is not equal")
	}

}
