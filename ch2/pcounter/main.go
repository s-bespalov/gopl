package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/s-bespalov/gopl/ch2/popcount"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		num, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pcounter: %v\n", err)
			os.Exit(1)
		}
		p1 := popcount.PopCount(num)
		p2 := popcount.PCLoop(num)
		p3 := popcount.PCShift(num)
		p4 := popcount.PCMask(num)
		fmt.Printf("%d: PopCount = %d, PCLoop = %d, PCShift = %d, PCMask = %d\n", num, p1, p2, p3, p4)
	}
}
