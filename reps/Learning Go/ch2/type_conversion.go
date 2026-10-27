package main

import "fmt"

func typeConversion() {
	var (
		x int     = 10
		y float64 = 30.2
		b byte    = 100
	)

	var sum1 float64 = y + float64(x)
	var sum2 int = x + int(y)
	var sum3 int = x + int(b)
	var sum4 byte = b + byte(x)

	fmt.Println("---------- Type Conversion ---------")
	fmt.Println("Sum1: ", sum1)
	fmt.Println("Sum2: ", sum2)
	fmt.Println("Sum3: ", sum3)
	fmt.Println("Sum4: ", sum4)
}
