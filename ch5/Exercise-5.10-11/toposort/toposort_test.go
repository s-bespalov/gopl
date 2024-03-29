package toposort

import "testing"

func TestTopoSortMaps(t *testing.T) {
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
	if result, err := TopoSortMaps(prereqs2); err != nil {
		t.Error(err)
	} else if result == nil {
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

func TestTopoSortMapsLoops(t *testing.T) {
	var prereqs = map[string]map[string]bool{
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
		"linear algebra":        {"calculus": true},
	}
	expected := "found a circle"
	if _, err := TopoSortMaps(prereqs); err == nil {
		t.Error("Should return error if prerequisites have loops")
	} else if err.Error() != expected {
		t.Errorf("Expected an error \"%s\", received \"%s\"", expected, err.Error())
	}
}
