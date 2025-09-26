package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	age int
	isMale bool
	ContactInfo
}

type ContactInfo struct {
	email string
	zipCode int
}

func main() {
	shashank := Person{
		firstName: "shashank",
		lastName: "singh",
		age: 23,
		isMale: true,
		ContactInfo: ContactInfo{
			email: "abc@gmail.com",
			zipCode: 12345,
		},
	}

	shashank.firstName = "joe"
	shashank.ContactInfo.zipCode = 54321
	shashank.updateName("newJoe")
	fmt.Println(shashank.firstName, shashank.ContactInfo.zipCode, shashank)
}

func (p *Person) updateName(newName string) {
	p.firstName = newName
}