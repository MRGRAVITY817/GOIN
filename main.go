package main

import (
	"fmt"
	"time"
)

// chan<- means it's send-only channel
func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("Sending %d\n", i)
		c <- i
	}
	close(c) // close channel after used
}

// <-chan means it's receive-only channel
// without arrow, you can send to channel
func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("Done!")
			break
		}
		fmt.Printf("Received %d\n\n", a)
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	receive(c)
}

// defer db.Close()
// cli.Start()
