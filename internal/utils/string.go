package utils

import (
	"math/rand"
	"time"
)

const (
	LChar   = "abcdefghijklmnopqrstuvwxyz"
	UChar   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NChar   = "0123456789"
	charset = LChar + UChar + NChar
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func MakeStr(n int, c string) string {
	chars := c
	if c == "" {
		chars = charset
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(chars))]
	}
	return string(b)
}
