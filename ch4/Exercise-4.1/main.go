package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Diff returns number of bits that are different in two SHA256 hashes
func Diff(sha1, sha2 *[32]byte) int {
	sum := 0
	for i := range sha1 {
		sum += int(pc[sha1[i]^sha2[i]])
	}
	return sum
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	c3 := sha256.Sum256([]byte("x"))
	c4 := sha256.Sum256([]byte("as"))
	fmt.Printf("c1 diff c2: %d\n", Diff(&c1, &c2))
	fmt.Printf("c1 diff c3: %d\n", Diff(&c1, &c3))
	fmt.Printf("c2 diff c3: %d\n", Diff(&c2, &c3))
	fmt.Printf("c4 diff c1: %d\n", Diff(&c4, &c1))
}
