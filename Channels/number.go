package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var c chan bool

func even(count int) {
	defer wg.Done()
	for i := 0; i < count; i += 2 {
		<-c
		fmt.Println("even:", i)
		c <- true
	}
	fmt.Println("closing chan")
	close(c)
}

func odd(count int) {
	defer wg.Done()
	for i := 1; i < count; i += 2 {
		fmt.Println("odd:", i)
		_, ok := <-c
		if ok {
			fmt.Println("Received from ch2:")
			c <- true
		} else {
			fmt.Println("ch2 closed")
		}
	}
}

func main() {
	c = make(chan bool)
	wg.Add(2)
	go even(10)
	c <- true
	go odd(10)
	wg.Wait()
}
