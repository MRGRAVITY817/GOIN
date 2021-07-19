package main

import (
	"fmt"
	"time"

	"github.com/MRGRAVITY817/goin/cli"
	"github.com/MRGRAVITY817/goin/db"
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
	defer db.Close()
	cli.Start()
}
