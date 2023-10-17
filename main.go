package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

type superman struct {
	person person
	ltk    bool
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

func (sm superman) speak() {
	fmt.Println("I'm a superman - this is my name", sm.person.first)
}

func main() {

	p1 := person{
		first: "hide",
	}

	sm1 := superman{
		person: person{
			first: "HIDE",
		},
		ltk: true,
	}

	var h1, h2 human
	h1 = p1
	h1.speak()

	h2 = sm1
	h2.speak()

}
