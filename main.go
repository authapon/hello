package main

import "fmt"

func main() {
	a := make([]int, 10)
	a = append(a, 10)
	a = append(a, 20)

	for i := range a {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

func main2() {
	a := make([]int, 10)
	b := 10
	a = append(a, b)
	a = append(a, b)

	fmt.Printf("a = %d\n", a[0])
}
