package hotelmap

import "fmt"

func Showmap() {
	fmt.Printf("MAP")
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
