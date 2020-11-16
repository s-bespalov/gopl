package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}

func valuate(key string) (r string) {
	if key == "foo" {
		r = "lalala"
		return
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var result string
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v", err)
			os.Exit(1)
		}
		result += s
	}
	result = expand(result, valuate)
	fmt.Println(result)
}
