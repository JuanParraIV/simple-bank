package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // [min, max]
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(5) + ".com"
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6)
}

// RandomHashedPassword generates a random hashed password
func RandomHashedPassword() string {
	return RandomString(8)
}

// RandomFullName generates a random full name
func RandomFullName() string {
	return RandomString(10)
}
