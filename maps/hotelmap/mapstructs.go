package hotelmap

type Floorplan struct {
	Nodes  map[string]MapNode
	Routes []Edge
}

type MapNode struct {
	Name        string
	Description string
	//edges       []string
}

type Edge struct {
	node1 string
	node2 string
}
