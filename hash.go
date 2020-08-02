package pwordsencode

import "fmt"

// Hash the given string in a unique way
func Hash(str string) (h string) {
	// sum of all the ascii values of all letters
	sum := 0

	for _, s := range str {
		sum += int(s)
	}

	var hash uint64 = uint64(len(str) * sum)

	for _, s := range str {
		n := hash * uint64(sum%int(s))
		h += fmt.Sprintf("%x", n)
		n = uint64(s)
		h += fmt.Sprintf("%x", n)
	}
	h += fmt.Sprintf("%x", sum*sum)
	return h
}
