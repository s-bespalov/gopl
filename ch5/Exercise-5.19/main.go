package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type result struct {
	upper string
}

func upper(s string) {
	r := strings.ToUpper(s)
	panic(result{r})
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			if reflect.TypeOf(p) == reflect.TypeOf(result{}) {
				r := p.(result)
				fmt.Printf("result: %s\n", r.upper)
			} else {
				panic(p)
			}
		}
	}()
	upper(os.Args[1])
}
