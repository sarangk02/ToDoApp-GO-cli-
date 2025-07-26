package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(10000)
}
