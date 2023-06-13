package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabetic = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomString Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetic)
	for i := 0; i < n; i++ {
		c := alphabetic[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney Generates a random amount of money
func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

// RandomCurrency Generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// Nombres y apellidos comunes para generar nombres realistas.
var firstNames = []string{"John", "Jane", "Mike", "Emily", "Alex", "Sophia", "Jacob", "Olivia", "Mason", "Ava", "Ethan", "Isabella"}
var lastNames = []string{"Smith", "Johnson", "Brown", "Taylor", "Miller", "Wilson", "Moore", "Anderson", "Thomas", "Jackson", "White", "Harris"}

// RandomRealName generates a realistic random name
func RandomRealName() string {
	first := firstNames[rand.Intn(len(firstNames))]
	last := lastNames[rand.Intn(len(lastNames))]
	return first + " " + last
}
