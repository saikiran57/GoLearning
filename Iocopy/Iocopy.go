package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://google.com"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
