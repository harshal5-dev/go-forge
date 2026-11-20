package main

import (
	"fmt"
	"maps"
)

func mapReadWrite() {
	totalWins := map[string]int{}
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	m1 := map[string]int{
		"hello": 5,
		"world": 0,
	}
	n := map[string]int{
		"hello": 5,
		"world": 0,
	}
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println("Orcas => ", totalWins["Orcas"])
	fmt.Println("Kittens => ", totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println("Kittens => ", totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println("Lions => ", totalWins["Lions"])

	v, ok := m["hello"]
	fmt.Println("v => ", v, "ok => ", ok)

	v, ok = m["world"]
	fmt.Println("v => ", v, "ok => ", ok)

	v, ok = m["goodbye"]
	fmt.Println("v => ", v, "ok => ", ok)

	fmt.Println("m => ", m)
	delete(m, "hello")
	fmt.Println("m => ", m)
	clear(m)
	fmt.Println("m => ", m)

	fmt.Println("maps.Equal(m,n) => ", maps.Equal(m, n))
	fmt.Println("maps.Equal(m1,n) => ", maps.Equal(m1, n))

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println("len(vals) => ", len(vals), "len(intSet) => ", len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
