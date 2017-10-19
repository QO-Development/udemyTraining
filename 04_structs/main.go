package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	/*
		//Three ways to define a struct
		//1. alex := person{"Alex", "Anderson"}
		//2. alex := person{firstName: "Alex", lastName: "Anderson"}

		//3.
		var alex person
		alex.firstName = "Alex"
		alex.lastName = "Anderson"

		fmt.Println(alex)
		fmt.Printf("%+v", alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")

	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
