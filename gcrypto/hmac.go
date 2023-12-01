package gcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHMAC256 generates a HMAC (Hash-Based Message Authentication Code)
// using SHA-256, a type of public key cryptography that uses hash functions.
// The HMAC process mixes a secret key with the message data,
// hashes the result with the hash function, then mixes that hash value with the secret key again.
// This provides a fast and simple way of verifying the authenticity of information transmitted.
func GenerateHMAC256(data []byte, key []byte) string {
	hash := hmac.New(sha256.New, key)

	hash.Write(data)

	hashSum := hash.Sum(nil)

	hexHash := hex.EncodeToString(hashSum)

	return hexHash
}

// CompareHMAC256 computes a HMAC of the input data with provided key and checks
// if it matches the provided HMAC.
// This method is useful to authenticate that an API request is from a trusted source,
// as it can generate a HMAC from request data and a secret key during API communication.
func CompareHMAC256(computedHmac string, hmacToCheck string) bool {
	return hmac.Equal([]byte(computedHmac), []byte(hmacToCheck))
}
