package move

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lexbarker/hoteladventure/data/player"
	"github.com/lexbarker/hoteladventure/maps/hotelmap"
)

func Move(playerloc *player.Stats, fp []hotelmap.Edge) {
	fmt.Printf("you see..\n")
	fmt.Printf("Current locale %s\n", playerloc.CurrentLocation)
	fmt.Println("You can move to-")
	moves := hotelmap.DirectionOptions(string(playerloc.CurrentLocation), fp)
	for _, place := range moves {
		fmt.Printf("\t%s\n", place)
	}
	var moveto string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Where to now? ...\n")
		var res bool
		nextmove, _ := reader.ReadString('\n')
		res, moveto = legalmove(moves, nextmove)
		if res {
			break
		}
		fmt.Println(" choice not recognised")
	}
	fmt.Println("Moving to  ", string(moveto))
	nextloc := strings.Trim(moveto, "\n")
	playerloc.CurrentLocation = nextloc
}

func legalmove(moveopts []string, playermove string) (bool, string) {
	for _, node := range moveopts {

		move := strings.TrimSuffix(playermove, "\n")

		if strings.EqualFold(node, move) {
			return true, node
		}
	}
	return false, playermove

}
