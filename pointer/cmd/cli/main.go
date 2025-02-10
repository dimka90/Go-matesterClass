package main

import (
	"fmt"
	"github.com/dimka90/pointer/types"
)

func main() {
	// num1 := 5
	// num2 := 3
	// fmt.Println(SwapModdify(&num1, &num2))

	acount1 := types.Account{
		Balance:  1000,
		Owner:    "Dimka",
		IsLocked: false,
	}

	// acount1.lock()
	if err := acount1.Deposit(100); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(acount1.TransferTo(&types.Account{
		Balance:  1000,
		Owner:    "James",
		IsLocked: false,
	}, 1000))

	fmt.Println(acount1.Balance)

	james := &types.Account{
		Balance:  1000.0,
		Owner:    "James",
		IsLocked: false,
	}

	sarah := &types.Account{
		Balance:  500.0,
		Owner:    "Sarah",
		IsLocked: false,
	}

	// Test 1: Successful deposit
	fmt.Println("\nTest 1: Regular deposit")
	fmt.Printf("James initial balance: %.2f\n", james.Balance)
	error1 := james.Deposit(200.0)
	if error1 != nil {
		fmt.Println(error1)
	}
	fmt.Printf("James new balance: %.2f\n", james.Balance)

	// Test 2: Successful transfer
	fmt.Println("\nTest 2: Regular transfer")
	fmt.Printf("Before transfer:\nJames: %.2f\nSarah: %.2f\n", james.Balance, sarah.Balance)
	error2 := james.TransferTo(sarah, 300.0)
	if error2 != nil {
		fmt.Println(error2)
	}
	fmt.Printf("\nAfter transfer:\nJames: %.2f\nSarah: %.2f\n", james.Balance, sarah.Balance)

	// Test 3: Lock account and try deposit
	fmt.Println("\nTest 3: Locked account deposit")
	james.Lock()
	error3 := james.Deposit(100.0)
	if error3 != nil {
		fmt.Println(error3)
	}

	// Test 4: Try transfer with insufficient funds
	fmt.Println("\nTest 4: Transfer with insufficient funds")
	error4 := sarah.TransferTo(james, 1000.0)
	if error4 != nil {
		fmt.Println(error4)
	}
}
