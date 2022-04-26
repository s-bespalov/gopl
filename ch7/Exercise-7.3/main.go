package main

import (
	"com/github/s-bespalov/gopl/ch7/Exercise-7.3/treesort"
	"fmt"
)

func main() {
	arr := []int{14, 7, 18, 20, 83, 1, 3, 61, 2, 99, 17, 81, 66, 92, 0}
	t := treesort.Sort(arr)
	fmt.Print(t)
}
