package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

func main() {

	p1 := person{
		first: "hide",
	}

	var h human
	h = p1
	h.speak()
}
