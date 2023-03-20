package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 1 - speed", i)
		time.Sleep(1 * time.Second)
	}
}

func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 2 - speed", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task1()
	go task2()

	time.Sleep(8 * time.Second)
}
