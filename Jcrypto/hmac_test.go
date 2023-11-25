package Jcrypto

import (
	"testing"
)

func TestHmacFunctions(t *testing.T) {
	// Define secret key and data
	key := []byte("secret")
	data := []byte("my confidential data")

	// Compute HMAC
	hash := HMAC256(data, key)

	// Validate it mismatches with incorrect HMAC
	wrong := "wrong_hash"
	if CheckHMAC256(hash, wrong) {
		t.Fail()
	}

	// Validate it matches with correct HMAC
	if !CheckHMAC256(hash, hash) {
		t.Fail()
	}
}
