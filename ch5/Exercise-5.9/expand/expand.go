package expand

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Expand(s string, f func(string) string) string {
	vals := map[string]string{}
	input := bufio.NewScanner(bytes.NewBufferString(s))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		if strings.HasPrefix(word, "$") {
			if rn, sz := utf8.DecodeLastRune([]byte(word)); unicode.IsPunct(rn) {
				word = word[:len(word)-sz]
			}
			vals[word] = ""
		}
	}
	if err := input.Err(); err != nil {
		return ""
	}
	for k := range vals {
		vals[k] = f(k[1:])
		s = strings.ReplaceAll(s, k, vals[k])
	}
	return s
}
