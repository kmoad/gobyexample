package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
	counter := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func(n int) { counter <- n }(i)
	}
	fmt.Println(<-counter)
	fmt.Println(<-counter)
	fmt.Println(<-counter)
	fmt.Println(<-counter)
	fmt.Println(<-counter)
	// fmt.Println(<-counter)
}
