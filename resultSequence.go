package statmachine

func gamesWonInARow(res []Result) int {
	return countSequence(res, isWin)
}

func gamesLostInARow(res []Result) int {
	return countSequence(res, isLose)
}

func gamesNotWonInARow(res []Result) int {
	return countSequence(res, func(r Result) bool{return !isWin(r)})
}

func gamesNotLostInARow(res []Result) int {
	return countSequence(res, func(r Result) bool{return !isLose(r)})
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