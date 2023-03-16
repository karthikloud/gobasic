package main

import "fmt"

func main() {

	number := 7

	for i := 1; i <= 10; i++ {

		if i == number {

			fmt.Println(number, "is the lucky number!")
		} else if i%2 == 0 {

			fmt.Println(i, "is even")
		} else {

			fmt.Println(i, "is odd")
		}
	}
}
