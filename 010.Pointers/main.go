package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) updateName(newFirstName string, newLastName string) {
	(*p).firstName = newFirstName
	(*p).lastName = newLastName
	fmt.Println(p)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
func main() {

	person1 := person{firstName: "A", lastName: "K"}

	// personPtr := &person1
	// personPtr.updateName("S", "K")
	fmt.Println(&person1)
	person1.updateName("S", "K")
	person1.print()

	myslice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(myslice)
	fmt.Println(myslice)
}
