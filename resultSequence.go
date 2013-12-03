package statmachine

func gamesWonInARow(res []Result) int {
	return countSequence(res, func(r Result) bool {return r.goals>r.opponentGoals})
}

func gamesLostInARow(res []Result) int {
	return countSequence(res, func(r Result) bool {return r.goals<r.opponentGoals})
}

func countSequence(res []Result, f func(Result) bool) int{
	count := 0
	for _,e := range res {
		if(f(e)){
			count++
		}else{
			break
		}
	}
	return count
}