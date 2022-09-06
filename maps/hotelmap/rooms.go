package hotelmap

// var nodelist = []string{"Lobby", "Entrance"}

// var Lobby = `{"room": "Lobby", "description": "A dilapidated 1920's style hotel Lobby, reeking of tobacco, weed, booze and misery"}`
// var Entrance = `{"room": "Main Entrance", "description": "Steps leaading to a neglected old hotel Entrance. Well past its heyday"}`

var rooms map[string]MapNode

func BuildRoomNodes() map[string]MapNode {
	rooms = make(map[string]MapNode)

	rooms["entrance"] = MapNode{
		Name:        "Main Entrance",
		Description: "Steps leaading to a neglected old hotel Entrance. Well past its heyday",
	}

	rooms["lobby"] = MapNode{
		Name:        "Lobby",
		Description: "A dilapidated 1920's style hotel Lobby, reeking of tobacco, weed, booze and misery",
	}

	rooms["reception"] = MapNode{
		Name:        "Reception",
		Description: "A messed up booking area a few empty bottles and a keyboard for the rooms ",
	}
	rooms["bar"] = MapNode{
		Name:        "Micks Bar",
		Description: "A hotel late night bar. Stools and a few tables, a small stage for a band, the detritous of many a bad night",
	}

	return rooms
}
