package main

import "fmt"

func ping(pings chan<- string, msg string) {
	fmt.Println("ping:", msg)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("pong:", msg)
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println("main:", <-pongs)
}
