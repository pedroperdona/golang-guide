package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	pedro := person{
		firstName: "Pedro",
		lastName:  "Perdoná",
		contact: contact{
			email:   "p.perdona@gmail.com",
			zipCode: 88054600,
		},
	}

	pedro.updateName("João")
	fmt.Printf("%+v", pedro)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
