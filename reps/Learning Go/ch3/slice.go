package main

import (
	"fmt"
	"slices"
)

func slice() {
	var x = []int{10, 20, 30}
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	z := []int{10, 20, 30}
	s := []string{"a,", "b", "c"}
	a := make([]int, 5)
	b := make([]int, 5, 10)
	c := make([]int, 0, 10)
	var d []int
	var e = []int{}

	fmt.Println("--------- Slices --------------")
	fmt.Println("s => ", s, "len => ", len(s))
	clear(s)
	fmt.Println("s => ", s, "len =>", len(s))
	y[1] = 2
	fmt.Println("y[1] => ", y[1])
	fmt.Println("Is x and z equals => ", slices.Equal(x, z))
	fmt.Println("Is y and z equals => ", slices.Equal(y, z))
	// fmt.Println("Is x and z equals => ", slices.Equal(x, s)) does not compile
	x = append(x, 40)
	x = append(x, 50, 60, 70)
	x = append(x, y...)
	fmt.Println("x => ", x)
	c = append(c, 5, 6, 7, 8)
	fmt.Println("a => ", a, "len => ", len(a), "cap => ", cap(a))
	fmt.Println("b => ", b, "len => ", len(b), "cap => ", cap(b))
	fmt.Println("c => ", c, "len => ", len(c), "cap => ", cap(c))
	fmt.Println("d => ", d, "len => ", len(d), "cap => ", cap(d))
	fmt.Println("e => ", e, "len => ", len(e), "cap => ", cap(e))
	fmt.Println("Is d equal to nil =>", d == nil)
	fmt.Println("Is e equal to nil =>", e == nil)
}
