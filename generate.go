package pwordsencode

import "math/rand"

// RandomString creates a string of random characters with the given length
func RandomString(l int) (gibberish string) {

	for len(gibberish) != l {
		n := rand.Int() % 128

		if (n > 32 && n < 91) || (n > 97 && n < 126) {
			gibberish += string(byte(n))
		}
	}
	return gibberish
}
