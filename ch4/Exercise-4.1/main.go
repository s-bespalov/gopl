package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Diff(c1, c2 [32]uint8) int {
	var sum int
	for i := range c1 {
		sum += int(pc[c1[i]^c2[i]])
	}
	return sum
}

func PrintBinSha256(f io.Writer, c [32]uint8) {
	for _, b := range c {
		fmt.Fprintf(f, "%08b", b)
	}
}

// alternative diff for check
func AltDiff(c1, c2 string) int {
	var sum int
	for i := range c1 {
		if c1[i] != c2[i] {
			sum++
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enought arguments")
	}
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("Count of different bytes in sha256 of \"%s\" and \"%s\" is %d\n", os.Args[1], os.Args[2], Diff(c1, c2))
	var buf1, buf2 bytes.Buffer
	PrintBinSha256(&buf1, c1)
	PrintBinSha256(&buf2, c2)
	fmt.Printf("Alternative Diff count: %d\n", AltDiff(buf1.String(), buf2.String()))
	fmt.Printf("sha256 of %q in binary form:\n", os.Args[1])
	fmt.Println(buf1.String())
	fmt.Printf("sha256 of %q in binary form:\n", os.Args[2])
	fmt.Println(buf2.String())
}
