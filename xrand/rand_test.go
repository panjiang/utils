package xrand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIntnInRange(t *testing.T) {
	summary := map[int]int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100000; i++ {
		n := IntnInRange(0, 9)
		summary[n]++
	}

	fmt.Println(summary)
}
