package main

import "fmt"

type varaible interface {
	int | int64 | float32 | string
}

func min[T varaible](x, y T) {

	fmt.Println(x + y)
}

func main() {
	a := 10
	b := 20
	min(a, b)

}
