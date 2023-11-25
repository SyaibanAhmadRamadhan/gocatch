package Jcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HMAC256(data []byte, key []byte) string {
	hash := hmac.New(sha256.New, key)

	hash.Write(data)

	hashSum := hash.Sum(nil)

	hexHash := hex.EncodeToString(hashSum)

	return hexHash
}

func CheckHMAC256(computedHmac string, hmacToCheck string) bool {
	return hmac.Equal([]byte(computedHmac), []byte(hmacToCheck))
}
