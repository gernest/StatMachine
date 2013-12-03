package statmachine

func lastGameWon(res []Result) Result {
	return findFirst(res, func(r Result) bool { return r.goals>r.opponentGoals})
}

func lastGameLost(res []Result) Result {
	return findFirst(res, func(r Result) bool {return r.goals<r.opponentGoals})
}

func lastGameDrawn(res []Result) Result {
	return findFirst(res, func(r Result) bool { return r.goals == r.opponentGoals})
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

