package statmachine

func gamesWonInARow(res []Result) int {
	return countSequence(res, isWin)
}

func gamesLostInARow(res []Result) int {
	return countSequence(res, isLose)
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