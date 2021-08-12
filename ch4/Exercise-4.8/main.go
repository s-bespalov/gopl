package main

import (
	"os"

	"github.com/s.bespalov/gopl/ch4/Exercise-4.8/charcount"
)

func main() {
	hist := charcount.Charcount(os.Stdin)
	charcount.PrintResults(os.Stdout, &hist)
}
