package main

import (
	"fmt"
	"os"

	"github.com/s-bespalov/gopl/ch7/Exercise-7.3/bytecounter"
)

func main() {
	w, n := bytecounter.CountingWriter(os.Stdout)
	fmt.Printf("Initial n: %d\n", *n)
	fmt.Fprintf(w, "Test string to stdoun\n")
	fmt.Printf("n after writing 22 runes: %d\n", *n)
}
