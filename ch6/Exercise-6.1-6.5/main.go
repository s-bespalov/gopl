package main

import (
	"fmt"

	"github.com/s-bespalov/gopl/ch6/Exercise-6.1-6.5/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)

	fmt.Println("Exercise 6.1 test")
	fmt.Println(x.Len())

	x.Remove(144)
	fmt.Println(&x)
	fmt.Println(x.Len())

	c := x.Copy()

	x.Clear()
	fmt.Println(&x)
	fmt.Println(x.Len())

	fmt.Println(c)
	fmt.Println(c.Len())

	fmt.Println("Exercise 6.2 test")
	x.AddAll(14, 111, 93, 66)
	c.AddAll(666)
	fmt.Println(&x)
	fmt.Println(x.Len())

	fmt.Println(c)
	fmt.Println(c.Len())

	fmt.Println("Exercise 6.3 test")
	fmt.Println("Intersect With")
	c.AddAll(15, 115, 1115, 9999999)
	x.AddAll(15, 115, 1115)
	fmt.Println(&x)
	fmt.Println(c)
	s := x.IntersectWith(c)
	fmt.Println(s)
	s = x.SymmetricDifference(c)
	fmt.Println(s)
	s = x.DifferenceWith(c)
	fmt.Println(s)

}
