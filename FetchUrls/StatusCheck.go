package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
	}

	fmt.Println("In main")
	var wg sync.WaitGroup

	c := make(chan string)
	done := make(chan struct{})

	// receiver
	go receiver(c, done)

	// sender
	for _, link := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			statusCheck(url, c)
		}(link)
	}

	fmt.Println("after go")
	wg.Wait()
	close(c)
	<-done

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

}

func statusCheck(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "failed")
		c <- link
		return
	}
	fmt.Println(link, "up")
	c <- link
}

func receiver(c chan string, done chan struct{}) {
	for l := range c {
		fmt.Println(l)
	}
	close(done)
}
