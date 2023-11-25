package Jcrypto

import (
	"testing"
)

func TestHash(t *testing.T) {
	testString := "testing123"

	hash, err := Hash(testString)
	if err != nil {
		t.Fail()
	}

	if hash == "" {
		t.Fail()
	}
}

func TestCompareHash(t *testing.T) {
	testString := "testing123"

	hash, _ := Hash(testString)

	isValid := CompareHash(testString, hash)
	if !isValid {
		t.Fail()
	}

	isValid = CompareHash("wrongpassword", hash)
	if isValid {
		t.Fail()
	}
}
