package random

import (
	"math/rand"

	"github.com/iancoleman/strcase"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// String return random string, Default length 5.
func String(length ...int) string {
	size := 5
	if len(length) > 0 {
		size = length[0]
	}
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] //nolint:gosec
	}
	return string(b)
}

// StringPascal return random string with pascal case, Default length 5.
func StringPascal(length ...int) string {
	size := 5
	if len(length) > 0 {
		size = length[0]
	}
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] //nolint:gosec
	}
	return strcase.ToCamel(string(b))
}
