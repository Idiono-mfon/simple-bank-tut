package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGH"

// Gnerate a random integer 4 between min and
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //random int64 between min and max
}

// Gnerate random string of n characters

func RandomString(n int) string {
	// Question: why do we have to use strings.Builder
	var sb strings.Builder

	k := len(alphabets)

	for i := 0; i < n; i++ {
		// rand.Intn(k) //Get a random number between 0 and k - 1
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate a randomOwner name

func RandOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
