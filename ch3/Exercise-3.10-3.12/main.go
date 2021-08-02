package main

import (
	"bytes"
	"fmt"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Exercise 3.10
func commabuf(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	i := 0
	c := len(s) % 3
	if c == 0 {
		c = 3
	}
	var buf bytes.Buffer
	for i+c < len(s) {
		buf.WriteString(s[i : i+c])
		buf.WriteByte(',')
		i += c
		c = 3
	}
	buf.WriteString(s[i:])
	return buf.String()
}

// Exercise 3.11
func commasf(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}
	var suf string
	if dot := strings.LastIndexByte(s, '.'); dot >= 0 {
		suf = s[dot:]
	}
	s = s[:len(s)-len(suf)]
	i := 0
	c := len(s) % 3
	if c == 0 {
		c = 3
	}
	for i+c < len(s) {
		buf.WriteString(s[i : i+c])
		buf.WriteByte(',')
		i += c
		c = 3
	}
	buf.WriteString(s[i:])
	buf.WriteString(suf)
	return buf.String()
}

func anograms(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	h := make(map[rune]int)
	for _, r := range s1 {
		h[r]++
	}
	for _, r := range s2 {
		h[r]--
		if h[r] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(comma("3443624670920040000010010"), "comma")
	fmt.Println(commabuf("3443624670920040000010010"), "commabuf")
	fmt.Println(commasf("-3443624670920040000010010.9999"), "commasf")
	fmt.Println(commasf("3443624670920040000010010"), "commasf")
	fmt.Println(commasf("-0.139846251897"), "commasf")
	fmt.Println("is anograms:", "tirades", "aridest", "-", anograms("tirades", "aridest"))
	fmt.Println("is anograms:", "колба", "бокал", "-", anograms("колба", "бокал"))
	fmt.Println("is anograms:", "колбочка", "бокалчик", "-", anograms("колбочка", "бокалчик"))
	fmt.Println("is anograms:", "情人", "人情", "-", anograms("情人", "人情"))
}
