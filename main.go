package main

import "fmt"

type person struct {
	first string
	last  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {

	var guy person
	guy.first = "Guy"
	guy.last = "Cherkesky"
	guy.contactInfo.email = "jsdfkjdsfksd"
	guy.contactInfo.zip = 987987

	caty := person{
		first: "Caty",
		last:  "Cherkesky",
		contactInfo: contactInfo{
			email: "SDfsdfsdf",
			zip:   3224424,
		},
	}

	guy.print()
	catyPointer := &caty
	catyPointer.updateName("Catherine")
	caty.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(pointerToPerson).first = newFirstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
