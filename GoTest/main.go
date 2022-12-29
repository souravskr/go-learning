package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email: "john@email.com",
			zip:   6700,
		},
	}

	john.updateLastName("Cena")
	fmt.Println()
	john.printName()
}

func (p person) printName() {
	fmt.Printf("Name is: %s %s", p.firstName, p.lastName)
}

func (p *person) updateLastName(lastName string) {
	(*p).lastName = lastName
}
