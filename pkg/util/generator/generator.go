package generator

import (
	"crypto/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateUniqueString(length int) string {
	timestamp := []byte(time.Now().Format("20060102150405"))

	result := make([]byte, 0, length)

	result = append(result, timestamp...)

	numRandomChars := length - len(result)
	if numRandomChars <= 0 {
		return string(result[:length])
	}

	randomBytes := make([]byte, numRandomChars)
	rand.Read(randomBytes)
	for i := 0; i < numRandomChars; i++ {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}

	result = append(result, randomBytes...)

	return string(result)
}
