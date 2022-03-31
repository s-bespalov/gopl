package intset

import (
	"bytes"
	"fmt"
)

// Exercise 6.5
var uintSize int

func init() {
	uintSize = 32 << (^uint(0) >> 63)
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Exercise 6.1
// return the number of elements
func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				l++
			}
		}
	}
	return l
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/uintSize, uint(x%uintSize)
	s.words[word] &= ^(1 << bit)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	t := IntSet{}
	t.words = append(t.words, s.words...)
	return &t
}

// Exercise 6.2
func (s *IntSet) AddAll(items ...int) {
	for _, item := range items {
		s.Add(item)
	}
}

// Exercise 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i := range s.words {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}
		s.words[i] &= ^t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}
		s.words[i] ^= t.words[i]
	}
	if len(s.words) < len(t.words) {
		s.words = append(s.words, t.words[len(s.words):]...)
	}
}

// Exercise 6.4
func (s *IntSet) Elems() []int {
	r := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, i*uintSize+j)
			}
		}
	}
	return r
}
