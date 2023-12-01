package gcommon

import (
	"math/rand"
	"strconv"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var seededRand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// RandomAlphabeticString generates a random string of a given length from the defined charset
func RandomAlphabeticString(length int) string {
	ranString := make([]rune, length)
	for i := range ranString {
		ranString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(ranString)
}

// RandomStringFromCharset generates a random string of a given length from a custom charset
func RandomStringFromCharset(length int, charset []rune) string {
	ranString := make([]rune, length)
	for i := range ranString {
		ranString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(ranString)
}

// RandomNumericString generates a random numeric string of a given length
func RandomNumericString(length int) string {
	seededRandInt := seededRand.Int()

	minValue := 1
	for i := 1; i < length; i++ {
		minValue *= 10
	}
	maxValue := minValue * 10

	numString := strconv.Itoa(seededRandInt % maxValue)
	return numString
}
