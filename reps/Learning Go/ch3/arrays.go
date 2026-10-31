package main

import "fmt"

func arrays() {
	var x = []int{10, 20, 30}
	var y = [...]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(x)
	fmt.Println(y)
	x[0] = 40
	fmt.Println(x[0])
}
