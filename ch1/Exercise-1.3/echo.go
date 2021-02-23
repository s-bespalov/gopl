package main

import (
	"fmt"
	"os"
	"strings"
)

func echoJoin() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoFor() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func main() {
}
