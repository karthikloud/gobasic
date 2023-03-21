package main

import "fmt"

func main() {
	net := make(chan int)

	go count(net)
	for u := range net {
		fmt.Println(u)
	}

}
func count(in chan int) {
	for i := 1; i <= 6; i++ {
		in <- i
	}
	close(in)
}
