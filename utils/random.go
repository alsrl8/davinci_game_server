package utils

import (
	"math/rand"
	"time"
)

func MakeRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
