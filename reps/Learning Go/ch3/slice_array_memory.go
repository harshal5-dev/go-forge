package main

import "fmt"

func sliceArrayMemory() {
	x := [4]int{5, 6, 7, 8}
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	// panicArray := [5]int(xSlice)
	xArrayPointer := (*[4]int)(xSlice)

	y := x[:2]
	z := x[2:]

	x[0] = 10
	xSlice[0] = 10

	fmt.Println("x => ", x)
	fmt.Println("y => ", y)
	fmt.Println("z => ", z)
	fmt.Println("xSlice => ", xSlice)
	fmt.Println("xArray => ", xArray)
	fmt.Println("smallArray => ", smallArray)
	// fmt.Println("panicArray => ", panicArray) panic: runtime error: cannot convert slice with length 4 to array
	// or pointer to array with length 5
	fmt.Println("xArrayPointer => ", xArrayPointer)
}
