package types

import (
	"fmt"
)

func (person Person) getAge() int {
	return person.Age
}

func PrintString(print Printer) {
	fmt.Println(print.String())
}
func (person Person) String() string {
	return fmt.Sprintf("My name is %s and i am %d years old\n", person.Name, person.Age)
}
