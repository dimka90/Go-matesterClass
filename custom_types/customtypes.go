package main

import (
	"fmt"
)

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

	var userinput int
	player1 := Player{
		name:         "Amad",
		position:     "Winger",
		jerseyNumber: 7,
		hoby:         "Dribbling players",
		club:         " Red Devils",
	}
	player2 := Player{
		name:         "Rashford",
		position:     "Winger",
		jerseyNumber: 10,
		hoby:         "Laziness in the field",
		club:         " Red Devils",
	}
	player1.addPlayer()
	player2.addPlayer()

	fmt.Printf("Enter Player number to get the player details>>")
	fmt.Scanln(&userinput)
	getPlayerDetails(userinput)
	// }
}
