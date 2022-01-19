package testutil

import (
	"math/rand"
	"time"
)

// RandomString generates a random string with given length.
// Notice: this function uses a pseudo random algorithm, only for use in test.
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
