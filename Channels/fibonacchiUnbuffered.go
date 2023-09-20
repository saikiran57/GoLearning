package main

import "fmt"

func fib(num int, val chan<- int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		val <- x
		x, y = y, x+y
	}

	close(val)
}

func main() {
	ch := make(chan int)

	go fib(10, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
