package main

import (
	"fmt"
	"time"
)

// chan<- means it's send-only channel
func countToTen(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf("Sending %d\n", i)
		c <- i
	}
	close(c) // close channel after used
}

// <-chan means it's receive-only channel
// without arrow, you can send to channel
func receive(c <-chan int) {
	for {
		time.Sleep(1 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("Done!")
			break
		}
		fmt.Printf("Received %d\n\n", a)
	}
}

func main() {
	// when you make buffered channel, you can rapidly store sending data
	// buffer and then send.
	// The code below will stack first 5 integers in buffer and then send
	// to the channel.
	c := make(chan int, 5)
	go countToTen(c)
	receive(c)
}

// defer db.Close()
// cli.Start()
