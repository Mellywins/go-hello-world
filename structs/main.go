package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method inside the Person struct (Value Receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}
func (p *Person) hasBirthday() { // (Reference receiver)
	p.age++
}

func (p *Person) getMarried(spouseLastName string) Person {
	p.lastName = spouseLastName
	return *p
}
func main() {
	// init person
	person1 := Person{firstName: "Samantha",
		lastName: "Jones",
		city:     "NYC",
		gender:   "M",
		age:      25}
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
}
