package main

import (
	"fmt"
	"reflect"
)

// Golang varibles
// A variable is a peice of storage

func main() {
	// integer type
	var x int = 10
	// float type
	var y float32 = 20.34
	// string type
	var z string = "new to golang"
	// bool type
	var isData = true

	fmt.Println(x, y, z, isData)

	// Declaration of multiple varibles
	a, b := 18.12, "tests"
	fmt.Println(a, b)
	var firstName, lastName string
	firstName = "FirstName"
	lastName = "secondName"
	fullName := firstName + " " + lastName
	fmt.Println(firstName, lastName)
	fmt.Println(fullName)

	// Check the type of x
	fmt.Println(reflect.TypeOf(x))

	// Declaring constants
	const PRODUCT = "phone"
	fmt.Println(PRODUCT)
}
