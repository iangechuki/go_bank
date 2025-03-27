package util

import (
	"math/rand"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func RandInt(min, max int) int64 {
	return int64(rnd.Intn(max-min+1) + min)
}
func RandomString(n int) string {
	k := len(alphabet)
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rnd.Intn(k)]
	}
	return string(b)
}

// generate a random owner
func RandomOwner() string {
	return RandomString(6)
}

// generates a random amount of money
func RandomMoney() int64 {
	return int64(RandInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rnd.Intn(n)]
}
