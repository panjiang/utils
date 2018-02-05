package xslice

import "math/rand"

// Contains whether slice exists certain value
func Contains(s []string, a string) bool {
	for _, b := range s {
		if b == a {
			return true
		}
	}
	return false
}

// Romove remove a element certain index
// method will be low performance along with elements growth
func Romove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// Mix mix a slice by random algorithm
func Mix(m []int) []int {
	for i := range m {
		j := rand.Intn(i + 1)
		m[i], m[j] = m[j], m[i]
	}
	return m

}
