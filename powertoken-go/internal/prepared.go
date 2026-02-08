package internal

import "fmt"

type PrepaidMeter struct {
	MeterNumber  string
	CustomerName string
	Address      string
	Balance      float64
	Tariff       float64
}

func NewPrepaidMeter(meterNumber string, customerName string, address string, tariff float64) *PrepaidMeter {
	return &PrepaidMeter{
		MeterNumber:  meterNumber,
		CustomerName: customerName,
		Address:      address,
		Balance:      0,
		Tariff:       tariff,
	}
}

func (p *PrepaidMeter) GetMeterNumber() string {
	return p.MeterNumber
}

func (p *PrepaidMeter) GetMeterType() MeterType {
	return Prepaid
}

func (p *PrepaidMeter) GetBalance() float64 {
	return p.Balance
}
func (p *PrepaidMeter) AddUnits(units float64) (float64, error) {
	if units <= 0 {
		return 0, fmt.Errorf("units must be greater than zero")
	}
	p.Balance += units
	return p.Balance, nil
}

func (p *PrepaidMeter) GetCustomerName() string {
	return p.CustomerName
}

func (p *PrepaidMeter) CalculateUnits(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("amount must be greater than zero")
	}
	return amount / p.Tariff, nil
}

func (p *PrepaidMeter) GetCustomerAddress() string {
	return p.Address
}
