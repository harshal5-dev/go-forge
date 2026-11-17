package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func constDeclaration() {
	const y = "hello"
	const a = 10
	var b int = a
	var c float64 = a
	var d byte = 10
	const typedX int = 10

	fmt.Println("---------- Const Declaration----------")
	fmt.Println("x => ", x)
	fmt.Println("y => ", y)

	// x = x + 1 this will not compile!
	// y = "bye" this will not compile!

	fmt.Println("b => ", b)
	fmt.Println("c => ", c)
	fmt.Println("d => ", d)
}
