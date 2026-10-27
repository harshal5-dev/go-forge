package main

import (
	"fmt"
)

func exercise3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("-------- Exercise3 ---------")
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
