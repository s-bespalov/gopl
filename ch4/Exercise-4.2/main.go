package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var hashType string

func init() {
	flag.StringVar(&hashType, "h", "sha256", "hash type, suported sha256, sha384, sha512")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	data, err := io.ReadAll(os.Stdin)
	check(err)
	//fmt.Printf("%08b\n", data)
	switch strings.ToLower(hashType) {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(data))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(data))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(data))
	}
}
