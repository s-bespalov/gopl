package main

import "fmt"

type mypanic struct{}

func testPanic() (result int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			//
		case mypanic{}:
			result = 10
		default:
			panic(p)
		}
	}()
	panic(mypanic{})
}

func main() {
	fmt.Println(testPanic())
}
