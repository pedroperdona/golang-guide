package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	pedro := person{
		firstName: "Pedro",
		lastName:  "Perdoná",
		contactInfo: contactInfo{
			email:   "p.perdona@gmail.com",
			zipCode: 88054600,
		},
	}

	fmt.Println(pedro)
}
