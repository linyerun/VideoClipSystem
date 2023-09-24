package utils

import "crypto/rand"

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateAuthCode(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		Logger().Error(err)
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
