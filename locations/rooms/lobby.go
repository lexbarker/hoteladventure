package location

import (
	"encoding/json"

	"github.com/lexbarker/hoteladventure/maps/hotelmap"
)

type response struct {
	Room        string `json:"room"`
	Description string `json:"description"`
}

var Lobby = `{"room": "Lobby", "description": "A dilapidated 1920's style hotel Lobby, reeking of tobacco, weed, booze and misery"}`
var Entrance = `{"room": "Main Entrance", "description": "Steps leaading to a neglected old hotel Entrance. Well past its heyday"}`

func RoomDiscription(room string) (string, string) {
	var detail string
	switch room {
	case "Lobby":
		detail = Lobby
	case "Entrance":
		detail = Entrance
	default:
		detail = `{"room": "not set", "description": "not set"}`
	}

	res := response{}
	json.Unmarshal([]byte(detail), &res)

	return res.Room, res.Description
}

func DescribeRoom(roomname string, rooms *hotelmap.Floorplan) (string, string) {

	//var a map[string]hotelmap.MapNode
	a := rooms.Nodes

	//fmt.Println(a[roomname].Name)

	return a[roomname].Name, a[roomname].Description

}
