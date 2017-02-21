package main

import (
	"./intset"
	"fmt"
)

func main() {
	var x, y intset.IntSet
	var z *intset.IntSet

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
	y.AddAll(1, 1, 2, 3, 5, 8, 13, 21, 2112)
	z.AddAll(1, 1, 2, 3, 5, 8)
	fmt.Println(z.String())
	//	y.IntersectWith(z)
	fmt.Println(y.String())

	fmt.Println("---------")
	fmt.Println(z.String())
	fmt.Println(y.String())
	y.SymmetricDifference(z)
	fmt.Println(y.String())

	fmt.Println(y.Elems())
	fmt.Println(32 << (^uint(0) >> 63))
}
