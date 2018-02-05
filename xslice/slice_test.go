package xslice

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMix1(t *testing.T) {
	rand.Seed(1)

	n := 100000000
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}

	s = Mix(s)
	fmt.Println("Mix1:", len(s), s[0], s[len(s)-1])
}

// 0.1Y 1s
// 1Y 10s
func TestMix2(t *testing.T) {
	rand.Seed(1)

	n := 100000000
	s := rand.Perm(n)
	fmt.Println("Mix2:", len(s), s[0], s[len(s)-1])
}

func BenchmarkMix1(b *testing.B) {
	s := []int{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

	Mix(s)
}
func BenchmarkMix2(b *testing.B) {
	rand.Perm(b.N)
}

func BenchmarkRemove(b *testing.B) {
	s := []int{}
	for i := 0; i <= 100000; i++ {
		s = append(s, i)
	}

	for i := range s {
		Romove(s, i)
	}
}
