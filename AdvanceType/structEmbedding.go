package main

import "fmt"

type Color int

const (
	ColorGreen Color = iota + 1
	ColorRed
	ColorYellow
)

func (c Color) String() Color {
	switch c {
	case ColorGreen:
		fmt.Printf("Green\n")
		return ColorGreen
	case ColorYellow:
		fmt.Printf("Yellow\n")
		return ColorYellow
	case ColorRed:
		fmt.Printf("Red\n")
		return ColorRed
	default:
		fmt.Println("Invalid traffic light code")
		panic("Error")
	}
}

func (c Color) Print() {
	color := ColorGreen
	for i := 0; i < 3; i++ {
		color = color.String()
		color++
	}
}

type Telegram struct {
	version float64
	color   string
	year    string
}

type TeleKilo struct {
	Telegram
	capacity int
}

type Player struct {
	position     int
	jerseyNumber int
	role         string
}

func main() {
	myapp := TeleKilo{
		Telegram: Telegram{
			version: 1,
			color:   "blue",
			year:    "2020",
		},
		capacity: 20000,
	}

	myapp.version = 3

	fmt.Printf("%+v\n", myapp.Telegram)

	fmt.Printf("%d\n", ColorGreen)
}
