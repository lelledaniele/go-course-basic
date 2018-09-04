package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Bruno"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Bruno"
	alex.contact = contactInfo{email: "email", zip: 999}

	alex.updateName("Ciao")

	//fmt.Println(alex)
	alex.print()
}

func (p *person) updateName(n string) {
	(p).firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) toString() string {
	return p.firstName
}
