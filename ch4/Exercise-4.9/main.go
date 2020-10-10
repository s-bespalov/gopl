package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(bufio.NewReader(os.Stdin))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standart input:", err)
		os.Exit(1)
	}
	fmt.Printf("word\tcount\n")
	for w, n := range counts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
