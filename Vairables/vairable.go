package main

import "fmt"

func main() {

	var (
		num1, num2 float64
		operator   string
	)

	fmt.Print("Enter a the first number>>")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number>>")
	fmt.Scanln(&num2)

	fmt.Print("Enter the operator>>")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Println(num1 + num2)

	case "-":
		fmt.Println(num2 - num1)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 != 0 {
			fmt.Println(num1 / num1)
		} else {
			fmt.Println("Error: Divsion by zero")
		}
	default:
		fmt.Println("Invalid operator")

	}

}
