package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/s-bespalov/gopl/ch2/Exercise-2.3-2.5/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseUint(arg, 10, 64)
		check(err)
		fmt.Println(popcount.PopCount(n), popcount.PopCountLoop(n), popcount.PopCountShift(n), popcount.PopCountClear(n))
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
