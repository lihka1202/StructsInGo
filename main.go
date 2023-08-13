package main

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Bjorn",
		contactInfo: contactInfo{
			email:   "me@gmail.com",
			pinCode: 6900,
		},
	}
	alex.updateName("Akhil")
	alex.print()
}
