package main

import "fmt"

type Printer interface {
	String() string
}

type Person struct {
	name string
	age  int
}

func (person Person) getAge() int {
	return person.age
}

func PrintString(print Printer) {
	fmt.Println(print.String())
}
func (person Person) String() string {
	return fmt.Sprintf("My name is %s and i am %d years ol\n", person.name, person.age)
}
func main() {
	var person1 Person = Person{
		name: "Dimka",
		age:  12,
	}

	// fmt.Println(person1.getAge())
	PrintString(person1)
}
