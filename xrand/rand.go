package xrand

import (
	"math/rand"
)

// IntnInRange returns a random int value in a range
// min <= N <= max
func IntnInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
