package main

import "fmt"

// Exercise-4.3
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Exercise-4.4
// Rotates a slice left by n elements
func rotate(s []int, n int) {
	n = n % len(s)
	buf := make([]int, n)
	for i := 0; i < len(s); i++ {
		if i < n {
			buf[i] = s[i]
		}
		if i < len(s)-n {
			s[i] = s[i+n]
		} else {
			s[i] = buf[i-(len(s)-n)]
		}
	}
}

// Exercise-4.5 duplicates function eliminates adjacent duplicates in a []string slice
func duplicates(s []string) []string {
	i, j := 0, 0
	for i < len(s) && j < len(s) {
		s[i] = s[j]
		i++
		for j < len(s) && s[j] == s[i-1] {
			j++
		}
	}
	return s[:i]
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println("reversed", a)
	b := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v rotated by %d :\n", b, 2)
	rotate(b[:], 2)
	fmt.Printf("%v\n", b)
	c := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v rotated by %d :\n", c, 10)
	rotate(c[:], 10)
	fmt.Printf("%v\n", c)

	// duplicates
	ds := []string{"milk", "coffee", "coffee", "coffee", "sugar", "sugar", "cacao", "tea", "tea", "tea"}
	fmt.Println("strings with duplicates:", ds)
	fmt.Println("duplicates remover", duplicates(ds))
}
