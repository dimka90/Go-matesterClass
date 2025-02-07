package main

import (
	"github.com/dimka90/interface/types"
)

func main() {
	person1 := types.Person{
		Name: "Dimka",
		Age:  12,
	}
	types.PrintString(person1)
}
