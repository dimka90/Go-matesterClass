package internal
type PrepaidMeter struct{
MeterNumber string
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
