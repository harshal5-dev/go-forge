package main

import (
	"fmt"
	"math/cmplx"
)

func complexNumber() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	var complexZeroVal complex128

	fmt.Println("-------- Complex Number --------")
	fmt.Println("Complex Literal Zero value:", complexZeroVal)
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)
	fmt.Println("Division:", x/y)
	fmt.Println("--------------------------------")
	fmt.Println("Real Portion:", real(x))
	fmt.Println("Imaginary Portion:", imag(x))
	fmt.Println("Absolute Value", cmplx.Abs(x))
	fmt.Println()
}
