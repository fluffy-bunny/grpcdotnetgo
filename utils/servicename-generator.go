package utils

import (
	"fmt"
	"math/rand"
)

func GenerateUnqueServiceName(prepend string) string {
	if len(prepend) == 0 {
		return RandomString(32)
	}
	return fmt.Sprintf("%v.%v", prepend, RandomString(32))
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
