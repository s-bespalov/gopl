package main

import "fmt"

// rotate rotates a slice s left by n elements
func rotate(s []int, n int) []int {
	z := make([]int, len(s), cap(s))
	for i := range s {
		z[i] = s[(i+n)%len(s)]
	}
	return z
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 8))
}
