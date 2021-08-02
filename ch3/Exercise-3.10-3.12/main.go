package main

import (
	"bytes"
	"fmt"
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

func main() {
	fmt.Println(comma("3443624670920040000010010"), "comma")
	fmt.Println(commabuf("3443624670920040000010010"), "commabuf")
}
