package pwordsencode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decode runs the encoding algorithm backwards and decodes the string
func Decode(enc, key string) (str string) {
	// divide the key to bytes
	divided := strings.Split(enc, "-")

	for i, d := range divided {
		ss, err := strconv.ParseUint(d, 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		str += decodeByte(ss, key, i)
	}

	return str
}

// decode a single byte of the encrypted string
func decodeByte(b uint64, key string, i int) string {

	sum := 0
	for _, s := range key {
		sum += int(s)
	}
	// fmt.Println("decoding>", sum)

	// iterate over the key
	for j, k := range key {
		b -= uint64(math.Pow(2.0, float64(k))) / uint64(sum+i+j)
	}
	b -= uint64(sum * len(key))
	b -= uint64(i * len(key))

	// if invalid character
	b %= 128
	if b < 33 && b != ' ' {
		return "lOl"
	}
	return string(byte(b))
}
