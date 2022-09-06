package hotelmap

import "fmt"

func BuildEdgeList(hr *[]Edge) {

	for i := range edgelist {
		Addroute(edgelist[i].nodesrc, edgelist[i].nodedest, hr)
	}

	//Addroute("Lobby", "Entrance", hr)
	//Addroute("Lobby", "Reception", hr)
	//Addroute("Lobby", "Bar", hr)

}

func Buildnode(hn *Floorplan) {
	//nodes := hn.Nodes
	nodes := hn.Nodes
	fmt.Println(nodes["bar"])

	//i :=connects[0].nodesrc

}
