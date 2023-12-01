package gcrypto

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash function accepts a string password, generates a hashed password from it using bcrypt, and then returns the hashed password as a string.
func Hash(str string) (hash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(str), 14)

	return string(hashByte), err
}

// CompareHash function accepts a string password and a hashed password, compares the password against the hashed password, and then returns true if they match, false otherwise.
func CompareHash(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
