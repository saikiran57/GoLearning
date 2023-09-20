package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", x)
			} else {
				fmt.Println("ch1 closed")
				ch1 = nil
			}
		case x, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", x)
			} else {
				fmt.Println("ch2 closed")
				ch2 = nil
			}
		default:
			// select a new channel to receive from
			if ch1 == nil && ch2 == nil {
				// both channels closed, exit the loop
				return
			}
		}
	}
}
