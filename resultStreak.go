package statmachine

func gamesWonInARow(res []Result) int {
	return countSequence(res, func(r Result) bool { return r.IsWin() })
}

func gamesLostInARow(res []Result) int {
	return countSequence(res, func(r Result) bool { return r.IsLoss() })
}

func gamesNotWonInARow(res []Result) int {
	return countSequence(res, func(r Result) bool { return !r.IsWin() })
}

func gamesNotLostInARow(res []Result) int {
	return countSequence(res, func(r Result) bool { return !r.IsLoss() })
}

func gamesScoredInInARow(res []Result) int {
	return countSequence(res, func(r Result) bool { return r.ScoredAGoal() })
}

func gamesConcededInARow(res []Result) int {
	return countSequence(res, concededAGoal)
}

func countSequence(res []Result, f func(Result) bool) int {
	count := 0
	for _, e := range res {
		if f(e) {
			count++
		} else {
			break
		}
	}
	return count
}
