package internal

type PostpaidMeter struct{
MeterNumber string
CustomerName string
Address string
DebtAmount float64
Tariff float64
}

func NewPostpaidMeter(meterNumber string, customerName string, address string, deptAmount float64, tariff float64) *PostpaidMeter{
return &PostpaidMeter{
MeterNumber: meterNumber,
CustomerName: customerName,
Address: address,
DebtAmount: deptAmount,
Tariff: tariff, 
} 
}
