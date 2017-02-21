package intset

import (
	"bytes"
	"fmt"
)

var pc []uint

func init() {
	var i uint
	pc = make([]uint, wordSize, wordSize)
	for i = 0; i < wordSize; i++ {
		pc[i] = pc[i/2] + (i % 2)
	}
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
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
		for j := 0; j < wordSize; j++ {
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

/*
练习6.1: 为bit数组实现下面这些方法

func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/
func (s *IntSet) Len() uint { // return the number of elements
	var count uint
	for _, x := range s.words {
		//fmt.Println(word)
		var c uint
		for i := 0; i < wordSize; i++ {
			c += pc[byte(x>>uint(i*8))]
		}
		count += c
	}
	return count
}

func (s *IntSet) Remove(x int) { // remove x from the set
	if s.Has(x) {
		word, bit := x/wordSize, uint(x%wordSize)
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

/*
练习 6.2：
定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以为一组IntSet值求和(翻译问题吧?)，比如s.AddAll(1,2,3)。
*/

func (s *IntSet)AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

/*
练习 6.3：
(*IntSet).UnionWith会用|操作符计算两个集合的并集，
我们再为IntSet实现另外的几个函数
IntersectWith(交集：元素在A集合B集合均出现),  A & B
DifferenceWith(差集：元素出现在A集合，未出现在B集合), A - B = A ^ (A & B)
SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)。  (A-B) | (B-A)
*/

func (s *IntSet) IntersectWith(t *IntSet) {
	l := len(t.words)
	for i,_ := range s.words {
		if i < l {
			s.words[i] = s.words[i] & t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	l := len(t.words)
	for i,_ := range s.words {
		if i < l {
			s.words[i] = s.words[i] ^ (s.words[i] & t.words[i])
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	l := len(t.words)
	for i,_ := range s.words {
		if i < l {
			s.words[i] = (s.words[i] ^ (s.words[i] & t.words[i])) | (t.words[i] ^ (s.words[i] & t.words[i]))
		}
	}

}

/*
练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
*/
func (s *IntSet) Elems() []uint64{
	res := []uint64{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, uint64(64*i+j))
			}
		}
	}
	return res
}


/*
练习 6.5： 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，
可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
*/

 const wordSize = 32 << (^uint(0) >> 63)
