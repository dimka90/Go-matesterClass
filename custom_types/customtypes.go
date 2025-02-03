package main

import "fmt"

type Player struct {
	name         string
	position     string
	jerseyNumber int
}

func main() {
	Amad := Player{
		name:         "Amad",
		position:     "Winger",
		jerseyNumber: 7,
	}

	players := map[int]Player{
		Amad.jerseyNumber: {
			name:         "Amad",
			position:     "Winger",
			jerseyNumber: 7,
		},
	}
	getPlayerDetails(players[Amad.jerseyNumber])

}

func getPlayerDetails(player Player) {

	fmt.Printf("%+v\n", player)
}
