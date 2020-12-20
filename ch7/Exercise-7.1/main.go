package main

import (
	"fmt"

	"github.com/s-bespalov/gopl/ch7/Exercise-7.1/wordcounter"
)

func main() {
	var w wordcounter.WordCounter
	w.Write([]byte("Hello"))
	fmt.Println(w)

	w = 0
	name := "Dolly"
	fmt.Fprintf(&w, "hello, %s, test", name)
	fmt.Println(w)
}
