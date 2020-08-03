//Dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files
//prints list of files where each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	infiles := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesFiles(f, counts, infiles, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fNames := ""
			sep := ""
			for fName := range infiles[line] {
				fNames += sep + fName
				sep = " "
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, fNames)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func countLinesFiles(f *os.File, counts map[string]int, infiles map[string]map[string]int, fname string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		//init nested map if needed
		if infiles[input.Text()] == nil {
			infiles[input.Text()] = make(map[string]int)
		}

		//put file name to set
		infiles[input.Text()][fname]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
