package types

import "fmt"

func Printer[T any](variable T) {
	fmt.Println(variable)
}
