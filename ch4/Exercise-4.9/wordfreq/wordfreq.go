package wordfreq

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func trimNotLetter(w string) string {
	r1, l1 := utf8.DecodeRuneInString(w)
	for !unicode.IsLetter(r1) && len(w) > 0 {
		w = w[l1:]
		r1, l1 = utf8.DecodeRuneInString(w)
	}
	r2, l2 := utf8.DecodeLastRuneInString(w)
	for !unicode.IsLetter(r2) && len(w) > 0 {
		w = w[:len(w)-l2]
		r2, l2 = utf8.DecodeLastRuneInString(w)
	}
	return w
}

func Wordferq(r io.Reader) map[string]int {
	hist := make(map[string]int)

	input := bufio.NewScanner(bufio.NewReader(r))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		w := strings.ToLower(input.Text())
		w = trimNotLetter(w)
		if len(w) > 0 {
			hist[w]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq:%v", err)
	}
	return hist
}

func PrintFreq(w io.Writer, hist map[string]int) {
	r := make(map[int]string) // reversed
	for k, v := range hist {
		if r[v] == "" {
			r[v] = k
		} else {
			r[v] = fmt.Sprintf("%s, %s", r[v], k)
		}
	}
	keys := make([]int, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, v := range keys {
		fmt.Fprintf(w, "%s:\t%d\n", r[v], v)
	}
}
