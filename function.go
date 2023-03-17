package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func sub(a string, b int) string {
	return a
}
func mul(k string, l string) string {

	return k + l
}

func main() {
	result := add(3, 5)
	fmt.Println(result)

	finish := sub("karthi", 5)
	fmt.Println(finish)

	submit := mul("siva", "balaji")
	fmt.Println(submit)
}
