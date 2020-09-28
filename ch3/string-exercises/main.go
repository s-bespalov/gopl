package main

import (
	"bytes"
	"log"
	"strings"
)

// insert commas in non negative, decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	if strings.HasPrefix(s, "-") {
		buf.WriteRune('-')
		s = s[1:]
	}
	var suf string
	if strings.Contains(s, ".") {
		p := strings.LastIndex(s, ".")
		suf = s[p+1:]
		s = s[:p]
	}
	for i := len(s) % 3; i <= len(s); i += 3 {
		if i > 3 {
			buf.WriteRune(',')
			buf.WriteString(s[i-3 : i])
		} else {
			buf.WriteString(s[:i])
		}
	}
	if len(suf) > 0 {
		buf.WriteRune('.')
		buf.WriteString(suf)
	}
	return string(buf.Bytes())
}

// reports whether two strings are anagrams of each other
func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hist := make(map[rune]int)
	histogram(hist, s1)
	for _, r := range s2 {
		hist[r]--
		if hist[r] < 0 {
			return false
		}
	}
	return true
}

func histogram(hist map[rune]int, s string) {
	for _, r := range s {
		hist[r]++
	}
}

func main() {
	log.Println(comma("12345"))
	log.Println(comma("123456789"))
	log.Println(comma("1234567"))
	log.Println(comma2("12345"))
	log.Println(comma2("123456789"))
	log.Println(comma2("123456.7"))
	log.Println(comma2("12345.6789"))
	log.Println(comma2("-12345.67"))
	log.Println(comma2("-1234567"))
	log.Println(isAnagrams("tirades", "aridest"))
	log.Println(isAnagrams("tirades", "aridist"))
	log.Println(isAnagrams("колба", "бокал"))
	log.Println(isAnagrams("колбочка", "бокалчик"))
	log.Println(isAnagrams("情人", "人情"))
}
