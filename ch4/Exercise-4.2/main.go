package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var t = flag.String("t", "SHA256", "hash type")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		if *t == "SHA256" {
			c := sha256.Sum256([]byte(arg))
			fmt.Printf("%x\n", c)
		} else if *t == "SHA384" {
			c := sha512.Sum384([]byte(arg))
			fmt.Printf("%x\n", c)
		} else if *t == "SHA512" {
			c := sha512.Sum512([]byte(arg))
			fmt.Printf("%x\n", c)
		} else {
			fmt.Println("unsupported hash type")
		}
	}
}
