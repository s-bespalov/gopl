package toposort

import "testing"

var prereqs2 = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},
	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func TestTopoSortMaps(t *testing.T) {
	if result := TopoSortMaps(prereqs2); result == nil {
		t.Error("The result of TopoSort is nil")
	} else if len(result) < len(prereqs2) {
		t.Errorf("After sorting some values were lost, should be %d or more but have %d", len(prereqs2), len(result))
	} else {
		seen := make(map[string]bool)
		for _, c := range result {
			for rqs := range prereqs2[c] {
				if !seen[rqs] {
					t.Errorf("Conditions are not met, topic \"%s\", requires \"%s\"", c, rqs)
				}
			}
			seen[c] = true
		}
	}
}
