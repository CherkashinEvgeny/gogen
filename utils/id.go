package utils

import (
	"math/rand"
	"strconv"
)

func InIds(n int) []string {
	return IdsSequence("arg", n)
}

func OutIds(n int) []string {
	return IdsSequence("res", n)
}

func IdsSequence(prefix string, n int) []string {
	names := make([]string, 0, n)
	for i := 0; i < n; i++ {
		names = append(names, prefix+strconv.Itoa(i))
	}
	return names
}

func RandomId(prefix string) string {
	return RandomIdN(prefix, 8)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomIdN(prefix string, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return prefix + string(b)
}
