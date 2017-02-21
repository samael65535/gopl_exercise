package main

/*
练习6.1: 为bit数组实现下面这些方法

func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/
import (
	"bytes"
	"fmt"
)

var pc []uint64

func init() {
	var i uint64
	pc = make([]uint64, 64, 64)
	for i = 0; i < 64; i++ {
		pc[i] = pc[i/2] + (i % 2)
	}
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
func (s *IntSet) Len() int { // return the number of elements
	count := 0
	for _, x := range s.words {
		//fmt.Println(word)
		c := int(pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
		count += c
	}
	return count
}

func (s *IntSet) Remove(x int) { // remove x from the set
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] ^= 1 << bit
	}

}

func (s *IntSet) Clear() { // remove all elements from the set
	for k := range s.words {
		s.words[k] = 0
	}
}

func (s *IntSet) Copy() *IntSet { // return a copy of the set
	var res IntSet
	for _, v := range s.words {
		res.words = append(res.words, v)
	}
	return &res
}

func main() {
	var x, y IntSet
	var z *IntSet

	x.Add(64)
	x.Add(2)
	z = x.Copy()
	fmt.Println(x.String())
	fmt.Println("---------")
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	fmt.Println("---------")
	x.UnionWith(&y)
	fmt.Println(x.String())
	x.Clear()
	x.Remove(64)
	x.Remove(1)
	fmt.Println(x.String())
	fmt.Println(z.String())
	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(x.Len())

	//for k, v := range pc {
	//	fmt.Println(k, v)
	//}
}
