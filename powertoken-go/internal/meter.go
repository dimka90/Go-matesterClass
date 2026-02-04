package internal
type MeterType int
const (
Prepaid  MeterType = iota
Postpaid
)
type Meter interface {
    GetMeterNumber() string
    GetMeterType()  MeterType // "Prepaid" or "Postpaid"
    GetBalance() float64
    AddUnits(units float64) error
    GetCustomerName() string
    CalculateUnits(amount float64) float64
}
