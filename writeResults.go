package statmachine
import "fmt"

func writeResults(res []Result) string{
	wins := gamesWon(res)
	draws := gamesDrawn(res)
	lost := gamesLost(res)
	return fmt.Sprintf("Total (W-D-L): %v-%v-%v",wins, draws,lost)
}