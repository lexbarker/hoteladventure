package main

import (
	"fmt"

	"github.com/lexbarker/hoteladventure/maps/hotelmap"
)

func main() {
	fmt.Printf("Floorplans POC\n")

	Hotel := hotelmap.NewFloorplan()

	//route := hotelmap.Setroute("Entrance", "Lobby")

	//hotelmap.Addroute(route, &Hotel.Routes)
	hotelmap.Addroute("Entrance", "Lobby", &Hotel.Routes)

	fmt.Printf("routes in map %d", len(Hotel.Routes))

	//fmt.Printf("You can move to %s", hotelmap.DirectionOptions("Lobby", &Hotel.Routes))

}
