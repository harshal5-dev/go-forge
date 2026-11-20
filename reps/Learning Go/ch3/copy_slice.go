package main

import "fmt"

func copySlice() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	z := make([]int, 2)
	a := make([]int, 2)
	d := [4]int{5, 6, 7, 8}
	e := make([]int, 2)

	num := copy(y, x)
	numZ := copy(z, x)
	numA := copy(a, x[2:])
	numE := copy(e, d[:])
	copy(d[:], x)

	fmt.Println("y => ", y, "num of element copied => ", num)
	fmt.Println("z => ", z, "num of element copied => ", numZ)
	fmt.Println("a => ", a, "num of element copied => ", numA)
	fmt.Println("e => ", e, "num of element copied => ", numE)
	fmt.Println("d => ", d)
}
