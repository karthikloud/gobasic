package main

import "fmt"

func main() {

	var num [5]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	num[3] = 4
	num[4] = 5

	fmt.Println(num)

	name := [3]string{"siva", "sanjay", "balaji"}

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Change the value
	name[1] = "karthi"
	fmt.Println(name)

}
