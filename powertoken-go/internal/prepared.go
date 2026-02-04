package internal
import "fmt"
type PrepaidMeter struct{
MeterNumber  string
CustomerName string
Address string
Balance float64
Tariff  float64
}

func NewPrepaidMeter(meterNumber string, customerName string, address string, tariff float64 ) *PrepaidMeter{
return &PrepaidMeter{
MeterNumber: meterNumber,
CustomerName: customerName,
Address: address,
Balance: 0,
Tariff: tariff,
}
}

func (p *PrepaidMeter) GetMeterNumber() string {
return p.MeterNumber
}

func (p *PrepaidMeter) GetMeterType() MeterType{
return  Prepaid
}

func (p *PrepaidMeter) GetBalance() float64{
return p.Balance
}
func (p *PrepaidMeter) Addunits(units float64) error{
if units <= 0 {
return fmt.Errorf("units must be greater than zero")
}
p.Balance += units
return nil
}

func (p *PrepaidMeter) GetCustomerName()  string{
return p.CustomerName
}
