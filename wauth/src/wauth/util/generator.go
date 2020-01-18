package util

import (
	"math/rand"
	"time"
)

func GenerateDigitToken(len int) string {
	str := "1234567890"
	shuffleArray := make([]byte, len)

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		index := rand.Intn(10)
		shuffleArray[i] = str[index]
	}

	return string(shuffleArray)
}

func GenerateDigitToken32() string {
	return GenerateDigitToken(32)
}
