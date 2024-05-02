package main

import "fmt"

func main() {
	a := make([]int, 10)
	a = append(a, 10)
	a = append(a, 20)
	a[20] = 30
	for i := range a {
		fmt.Printf("%d\n", i)
	}
}

func main2() {
	a := make([]int, 10)
	b := 10
	a = append(a, b)
	a = append(a, b)

	fmt.Printf("a = %d\n", a[0])
}
