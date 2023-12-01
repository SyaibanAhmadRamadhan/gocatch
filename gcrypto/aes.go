package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

// GenerateKey takes a string input and generates a 32 characters key.
// Whitespaces from the string are removed and if the resulting string length is less than 32,
// zeros are appended at the end until its length is exactly 32.
func GenerateKey(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	if len(str) < 32 {
		str += strings.Repeat("0", 32-len(str))
	}
	return str[:32]
}

// AesDecryptCFB uses Cipher Feedback Mode (CFB) decryption and decrypts the given encrypted text using the provided 32 characters key.
// It returns the decrypted text or an error if any occurred during the decryption process.
func AesDecryptCFB(text string, key string) (str string, err error) {
	if len(key) > 32 {
		err = errors.New("key len must be 32")
		return
	}

	cipherText, err := hex.DecodeString(text)
	if err != nil {
		return str, fmt.Errorf("failed decode hex string | err : %v | text : %s", err, text)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return str, fmt.Errorf("failed generate cipher block | err : %v | text : %s", err, text)
	}

	if len(cipherText) < aes.BlockSize {
		return str, fmt.Errorf("text is too short | text : %s", text)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	str = string(cipherText)
	return
}

// AesEncryptCFB uses Cipher Feedback Mode (CFB) for encryption and encrypts the given text with the provided 32 characters key.
// It returns the hex-encoded encrypted cipherText or an error.
func AesEncryptCFB(text string, key string) (str string, err error) {
	if len(key) > 32 {
		err = errors.New("key len must be 32")
		return
	}

	plainText := []byte(text)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return str, fmt.Errorf("failed generate cipher block | err : %v | text : %s", err, text)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return str, fmt.Errorf("failed generate vector | err : %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	str = fmt.Sprintf("%x", cipherText)
	return
}
