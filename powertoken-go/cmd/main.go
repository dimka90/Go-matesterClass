package main
import "github.com/dimka90/powertoken/internal"
import "fmt"
func main(){
meter := internal.NewPrepaidMeter("12234", "Dimka", "Jos", 9)
fmt.Println(meter.GetMeterNumber())
fmt.Println("Codig is fun")
}
