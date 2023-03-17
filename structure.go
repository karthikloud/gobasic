package main

import "fmt"

//structure

type mobile struct {
	karthi int
	colour string
	price  float32
}
type cars struct {
	mobile
	name  string
	model string
	num   int
	price float64
}

func main() {
	a := mobile{karthi: 22, colour: "white", price: 12222}
	b := cars{name: "car", model: "benz", num: 0000, price: 5000000}
	c := cars{mobile: mobile{karthi: 25, colour: "black", price: 80000.9}}
	fmt.Printf("%v", a.colour)
	fmt.Printf("%v", b)
	fmt.Printf("%v", c)
}
