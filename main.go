package main

import "fmt"

func main() {
	a := make([]int, 10)
	a = append(a, 10)
	a = append(a, 20)

	b := "tttt"

	for i := range a {
		fmt.Printf("a[%02d] = %d, %s\n", i, a[i], b)
	}
}

func main2() {
	a := make([]int, 10)
	b := 10
	a = append(a, b)
	a = append(a, b)

	fmt.Printf("a = %d\n", a[0])
}
