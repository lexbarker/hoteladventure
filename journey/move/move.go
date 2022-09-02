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
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Where to now? ...\n")
	nextmove, _ := reader.ReadString('\n')
	fmt.Println("Moving to  ", string(nextmove))
	nextloc := strings.Trim(nextmove, "\n")
	playerloc.CurrentLocation = nextloc
}
