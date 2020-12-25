package main

import (
	"fmt"

	"github.com/s-bespalov/gopl/ch7/Exercise-7.3/treesort"
)

func main() {
	t := []int{5, 6, 1, 8, 17, 11, 4}
	fmt.Println(t)
	treesort.Sort(t)
	fmt.Println(t)

}
