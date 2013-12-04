package statmachine

func gamesWon(res []Result) int{
	return countTotal(res, isWin)
}

func gamesLost(res []Result) int{
	return countTotal(res, isLose)
}

func gamesDrawn(res []Result) int{
	return countTotal(res, isDraw)
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