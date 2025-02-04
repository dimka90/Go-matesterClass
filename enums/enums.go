package main

import "fmt"

type StatusCode int

const (
	Ok          StatusCode = 200
	NotFound    StatusCode = 404
	BadRequest  StatusCode = 400
	ServerError StatusCode = 500
)

func (clientMessage StatusCode) serveResponse() {
	switch clientMessage {
	case Ok:
		fmt.Printf("Success")
	case NotFound:
		fmt.Printf("Page not found")
	case ServerError:
		fmt.Printf("An internal error occured")
	case BadRequest:
		fmt.Printf("Bad request")
	default:
		panic("Invalid code")
	}
}
func main() {
	var message StatusCode
	fmt.Printf("Enter status code>>>")
	fmt.Scanln(&message)
	message.serveResponse()

}
