package main

import "fmt"

func join(sep string, elements ...string) (r string, err error) {
	if len(elements) == 0 {
		err = fmt.Errorf("no incoming arguments in func join(string, ...string)")
		return
	}
	for i, e := range elements {
		r += e
		if i != len(elements)-1 {
			r += sep
		}
	}
	return
}

func main() {
	fmt.Println(join("-", "test", "test2", "abc"))
	fmt.Println(join(" "))
}
