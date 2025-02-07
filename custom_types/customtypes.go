package main

import (
	"fmt"
)

type Command int

const (
	ADDPLAYER Command = iota
	GETPLAYERDETAILS
	GETALLPLAYERDETAILS
	DELETEPLAYER
)

func handleUserInput(keyBoardInput Command, player *Player) {

	switch keyBoardInput {
	case ADDPLAYER:
		player.addPlayer()
	case GETPLAYERDETAILS:
		fmt.Println("Enter Player Id")
		playerId, _ := fmt.Scanln(&playerId)
		player.getPlayerDetails(playerId)
	case GETALLPLAYERDETAILS:
		player.getallPlayerDetails()
	case DELETEPLAYER:
		fmt.Println("Enter Player Id to Delete")
		playerId, _ := fmt.Scanln(&playerId)
		player.deletePlayer(playerId)
	default:
		fmt.Printf("Invalid command")
	}
}

type Player struct {
	name         string
	position     string
	jerseyNumber int
	hoby         string
	club         string
	playerId     int
}

var players = make(map[int]Player)
var playerId int

func (player *Player) addPlayer() {
	playerId++
	player.playerId = playerId
	players[player.playerId] = *player
	fmt.Println("Player successfully added")
}

func (player Player) getPlayerDetails(key int) {
	player, exists := players[key]

	if !exists {
		fmt.Println("Player does not exist")
		return
	}
	fmt.Printf("%+v\n", player)
}

func (player Player) getallPlayerDetails() {
	for _, v := range players {
		fmt.Printf("%+v\n", v)
	}

}

func (player Player) deletePlayer(playerId int) {
	_, ok := players[playerId]

	if !ok {
		panic("Player with the Id does not exist")
	}
	delete(players, playerId)
	fmt.Println("Player Deleted successfully")
}
func getPlayerDetails(key int) {
	players, exists := players[key]

	if !exists {
		fmt.Println("Player does not exist")
		return
	}
	fmt.Printf("%+v\n", players)
}

func main() {

	var userinput Command
	player1 := Player{
		name:         "Amad",
		position:     "Winger",
		jerseyNumber: 7,
		hoby:         "Dribbling players",
		club:         " Red Devils",
	}

	for {

		fmt.Printf("Choose command\n 0.create a user\n 1. Get Player Detaisl \n 2. Get all Players \n 3. Delete a Player>>\n ")
		fmt.Scanln(&userinput)
		handleUserInput(userinput, &player1)
	}

}
