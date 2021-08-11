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
	uncategorized
)

var categoriesNames []string = []string{
	control:       "Control",
	digit:         "Digit",
	letter:        "Letter",
	graphic:       "Graphic",
	mark:          "Mark",
	number:        "Number",
	punct:         "Punct",
	space:         "Space",
	symbol:        "Symbol",
	uncategorized: "Uncategoriced",
}

func charcount() {
	counts := [10]int{}
	invalid := 0
	uncategorized := 0

	in := bufio.NewReader(os.Stdin)
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
	}
}
