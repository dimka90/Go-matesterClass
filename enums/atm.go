package main

import "fmt"

type Transaction int
type BalanceAccount struct {
	balance float64
}

const (
	Deposit = iota + 1
	Withdraw
	CheckBalance
)

var customer BalanceAccount

func (transactionType Transaction) processTraction() {
	switch transactionType {
	case Deposit:
		var amount float64
		fmt.Println("Enter deposit amount:")
		fmt.Scanln(&amount)
		customer.balance += amount
		fmt.Printf("Deposit successful! New balance: %.2f\n", customer.balance)
	case Withdraw:
		var amount float64
		fmt.Println("Enter withdrawal amount:")
		fmt.Scanln(&amount)
		if amount > customer.balance {
			fmt.Println("Insufficient Balance")
			panic("Insuffient balance")
		}
		customer.balance -= amount
		fmt.Printf("Withdrawal successful! New balance:%2.f\n", customer.balance)

	case CheckBalance:
		fmt.Printf("Current balance: %2.f", customer.balance)
	}
}
func main() {
	var action Transaction
	for {
		fmt.Printf("Welcome to Go-ATM 🚀\nChoose a transaction:\n1. Deposit\n2. Withdraw\n3. Check Balance\n")
		fmt.Scanln(&action)
		if action == 4 {
			fmt.Println("👋 Thank you for using Go-ATM! Have a great day!")
			break
		}
		action.processTraction()

	}

}
