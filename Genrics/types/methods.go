package types

import "fmt"

func Printer[T any](variable T) {
	fmt.Printf("%+v\n", variable)
}
func (student *Student) CreateId() {

	studentId++
	student.StudentId = studentId
}
