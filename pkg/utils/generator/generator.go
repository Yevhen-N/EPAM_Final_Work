package generator

import (
	"math/rand"
	"time"
)

var numberRuns = []rune("0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeCardNumber() string {
	return generateString(16)
}

func MakeIBan() string {
	return "UA" + generateString(29)
}

func generateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRuns[rand.Intn(len(numberRuns))]
	}
	return string(b)
}
