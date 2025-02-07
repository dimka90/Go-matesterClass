package types

import (
	"fmt"
)

var players = make(map[int]Player)
var PlayerId int

type Command int

const (
	ADDPLAYER Command = iota
	GETPLAYERDETAILS
	GETALLPLAYERDETAILS
	DELETEPLAYER
)

func (player *Player) addPlayer() {
	PlayerId++
	player.PlayerId = PlayerId
	players[player.PlayerId] = *player
	fmt.Println("Player successfully added")
}

func (player Player) deletePlayer(playerId int) {
	_, ok := players[playerId]

	if !ok {
		panic("Player with the Id does not exist")
	}
	delete(players, playerId)
	fmt.Println("Player Deleted successfully")
}
func GetPlayerDetails(key int) {
	players, exists := players[key]

	if !exists {
		fmt.Println("Player does not exist")
		return
	}
	fmt.Printf("%+v\n", players)
}

func (player Player) GetallPlayerDetails() {
	for _, v := range players {
		fmt.Printf("%+v\n", v)
	}

}

func HandleUserInput(keyBoardInput Command, player *Player) {

	switch keyBoardInput {
	case ADDPLAYER:
		player.addPlayer()
	case GETPLAYERDETAILS:
		fmt.Println("Enter Player Id")
		var playerId int
		playerId, _ = fmt.Scanln(&playerId)
		GetPlayerDetails(playerId)
	case GETALLPLAYERDETAILS:
		player.GetallPlayerDetails()
	case DELETEPLAYER:
		fmt.Println("Enter Player Id to Delete")
		var playerId int
		playerId, _ = fmt.Scanln(&playerId)
		player.deletePlayer(playerId)
	default:
		fmt.Printf("Invalid command")
	}
}
