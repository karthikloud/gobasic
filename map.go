package main

import "fmt"

//multiple return value
func swap(x, y int) (int, int) {
	return y, x


// closure
func adder() func(int) int {
    Sum := 0
    return func(x int) int {
        Sum += x
        return Sum
    }
}
// recursion
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
func main() {
//
	x, y := 1, 2
	fmt.Printf("Before swap:", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swap: ", x, y)
}

//
	a := adder()
    for i := 1; i <= 5; i++ {
        fmt.Printf("Sum after adding :", i, a(i))
    }
	result := factorial(5)
    fmt.Println(result)
}

 }
