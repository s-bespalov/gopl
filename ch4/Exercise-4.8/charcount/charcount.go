package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	control = iota
	digit
	letter
	graphic
	mark
	number
	punct
	space
	symbol
)

var categoriesNames []string = []string{
	control: "Control",
	digit:   "Digit",
	letter:  "Letter",
	graphic: "Graphic",
	mark:    "Mark",
	number:  "Number",
	punct:   "Punct",
	space:   "Space",
	symbol:  "Symbol",
}

func categorize(r rune, c *[9]int) {
	if unicode.IsControl(r) {
		c[control]++
	}
	if unicode.IsDigit(r) {
		c[digit]++
	}
	if unicode.IsLetter(r) {
		c[letter]++
	}
	if unicode.IsGraphic(r) {
		c[graphic]++
	}
	if unicode.IsMark(r) {
		c[mark]++
	}
	if unicode.IsNumber(r) {
		c[number]++
	}
	if unicode.IsPunct(r) {
		c[punct]++
	}
	if unicode.IsSpace(r) {
		c[space]++
	}
	if unicode.IsSymbol(r) {
		c[symbol]++
	}
}

func PrintResults(w io.Writer, c *[9]int) {
	for i, v := range c {
		fmt.Fprintf(w, "%s:\t%d\n", categoriesNames[i], v)
	}
}

func Charcount(r io.Reader) [9]int {
	counts := [9]int{}
	invalid := 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		categorize(r, &counts)
	}
	return counts
}
