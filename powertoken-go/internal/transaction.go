package internal

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type TransactionStatus int

const (
	Successfull TransactionStatus = iota
	Failed
	Pending
)

type Transaction struct {
	TransactionNo   string
	MeterNumber     string
	CustomerName    string
	Amount          float64
	UnitsPurchased  float64
	Token           string
	PaymentMethod   string
	TransactionDate time.Time
	Status          TransactionStatus // "Successful", "Failed", "Pending"
	MeterType       MeterType
	Address         string
}

func generateTransactionNo() string {
	// Use timestamp + random number
	return fmt.Sprintf("%d%d", time.Now().Unix(), rand.Intn(10000))
}

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator

	token := ""
	for i := 0; i < 5; i++ { // 5 groups
		group := ""
		for j := 0; j < 4; j++ { // 4 digits per group
			digit := rand.Intn(10)       // Random number between 0-9
			group += strconv.Itoa(digit) // Convert int to string and append
		}
		token += group
		if i < 4 {
			token += "-" // Add dash between groups, but not at the end
		}
	}
	return token
}
func NewTransaction(meter Meter, amount float64, paymentMethod string) (*Transaction, error) {
	amountToUint, err := meter.CalculateUnits(amount)
	if err != nil {
		return &Transaction{}, err
	}
	return &Transaction{
		TransactionNo:   generateTransactionNo(),
		MeterNumber:     meter.GetMeterNumber(),
		CustomerName:    meter.GetCustomerName(),
		Amount:          amount,
		UnitsPurchased:  amountToUint,
		Token:           GenerateToken(),
		PaymentMethod:   paymentMethod,
		TransactionDate: time.Now(),
		Status:          Pending,
		MeterType:       meter.GetMeterType(),
		Address:         meter.GetCustomerAddress(),
	}, nil
}

func (t *Transaction) String() string {
	return fmt.Sprintf(`
========== RECEIPT ==========
Transaction No: %s
Date: %s
Status: %v

Customer: %s
Meter Number: %s
Address: %s
Meter Type: %v

Amount Paid: ₦%.2f
Units Purchased: %.2f kWh
Token: %s

Payment Method: %s
=============================
`,
		t.TransactionNo,
		t.TransactionDate.Format("Jan 2, 2006 15:04:05"),
		t.Status,
		t.CustomerName,
		t.MeterNumber,
		t.Address,
		t.MeterType,
		t.Amount,
		t.UnitsPurchased,
		t.Token,
		t.PaymentMethod,
	)
}

func (ts TransactionStatus) String() string {
	switch ts {
	case Successfull:
		return "Successfull"
	case Failed:
		return "Failed"
	case Pending:
		return " Pending"
	default:
		return "Unknown"
	}
}
