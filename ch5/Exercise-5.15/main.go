package main

import "fmt"

func max(vals ...int) (max int, err error) {
	if len(vals) <= 0 {
		err = fmt.Errorf("no incoming arguments in func max()")
		return
	}
	max = vals[0]
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return
}

func min(vals ...int) (min int, err error) {
	if len(vals) <= 0 {
		err = fmt.Errorf("no incoming arguments in func min()")
		return
	}
	min = vals[0]
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return
}

func main() {
	fmt.Println(max(8, 9, 10, 2, -2, 1))
	fmt.Println(min(8, 9, 10, 2, -2, 1))
}
