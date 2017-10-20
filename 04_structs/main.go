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

	// &variable means give me the memory address of the value this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

//Turn address into value with *address
//Turn value into address with &value

func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson means give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
