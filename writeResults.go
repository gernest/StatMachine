package statmachine

import "fmt"

func ResultInfoString(res []Result) string {
	wins := gamesWon(res)
	draws := gamesDrawn(res)
	lost := gamesLost(res)
	return fmt.Sprintf("Total (W-D-L): %v-%v-%v", wins, draws, lost)
}

func ResultSequenceString(res []Result) string {
	sequence := ""
	for _, r := range res {
		if r.IsWin() {
			sequence += "W"
		} else if r.IsLoss() {
			sequence += "L"
		} else {
			sequence += "D"
		}
	}
	return sequence
}

func GoalInfoString(res []Result) string {
	cleanSheets := cleanSheets(res)
	gamesScoredIn := gamesScoredIn(res)
	return fmt.Sprintf("Games Scored In/Clean Sheets: %v/%v", gamesScoredIn, cleanSheets)
}

func GoalDifferenceString(res []Result) string {
	goals := TotalGoalsScored(res)
	against := TotalGoalsConceded(res)
	return fmt.Sprintf("Goals Scored- Goals Against (difference): %v-%v %v", goals, against, goals-against)
}
