package util

import "math/rand"

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"BDT", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
