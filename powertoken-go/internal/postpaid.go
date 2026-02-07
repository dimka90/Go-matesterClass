package internal

import "fmt"

type PostpaidMeter struct {
	MeterNumber  string
	CustomerName string
	Address      string
	DebtAmount   float64
	Tariff       float64
}

func NewPostpaidMeter(meterNumber string, customerName string, address string, deptAmount float64, tariff float64) *PostpaidMeter {
	return &PostpaidMeter{
		MeterNumber:  meterNumber,
		CustomerName: customerName,
		Address:      address,
		DebtAmount:   deptAmount,
		Tariff:       tariff,
	}
}

func (p *PostpaidMeter) GetMeterNumber() string {
	return p.MeterNumber
}

func (p *PostpaidMeter) GetMeterType() MeterType {
	return Postpaid
}

func (p *PostpaidMeter) GetBalance() float64 {
	return p.DebtAmount
}

func (p *PostpaidMeter) Addunits() (float64, error) {
	fmt.Println("You don't have implementation for that")
	return 0, nil
}
func (p *PostpaidMeter) GetCustomerName() string {
	return p.CustomerName
}

func (p *PostpaidMeter) AddDebt(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf(" Amount must be greater than zero")
	}
	p.DebtAmount += amount
	return nil
}

func (p *PostpaidMeter) PayBill(amount float64) (float64, error) {

	if amount <= 0 {
		return 0, fmt.Errorf("Amount must be greater thatn zero")
	}
	p.DebtAmount -= amount

	return p.DebtAmount, nil
}

func (p *PostpaidMeter) CalculateCost(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("amount must be greater than zero")
	}
	return amount / p.Tariff, nil
}
func (p *PostpaidMeter) GetCustomerAdress() string {
	return p.Address
}
