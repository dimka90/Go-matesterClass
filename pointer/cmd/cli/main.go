package main

import "fmt"

func SwapModdify(num1, num2 *int) int {
	num1Copy := *num1
	*num1 = *num2 + 10
	*num2 = num1Copy * 2

	return *num1 + *num2
}

func main() {
	num1 := 5
	num2 := 3
	fmt.Println(SwapModdify(&num1, &num2))

}
