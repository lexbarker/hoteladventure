package hotelmap

import "fmt"

type Floorplan struct {
	Nodes  []string
	Routes []Edge
}

type MapNodes struct {
	name        string
	description string
}

type Edge struct {
	node1 string
	node2 string
}

func NewFloorplan() Floorplan {
	var nodeslice []string
	var edgeslice []Edge

	nodemap := Floorplan{
		Nodes:  nodeslice,
		Routes: edgeslice,
	}
	return nodemap
}

func setnodes() {

}

func setroute(node1, node2 string) Edge {
	edgeroute := Edge{
		node1: node1,
		node2: node2,
	}
	fmt.Print(edgeroute)
	return edgeroute

}

func addroute(newroute Edge, nr *[]Edge) {

	*nr = append(*nr, newroute)
}

func Addroute(node1, node2 string, nr *[]Edge) {
	fmt.Printf("route %s", nr)
	addroute(setroute(node1, node2), nr)
	addroute(setroute(node2, node1), nr)
}
