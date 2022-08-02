package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

type contactinfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstname string) {
	p.firstName = newFirstname
}
func main() {

	// Person without Contact Information
	// person1 := person{"S", "K"}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// fmt.Println(person1.lastName)

	person2 := person{firstName: "A", lastName: "K"}
	fmt.Println(person2)
	fmt.Println(person2.firstName)
	fmt.Println(person2.lastName)

	var sk person
	fmt.Println(sk)
	fmt.Printf("%+v", sk)

	sk.firstName = "S"
	sk.lastName = "K"

	fmt.Println(sk)
	fmt.Printf("%+v", sk)

	person3 := person{
		firstName: "A",
		lastName:  "SK",
		contact: contactinfo{
			email:   "ask@gmail.com",
			zipCode: 411030,
		},
	}

	fmt.Println("----------")
	person3.print()
	person3.updateFirstName("Av")
	person3.print()
}
