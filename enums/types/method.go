package types

import "fmt"

const (
	Deposit = iota + 1
	Withdraw
	CheckBalance
)

func (account *BalanceAccount) ProcessTraction(transactionType Transaction) {
	switch transactionType {
	case Deposit:
		var amount float64
		fmt.Println("Enter deposit amount:")
		fmt.Scanln(&amount)
		account.Balance += amount
		fmt.Printf("Deposit successful! New balance: %.2f\n", account.Balance)
	case Withdraw:
		var amount float64
		fmt.Println("Enter withdrawal amount:")
		fmt.Scanln(&amount)
		if amount > account.Balance {
			fmt.Println("Insufficient Balance")
			panic("Insuffient balance")
		}
		account.Balance -= amount
		fmt.Printf("Withdrawal successful! New balance:%2.f\n", account.Balance)

	case CheckBalance:
		fmt.Printf("Current balance: %2.f", account.Balance)
	}
}
