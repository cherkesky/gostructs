package main

import "fmt"

type person struct {
	first   string
	last    string
	contact contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {

	var guy person
	guy.first = "Guy"
	guy.last = "Cherkesky"
	guy.contact.email = "jsdfkjdsfksd"
	guy.contact.zip = 987987

	caty := person{
		first: "Caty",
		last:  "Cherkesky",
		contact: contactInfo{
			email: "SDfsdfsdf",
			zip:   3224424,
		},
	}
	fmt.Println(guy)
	// fmt.Printf("%+v", guy)

	fmt.Println(caty)
	fmt.Printf("%+v", caty)
}
