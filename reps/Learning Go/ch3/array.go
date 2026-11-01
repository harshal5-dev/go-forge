package main

import "fmt"

func array() {
	var x = []int{10, 20, 30}
	var y = [...]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println("---------- Arrays --------------")
	fmt.Println("x => ", x)
	fmt.Println("y => ", y)
	x[0] = 40
	fmt.Println("x[0] => ", x[0])
	fmt.Println("length of y => ", len(y))
}
