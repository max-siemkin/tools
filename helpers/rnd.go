package tools

import (
	"math/rand"
	"strings"
	"time"
)

func newRnd() *rand.Rand { return rand.New(rand.NewSource(time.Now().UnixNano())) }

// RndHash -
func RndHash(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rnd := newRnd()
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}

// RndNumbers -
func RndNumbers(n int) string {
	letterRunes := []rune("0123456789")
	b := make([]rune, n)
	rnd := newRnd()
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}

// RndLetters -
func RndLetters(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rnd := newRnd()
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}

// RndPasswSpecChar -
func RndPasswSpecChar(n int) string {
	for {
		letterRunes := []rune("abcdefghijklmnopqrstuvwxyz0123456789-!@#$%&()_ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, n)
		rnd := newRnd()
		for i := range b {
			b[i] = letterRunes[rnd.Intn(len(letterRunes))]
		}
		for _, c := range "-!@#$%&()_" {
			if strings.Contains(string(b), string(c)) {
				return string(b)
			}
		}
	}
}
