package main

import "fmt"

func slicingSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("--------- Slicing Slices ----------")
	fmt.Println("x => ", x)
	fmt.Println("y => ", y)
	fmt.Println("z => ", z)
	fmt.Println("d => ", d)
	fmt.Println("e => ", e)

	fmt.Println("------ Slice Share Storage -------")
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x => ", x)
	fmt.Println("y => ", y)
	fmt.Println("z => ", z)

	fmt.Println("------ Slice append Storage -------")
	f := []string{"a", "b", "c", "d"}
	h := f[:2]
	fmt.Println("cap(f) => ", cap(f), "cap(h) => ", cap(h))
	h = append(h, "z")
	fmt.Println("f => ", f)
	fmt.Println("h => ", h)

	fmt.Println("-------- Confusing Slices --------")
	i := make([]string, 0, 5)
	i = append(i, "a", "b", "c", "d")
	j := i[:2]
	k := i[2:]
	fmt.Println("cap(i) => ", cap(i), "cap(j) => ", cap(j), "cap(k) => ", cap(k))
	j = append(j, "i", "j", "k")
	i = append(i, "x")
	k = append(k, "y")
	fmt.Println("i => ", i)
	fmt.Println("j => ", j)
	fmt.Println("k => ", k)

	fmt.Println("------ Full Slice Expression ------")
	l := make([]string, 0, 5)
	l = append(l, "a", "b", "c", "d")
	m := l[:2:2]
	n := l[2:4:4]
	fmt.Println("cap(l) => ", cap(l), "cap(m) => ", cap(m), "cap(n) => ", cap(n))
	fmt.Println("l => ", l)
	fmt.Println("m => ", m)
	fmt.Println("n => ", n)
}
