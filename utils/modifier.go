package utils

import (
	"strings"
)

func Public(id string) string {
	runes := []rune(id)
	firstLetter := strings.ToUpper(string(runes[0]))
	runes[0] = []rune(firstLetter)[0]
	return string(runes)
}

func Private(id string) string {
	runes := []rune(id)
	firstLetter := strings.ToLower(string(runes[0]))
	runes[0] = []rune(firstLetter)[0]
	return string(runes)
}
