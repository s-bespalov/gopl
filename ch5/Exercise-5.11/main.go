package main

import (
	"fmt"
	"log"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	position := 1
	var visitAll func(key string)
	visitAll = func(key string) {
		chain := make(map[string]bool)
		var visit func(key string)
		visit = func(key string) {
			if chain[key] {
				log.Println("--------------------")
				log.Println("Cycle detected:")
				for k := range chain {
					log.Printf("%s: %v\n", k, m[k])
				}
				log.Println("--------------------")
			} else {
				chain[key] = true
			}
			for _, item := range m[key] {
				if !seen[item] {
					seen[item] = true
					visit(item)
					order[position] = item
					position++
				}
			}
			if !seen[key] {
				seen[key] = true
				order[position] = key
				position++
			}
		}
		visit(key)
	}
	for k := range m {
		visitAll(k)
	}
	return order
}

func main() {
	for k, v := range topoSort(prereqs) {
		fmt.Printf("%d: %s\n", k, v)
	}
}
