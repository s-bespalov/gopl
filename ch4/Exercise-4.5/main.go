package main

import "fmt"

// rmdup eliminates adjacent duplicates
func rmdup(s []string) []string {
	i := 1
	for _, v := range s {
		if s[i-1] != v {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func main() {
	data := []string{"two", "two", "one", "one", "one", "three", "two", "two"}
	fmt.Println(rmdup(data))
}
