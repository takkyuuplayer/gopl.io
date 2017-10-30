package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("b"))

	fmt.Println(diff(c1, c2))
}

func diff(c1, c2 [32]byte) int {
	ret := 0

	for i := 0; i < 32; i++ {
		ret += popcount(c1[i] ^ c2[i])
	}

	return ret
}

func popcount(b byte) int {
	ret := 0

	for b > 0 {
		b = b & (b - 1)
		ret++
	}

	return ret
}
