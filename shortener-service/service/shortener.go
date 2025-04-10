package service

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var randsequqnce = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randsequqnce.Intn(len(charset))]
	}
	return string(b)
}
