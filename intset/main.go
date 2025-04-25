// An IntSet is a set of small non-negative integers. Its zero value represents the empty set

package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	bits = 32 << (^uint(0) >> 63) // Exercise 6.5
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Printf("bit = %v\n", (1 << bit))
	fmt.Printf("word=%d, bit=%d\n", word, bit)
	fmt.Printf("len word=%d, other=%d", len(s.words), s.words[word])
	return word < len(s.words) && s.words[word]&(1<<bit) != 0 // verschiebe bit um eine Stelle nach links
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)

	}
	s.words[word] |= 1 << bit // Bit OR
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the len of IntSet
func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9,42}

	x.UnionWith(&y)
	fmt.Println(x.String()) // { 1, 9, 42, 144}

	fmt.Printf("Elements of x = %d\n", x.Len())

	os.Exit(0)

}
