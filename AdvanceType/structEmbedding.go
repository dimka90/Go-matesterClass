package main

import "fmt"

type Telegram struct {
	version float64
	color   string
	year    string
}

type TeleKilo struct {
	Telegram
	capacity int
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

	fmt.Printf("%+v\n", myapp)
}
