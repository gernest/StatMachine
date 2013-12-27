package statmachine
import "fmt"

func ResultInfoString(res []Result) string{
	wins := gamesWon(res)
	draws := gamesDrawn(res)
	lost := gamesLost(res)
	return fmt.Sprintf("Total (W-D-L): %v-%v-%v",wins, draws,lost)
}

func GoalInfoString(res []Result) string{
	cleanSheets := cleanSheets(res)
	gamesScoredIn := gamesScoredIn(res)
	return fmt.Sprintf("Games Scored In/Clean Sheets: %v/%v", gamesScoredIn, cleanSheets)
}

func GoalDifferenceString(res []Result) string{
	goals := totalGoalsScored(res)
	against := totalGoalsConceded(res)
	return fmt.Sprintf("Goals Scored- Goals Against (difference): %v-%v %v", goals, against, goals-against)
}