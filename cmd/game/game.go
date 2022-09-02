package main

import (
	"fmt"

	"github.com/lexbarker/hoteladventure/data/player"
	"github.com/lexbarker/hoteladventure/journey/move"
	"github.com/lexbarker/hoteladventure/maps/hotelmap"
)

func main() {
	fmt.Printf("The Hotel\n\n")
	fmt.Println("Explore the Hotel, but watch out for surprises")

	gameloop()
}

func gameloop() {
	//set up play info
	var player1 player.Player
	player1.PlayerInfo = player.Setup()
	//set up hotel map
	Hotel := hotelmap.NewFloorplan()
	hotelmap.Buildmap(&Hotel.Routes)

	//display info
	fmt.Printf("Health:\t %d\n", player1.PlayerInfo.Health)
	fmt.Printf("hello, please enter the Hotel\n")
	fmt.Printf("routes %d\n", len(Hotel.Routes))
	//start game loop
	for {
		move.Move(&player1.PlayerInfo, Hotel.Routes)
		fmt.Printf("\t moving to  %s\n", player1.PlayerInfo.CurrentLocation)
	}
}
