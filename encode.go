package whistle

import (
	"fmt"
	"math"
)

// Encode the string with the given key
func Encode(str, key string) (enc string) {

	for i, s := range str {
		ss := byte(s)
		enc += string(EncodeByte(ss, key, i))
	}

	return enc[:len(enc)-1]
}

// EncodeByte with the given key and index of the
// byte in the text
func EncodeByte(b byte, key string, i int) string {
	var temp uint64 = uint64(b)

	// sum of all key letters
	sum := 0
	for _, s := range key {
		sum += int(s)
	}

	// iterate over the key
	for j, k := range key {

		temp += uint64(math.Pow(2.0, float64(k))) / uint64(sum+i+j)
	}

	temp += uint64(sum * len(key))
	temp += uint64(i * len(key))

	// return temp in hexadecimal
	return fmt.Sprintf("%x-", temp)
}
