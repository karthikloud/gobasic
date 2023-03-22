package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	//  send messages to the channels
	go func() {
		for {
			ch1 <- "specific"
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			ch2 <- "unorder"
			time.Sleep(2 * time.Second)
		}
	}()

	//  receive messages from the channels
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}

	}

}
