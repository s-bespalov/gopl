package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// spacesquash squashes adjacent Unicode spaces in one ASCII space
func spacesquash(s []byte) []byte {
	k, i := 0, 0
	for i < len(s) {
		r, rl := utf8.DecodeRune(s[i:])
		if i > 0 && unicode.IsSpace(r) {
			if s[k-1] != '\x20' {
				s[k] = '\x20'
				k++
			}
		} else {
			for _, b := range s[i : i+rl] {
				s[k] = b
				k++
			}
		}
		i += rl
	}
	return s[:k]
}

func main() {
	data := "It\n is\t\u12e4  \v⽏\t⽹\v ⾧⾯a\vtest"
	fmt.Println(string(spacesquash([]byte(data))))
}
