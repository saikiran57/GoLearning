package main

import "fmt"

//Declare structs abc
type person struct {
	firstName string
	lastName  string
	age       int
}

func newPerson() (p *person) {
	p = &person{
		"ss",
		"kkk",
		33,
	}
	return p
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func main() {
	p := person{firstName: "sai", lastName: "kiran", age: 33}
	r := &p
	s := &r
	np := newPerson()
	fmt.Printf("%p\n", &p)
	fmt.Printf("%p\n", r)
	fmt.Printf("%p\n", *s)
	p.updateName("sss")
	p.print()
	r.updateName("saii")
	p.print()
	np.print()
}
