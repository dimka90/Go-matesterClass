package internal

type MeterType int

const (
	Prepaid MeterType = iota
	Postpaid
)

func (mt MeterType) String() string {
	switch mt {
	case Prepaid:
		return "Prepaid"
	case Postpaid:
		return "Postpaid"
	default:
		return "Unknown"
	}
}

type Meter interface {
	GetMeterNumber() string
	GetMeterType() MeterType // "Prepaid" or "Postpaid"
	GetBalance() float64
	AddUnits(units float64) (float64, error)
	GetCustomerName() string
	CalculateUnits(amount float64) (float64, error)
	GetCustomerAddress() string
}
