package statmachine

func lastGameWon(res []Result) Result {
	return findFirst(res, isWin)
}

func lastGameLost(res []Result) Result {
	return findFirst(res, isLoss)
}

func lastGameDrawn(res []Result) Result {
	return findFirst(res, isDraw)
}

func findFirst(res []Result, f func(Result) bool) Result {
	found := Result{}
	for _,e := range res {
		if f(e){
			found = e
			break;
		}
	}
	return found
}

