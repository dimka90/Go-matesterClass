package internal
import (
"time"
"math/rand"
"strconv"
"fmt"
)
type TransactionStatus int

const (
Successfull   TransactionStatus = iota
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
    Status           TransactionStatus // "Successful", "Failed", "Pending"
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
            digit := rand.Intn(10) // Random number between 0-9
            group += strconv.Itoa(digit) // Convert int to string and append
        }
        token += group
        if i < 4 {
            token += "-" // Add dash between groups, but not at the end
        }
    }
    return token
}
func NewTransaction(meter Meter, amount float64, paymentMethod string)  (*Transaction, error){

return &Transaction{
TransactionNo:generateTransactionNo(),
MeterNumber: meter.GetMeterNumber(),
CustomerName: meter.GetCustomerName(),
Amount: amount,
UnitsPurchased:meter.CalculateUnits(amount),
Token: GenerateToken(),
PaymentMethod: paymentMethod,
TransactionDate:  time.Now(),
Status: Pending,
}, nil
}

