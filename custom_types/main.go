package main

import (
	"fmt"
	"github.com/dimka90/football-players/types"
)

func main() {

	var userinput types.Command
	player1 := types.Player{
		Name:         "Amad",
		Position:     "Winger",
		JerseyNumber: 7,
		Hoby:         "Dribbling players",
		Club:         " Red Devils",
	}

	for {

		fmt.Printf("Choose command\n 0.create a user\n 1. Get Player Detaisl \n 2. Get all Players \n 3. Delete a Player>>\n ")
		fmt.Scanln(&userinput)
		types.HandleUserInput(userinput, &player1)
	}

}
