package main

import "fmt"

func main() {
	// Numeric values
	var i int = 10
	var f float64 = 3.14
	var c complex128 = complex(1, 2)
	fmt.Println(i, f, c)

	// String values
	var s string = "karthi"
	fmt.Println(s)

	// Boolean values
	var b bool = true
	fmt.Println(b)

	// Array values
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// Function values
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))
}
