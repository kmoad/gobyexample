package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	// Below is added by me later.
	// Took until closing channels lesson to
	// learn how to do this.
	counter := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			counter <- i
		}
		close(counter)
	}()
	for {
		count, valid := <-counter
		if valid {
			fmt.Println("count:", count)
		} else {
			break
		}
	}

	// Better method is to use for range
	counter = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			counter <- i
		}
		close(counter)
	}()
	for count := range counter {
		fmt.Println("count:", count)
	}
}
