package main

import "fmt"

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func printArea(s Shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	square := Square{side: 5.8}
	printArea(square)
}
