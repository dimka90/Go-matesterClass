// Create a Traffic Light System using enums (iota) and methods.
package main

import "fmt"

type TrafficLight int

const (
	Red TrafficLight = iota
	Green
	Yellow
)

func (trafficSign TrafficLight) nextSign() TrafficLight {
	switch trafficSign {
	case Red:
		fmt.Println("Current Light: Red - Stop")
		return Red
	case Green:
		fmt.Println("Current Light: Green- Go")
		return Green
	case Yellow:
		fmt.Println("Current Light: Yellow- Get Ready")
		return Yellow
	default:
		panic("Error: Unkown light")

	}
}
func main() {
	passengers := Red

	for i := 0; i < 3; i++ {
		passengers = passengers.nextSign()
		passengers++
	}
}
