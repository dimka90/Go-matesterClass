package types

import "fmt"
import "errors"

func (a *Account) Deposit(amount float64) error {
	if a.IsLocked {
		return errors.New("Accout is Lock")
	}
	a.Balance += amount
	fmt.Println("Amount Deposited successfully")
	return nil
}

func (a *Account) Lock() {
	if a.IsLocked {
		return
	}
	a.IsLocked = true
	fmt.Println("Account successfully lock")
}

func (a *Account) TransferTo(recipient *Account, amount float64) error {
	if a.IsLocked || recipient.IsLocked {
		return errors.New("Account is lock")
	}
	if a.Balance < amount {
		return errors.New("insufficeint blance")
	}
	a.Balance -= amount
	recipient.Balance += amount
	fmt.Printf("%.2f successfully transfered to the %s Balance", amount, recipient.Owner)

	return nil
}
func SwapModdify(num1, num2 *int) int {
	num1Copy := *num1
	*num1 = *num2 + 10
	*num2 = num1Copy * 2

	return *num1 + *num2
}
