package main

import (
	"fmt"
	"time"
)

func digitsSum(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}

	return sum
}

// create worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for v := range jobs {
		fmt.Println("woker id:", id, "started job:", v)
		time.Sleep(time.Second)
		fmt.Println("worker id:", id, "finished job:", v)
		results <- v * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
