package util

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

// Generate a random integer between min to max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate a random string of length n
func RandomString(n int) string {
	var substr strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		substr.WriteByte(c)
	}

	return substr.String()

}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
