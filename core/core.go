package core

import (
	"math/rand"
	"time"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fileNameLength = 10
)

func GenerateRandomFileName() string {
	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(
		fileNameLength,
	) + 1
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
