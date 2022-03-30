package intset

import (
	"fmt"
	"testing"
)

func TestAddAll(t *testing.T) {
	set := []int{13, 177, 5, 40}
	donotadd := []int{812, 17, 14, 15}
	intset := &IntSet{}
	intset.AddAll(set...)
	for _, item := range set {
		if !intset.Has(item) {
			t.Error("set does not contain", item)
		}
	}
	for _, item := range donotadd {
		if intset.Has(item) {
			t.Error("set contains extra", item)
		}
	}
}

func TestUnionWith(t *testing.T) {
	set1 := []int{5, 99}
	set2 := []int{10, 77, 88}
	intset := &IntSet{}
	intset.AddAll(set1...)
	intset2 := &IntSet{}
	intset2.AddAll(set2...)
	intset.UnionWith(intset2)
	set1 = append(set1, set2...)
	for _, item := range set1 {
		if !intset.Has(item) {
			t.FailNow()
		}
	}
}

func TestString(t *testing.T) {
	set := []int{11, 17, 88}
	str := "{11 17 88}"
	intset := &IntSet{}
	intset.AddAll(set...)
	if str != fmt.Sprint(intset) {
		t.Fail()
	}
}

func TestLen(t *testing.T) {
	set := []int{11, 17, 88}
	intset := &IntSet{}
	intset.AddAll(set...)
	if len(set) != intset.Len() {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	set := []int{11, 17, 88}
	intset := &IntSet{}
	intset.AddAll(set...)
	intset.Remove(set[1])
	if intset.Has(set[1]) {
		t.Fail()
	}
	if !intset.Has(set[0]) {
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	set := []int{11, 17, 88}
	intset := &IntSet{}
	intset.AddAll(set...)
	intset.Clear()
	if intset.Has(set[0]) {
		t.Fail()
	}
	if intset.Len() != 0 {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	set := []int{11, 17, 88}
	intset := &IntSet{}
	intset.AddAll(set...)
	intset2 := intset.Copy()
	if intset2 == nil {
		t.FailNow()
	}
	str1 := fmt.Sprint(intset)
	str2 := fmt.Sprint(intset2)
	if str1 != str2 {
		t.FailNow()
	}
	intset2.Remove(set[0])
	if !intset.Has(set[0]) {
		t.FailNow()
	}
}
