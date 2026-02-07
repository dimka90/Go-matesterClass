package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type BillingSystem struct {
	meters       []Meter
	transactions []Transaction
}

func NewBillingSystem() *BillingSystem {

	return &BillingSystem{
		meters:       []Meter{},
		transactions: []Transaction{},
	}
}

func (b *BillingSystem) RegisterMeter(meter Meter) {
	b.meters = append(b.meters, meter)
}

func (b *BillingSystem) ProcessPayment(meterNumber string, amount float64, paymentMethod string) (*Transaction, error) {
	if amount <= 0 {
		return &Transaction{}, fmt.Errorf("Amount cannot be zero")
	}
	for _, meter := range b.meters {
		if meter.GetMeterNumber() == meterNumber {
			new_transaction, err := NewTransaction(meter, amount, paymentMethod)
			if err != nil {
				return &Transaction{}, err
			}
			return new_transaction, nil
		}
	}
	return &Transaction{}, nil
}
func (b *BillingSystem) GetMeterByNumber(meterNumber string) (Meter, error) {
	for _, meter := range b.meters {
		if meter.GetMeterNumber() == meterNumber {
			return meter, nil
		}
	}
	return nil, fmt.Errorf("Meter with such number does not exist")
}
func (b *BillingSystem) ListTransactions() error {
	if len(b.transactions) <= 0 {
		return fmt.Errorf("Zero transactions")
	}
	for _, transaction := range b.transactions {
		fmt.Println(transaction)
	}
	return nil
}
func (b *BillingSystem) SaveTransactionsToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(b.transactions)
	return nil
}

func GenerateReceipt(transaction *Transaction) {
	fmt.Println(" Dimka electricity company")
	fmt.Println(" Transaction Details")
	fmt.Println("[%s]", transaction.Status)
	fmt.Println("Token              [%s]", transaction.Token)
	fmt.Println("Meter Type         [%s]", transaction.MeterType)
	fmt.Println("Address            [%s]", transaction.Address)
	fmt.Println("Payment Method     [%s]", transaction.PaymentMethod)
	fmt.Println("Customenr Name     [%s]", transaction.CustomerName)
	fmt.Println("Uints Purchased    [%s]", transaction.UnitsPurchased)
	fmt.Println("Uints  Amount      [%s]", transaction.Amount)
	fmt.Println("Transaction Date    [%s]", time.Now().Unix())
}
