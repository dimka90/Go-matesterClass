package main

import "fmt"

type UserStore interface {
	PUT(username string, age int)
	GET(username string)
	DELETE(username string)
	PATCH(username string, newUsername string)
	// GETALL() []User
}

type User struct {
	username string
	age      int
}

var allUsers []User

type MongoDb struct {
}

var users = make(map[string]User)

func (db MongoDb) PUT(username string, age int) {
	users[username] = User{
		username: username,
		age:      age,
	}
	fmt.Printf("%s  successfully Stored", username)
}

func (db MongoDb) GET(username string) {
	userDetails, ok := users[username]
	if !ok {
		panic("User data not found")
	}

	fmt.Printf("%v", userDetails)
}
func (db MongoDb) DELETE(username string) {

	_, ok := users[username]
	if !ok {
		panic("User does not exisit")
	}
	delete(users, username)
}

func (db MongoDb) PATCH(username string, newUsername string) {

	_, ok := users[username]
	if !ok {
		panic("User does not exisit")
	}
	a := users[username].age
	users[username] = users[newUsername]
	users[newUsername] = User{
		username: newUsername,
		age:      a,
	}
}

// func (db MongoDb) GETALL() {

// }

func main() {
	var database UserStore = MongoDb{}

	// Test functionality
	database.PUT("Alice", 25)
	database.GET("Alice")
	database.PATCH("Alice", "Joy")
	database.GET("Joy")
	database.GET("Alice")
	// database.DELETE("Alice")
}
