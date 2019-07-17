package golib

import (
	"math/rand"
	"time"
)

// RandomChoice returns a random index of an array
func RandomChoice(arr []int) int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(len(arr))
	return arr[a]
}
