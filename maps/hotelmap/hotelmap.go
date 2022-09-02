package hotelmap

func Buildmap(hr *[]Edge) {
	Addroute("Lobby", "Entrance", hr)
	Addroute("Lobby", "Reception", hr)
	Addroute("Lobby", "Bar", hr)
}
