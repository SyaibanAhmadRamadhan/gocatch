package Jcrypto

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) (hash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(str), 14)

	return string(hashByte), err
}

func CompareHash(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
