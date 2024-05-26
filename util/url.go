package util

import "math/rand"

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
)

// GenerateShortUrl generates a random string with the size of six characters
// character set is lowercase and uppercase alphabets and numbers
func GenerateShortUrl() string {
	return generateRandomString(6, lowercase+uppercase+numbers)
}

// generateRandomString generates a random string with the given length
func generateRandomString(length int, characterSet string) string {
	charSetLength := len(characterSet)

	shortUrl := make([]uint8, length)
	for i := 0; i < 6; i++ {
		shortUrl[i] = characterSet[rand.Int31n(int32(charSetLength))]
	}

	return string(shortUrl)
}
