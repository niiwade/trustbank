package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// generates a random integer between min and max
func RandomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate random owner
func RandomOwner() string {
	return RandomString(6)
}

// generate random amount of money
func RandomAmount() int64 {
	return RandomInit(0, 1000)

}

// generate random currency
func RandomCurrency() string {

	currencies := []string{"EUR", "GHS", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
