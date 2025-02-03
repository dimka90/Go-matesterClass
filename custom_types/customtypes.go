package main

import "fmt"

type Player struct {
	name         string
	position     string
	jerseyNumber int
	hoby         string
	club         string
}

var players map[int]Player

func main() {
	var userinput int

	players = map[int]Player{
		7: {
			name:         "Amad",
			position:     "Winger",
			jerseyNumber: 7,
			hoby:         "Dribbling players",
			club:         " Red Devils",
		},

		10: {
			name:         "Jackson",
			position:     "Striker",
			jerseyNumber: 9,
			hoby:         "Missing empty post",
			club:         "The blues",
		},
	}

	fmt.Printf("Enter Player number to get the player details>>")
	fmt.Scanln(&userinput)
	getPlayerDetails(userinput)

}

func getPlayerDetails(key int) {
	if players[key].jerseyNumber == 0 {
		fmt.Println("Player does not exist")
		return
	}
	fmt.Printf("%+v\n", players[key])
}
