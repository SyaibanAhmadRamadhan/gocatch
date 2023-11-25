package jolly

import (
	"math/rand"
	"strconv"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randTimeUnix = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[randTimeUnix.Intn(len(charset))]
	}
	return string(b)
}

func RandomStringWithCustomCharset(length int, charset []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[randTimeUnix.Intn(len(charset))]
	}
	return string(b)
}

func RandomInt(length int) string {
	randTimeUnixInt := randTimeUnix.Int()

	minValue := 1
	for i := 1; i < length; i++ {
		minValue *= 10
	}
	maxValue := minValue * 10

	sixDigitsString := strconv.Itoa(randTimeUnixInt % maxValue)
	return sixDigitsString
}
