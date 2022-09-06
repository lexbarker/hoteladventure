package hotelmap

import "fmt"

func Showmap() {
	fmt.Printf("MAP")
}

func NewFloorplan() Floorplan {
	var nodeslice map[string]MapNode
	var edgeslice []Edge

	nodemap := Floorplan{
		Nodes:  nodeslice,
		Routes: edgeslice,
	}
	return nodemap
}

func setnode(name string) {

}

func setroute(node1, node2 string) Edge {
	edgeroute := Edge{
		node1: node1,
		node2: node2,
	}
	//fmt.Print(edgeroute)
	return edgeroute

}

func addroute(newroute Edge, nr *[]Edge) {

	*nr = append(*nr, newroute)
}

func Addroute(node1, node2 string, nr *[]Edge) {
	//fmt.Printf("route %s", nr)
	addroute(setroute(node1, node2), nr)
	addroute(setroute(node2, node1), nr)
}
func AddNode(name string, content string, graphnode *[]MapNode) {

}

func DirectionOptions(location string, r []Edge) []string {
	var nextplace []string

	for _, ro := range r {
		if string(ro.node1) == location {
			nextplace = append(nextplace, ro.node2)
		}
	}
	return nextplace
}
