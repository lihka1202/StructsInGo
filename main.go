package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Bjorn",
		contact: contactInfo{
			email:   "me@gmail.com",
			pinCode: 6900,
		},
	}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
