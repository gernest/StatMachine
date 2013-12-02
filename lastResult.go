package statmachine

func lastWin(res []Result) Result {
	win := Result{}
	for _,e := range res {
		if e.goals>e.opponentGoals {
			win = e
			break
		}
	}
	return win
}