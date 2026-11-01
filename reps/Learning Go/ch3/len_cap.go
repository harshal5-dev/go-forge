package main

import "fmt"

func lenCap() {
	var x []int
	fmt.Println("--------- Length and Capacity ----------")
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
	x = append(x, 10)
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
	x = append(x, 20)
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
	x = append(x, 30)
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
	x = append(x, 40)
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
	x = append(x, 50)
	fmt.Println("x => ", x, "len => ", len(x), "cap => ", cap(x))
}
