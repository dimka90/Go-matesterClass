package main

import (
	"fmt"
	"github.com/dimka0/enums/types"
)

var customer types.BalanceAccount

func main() {
	var action types.Transaction
	var userAcount types.BalanceAccount = types.BalanceAccount{}
	for {
		fmt.Printf("Welcome to Go-ATM 🚀\nChoose a transaction:\n1. Deposit\n2. Withdraw\n3. Check Balance\n")
		fmt.Scanln(&action)
		if action == 4 {
			fmt.Println("👋 Thank you for using Go-ATM! Have a great day!")
			break
		}
		userAcount.ProcessTraction(action)

	}

}
