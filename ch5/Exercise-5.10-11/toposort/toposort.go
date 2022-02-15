package toposort

import (
	"fmt"
	"sort"
)

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func TopoSortMaps(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var loop map[string]bool
	var visitAll func(items map[string]bool) error
	visitAll = func(items map[string]bool) error {
		for item := range items {
			if loop[item] {
				err := fmt.Errorf("found a circle")
				return err
			} else {
				loop[item] = true
			}
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			}
		}
		return nil
	}
	for item, reqs := range m {
		if !seen[item] {
			loop = map[string]bool{item: true}
			if err := visitAll(reqs); err != nil {
				return nil, err
			}
			order = append(order, item)
			seen[item] = true
		}
	}
	return order, nil
}
