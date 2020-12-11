package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits.UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits.UintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Exercise 6.1

// Len returns count of values in set
func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits.UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				l++
			}
		}
	}
	return l
}

// Remove removes a value from set
func (s *IntSet) Remove(x int) {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= (1 << bit)
}

// Clear clears the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy copies the set
func (s *IntSet) Copy() *IntSet {
	var y IntSet
	y.words = make([]uint, len(s.words))
	copy(y.words, s.words)
	return &y
}

// Exercise 6.2

// AddAll ads all values to set
func (s *IntSet) AddAll(values ...int) {
	for _, x := range values {
		s.Add(x)
	}
}

// Exercise 6.3

// IntersectWith returns a set that is a result of intersection of two sets
func (s *IntSet) IntersectWith(s2 *IntSet) *IntSet {
	s3 := &IntSet{}
	for i, word := range s.words {
		if i >= len(s2.words) {
			break
		}
		word &= s2.words[i]
		s3.words = append(s3.words, word)
	}
	return s3
}

// SymmetricDifference returns a set that is a symmetric difference of two sets
func (s *IntSet) SymmetricDifference(s2 *IntSet) *IntSet {
	s3 := &IntSet{}
	for i, word := range s.words {
		if i < len(s2.words) {
			word ^= s2.words[i]
		}
		s3.words = append(s3.words, word)
	}

	if len(s2.words) > len(s.words) {
		l2 := len(s.words)
		s3.words = append(s3.words, s2.words[l2:]...)
	}
	return s3
}

// DifferenceWith returns a set that is a difference of two sets
func (s *IntSet) DifferenceWith(s2 *IntSet) *IntSet {
	s3 := &IntSet{}
	for i, word := range s.words {
		if i < len(s2.words) {
			word &^= s2.words[i]
		}
		s3.words = append(s3.words, word)
	}
	return s3
}

// Exercise 6.4

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() []int {
	var r []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits.UintSize; j++ {
			if word&(1<<j) != 0 {
				r = append(r, bits.UintSize*i+j)
			}
		}
	}
	return r
}
