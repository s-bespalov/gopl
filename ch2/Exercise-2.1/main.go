package main

import (
	"fmt"

	"github.com/s-bespalov/gopl/ch2/Exercise-2.1/tempconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("10 C to F= %v\n", tempconv.CToF(10))
	fmt.Printf("12C = %v\n", tempconv.CToK(12))
	fmt.Printf("71F = %v\n", tempconv.FToK(71))
}
