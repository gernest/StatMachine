package statmachine

func gamesWon(res []Result) int {
	return countTotal(res, func(r Result) bool { return r.IsWin() })
}

func gamesLost(res []Result) int {
	return countTotal(res, func(r Result) bool { return r.IsLoss() })
}

func gamesDrawn(res []Result) int {
	return countTotal(res, func(r Result) bool { return r.IsDraw() })
}

func gamesScoredIn(res []Result) int {
	return countTotal(res, func(r Result) bool {return r.ScoredAGoal() })
}

func cleanSheets(res []Result) int {
	return countTotal(res, func(r Result) bool { return !concededAGoal(r) })
}

func totalGoalsScored(res []Result) int {
	return sumTotal(res, func(r Result) int { return r.Goals })
}

func totalGoalsConceded(res []Result) int {
	return sumTotal(res, func(r Result) int { return r.OpponentGoals })
}

func sumTotal(res []Result, f func(Result) int) int {
	sum := 0
	for _, r := range res {
		sum += f(r)
	}
	return sum
}
func countTotal(res []Result, f func(Result) bool) int {
	count := 0
	for _, r := range res {
		if f(r) {
			count++
		}
	}
	return count
}
