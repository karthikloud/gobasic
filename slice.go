package main

import "fmt"

func main() {
	slice := []int{2, 4, 7, 5, 8}

	num := slice[2:3]

	fmt.Println(slice)
	fmt.Println(num)

	slice = append(slice, 9, 10)
	fmt.Println(slice)
}
