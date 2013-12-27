package statmachine

func gamesWon(res []Result) int{
	return countTotal(res, isWin)
}

func gamesLost(res []Result) int{
	return countTotal(res, isLoss)
}

func gamesDrawn(res []Result) int{
	return countTotal(res, isDraw)
}

func gamesScoredIn(res []Result) int{
	return countTotal(res, func (r Result) bool{return r.goals>0})
}

func cleanSheets(res []Result) int{
	return countTotal(res, func (r Result) bool {return 0==r.opponentGoals})
}

func countTotal(res []Result, f func(Result) bool) int{
	count :=0
	for _,r := range res{
		if(f(r)){
			count++
		}
	}
	return count
}