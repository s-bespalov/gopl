package main

import (
	"fmt"
	"os"

	"github.com/s.bespalov/gopl/ch4/Exercise-4.9/wordfreq"
)

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModePerm)
	check(err)
	defer func() {
		err = f.Close()
		check(err)
	}()

	h := wordfreq.Wordferq(f)
	wordfreq.PrintFreq(os.Stdout, h)
}

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "wordferq:%v", e)
		panic(e)
	}
}
