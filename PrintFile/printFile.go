package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Invlaid lenght of args ")
		os.Exit(1)
	}

	fileName := os.Args[1]
	bs, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, bs)
}
