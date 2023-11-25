package Jcrypto

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

func GenerateKey(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	if len(str) < 32 {
		str += strings.Repeat("0", 32-len(str))
	}
	return str[:32]
}

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
