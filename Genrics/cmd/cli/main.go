package main

import "github.com/dimka90/generic/types"
import "github.com/dimka90/generic/mockUser"

func main() {
	mockuser.Student.CreateId()
	mockuser.Student2.CreateId()
	types.Printer(mockuser.Student)
	types.Printer(mockuser.Student2)
}
