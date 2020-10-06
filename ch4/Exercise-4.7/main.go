package main

import (
	"fmt"
	"unicode/utf8"
)

// reverse reverses a byte slice that represents UTF-8 encoded string in place
func reverse(s []byte) {
	end := len(s)
	for end > 0 {
		r, rl := utf8.DecodeRune(s)
		copy(s, s[rl:end])
		rbuf := make([]byte, rl)
		utf8.EncodeRune(rbuf, r)
		end -= rl
		copy(s[end:], rbuf)
	}
}

func main() {
	data := []byte("It is \u12e4 ⽏⽹ a test ⾧⾯")
	fmt.Println(string(data))
	reverse(data)
	fmt.Println(string(data))
}
