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
	caty.print()
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
