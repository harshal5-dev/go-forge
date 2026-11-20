package main

import "fmt"

func stringSlicing() {
	// var s string = "Hello there"
	var s string = "Hello, 🌞"
	var b byte = s[6]
	var s2 string = s[4:8]
	var s3 string = s[:5]
	var s4 string = s[6:]
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	var x rune = 'x'
	var y byte = 'y'
	var s5 string = string(x)
	var s6 string = string(y)

	fmt.Println("s => ", s)
	fmt.Println("len(s) => ", len(s))
	fmt.Println("b => ", b)
	fmt.Println("s2 => ", s2)
	fmt.Println("s3 => ", s3)
	fmt.Println("s4 => ", s4)
	fmt.Println("bs => ", bs)
	fmt.Println("rs => ", rs)
	fmt.Println("s5 => ", s5)
	fmt.Println("s6 => ", s6)
}
