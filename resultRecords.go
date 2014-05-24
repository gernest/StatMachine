package statmachine

func BiggestWins(res []Result) []Result {
	return findMaximizingResults(res, func(r Result) bool { return r.IsWin() }, func(r Result) int { return r.Goals - r.OpponentGoals })
}

func BiggestLosses(res []Result) []Result {
	return findMaximizingResults(res, func(r Result) bool {return r.IsLoss() }, func(r Result) int { return r.OpponentGoals - r.Goals })
}

func HighestScoring(res []Result) []Result {
	return findMaximizingResults(res, nil, func(r Result) int { return r.OpponentGoals + r.Goals })
}

func findMaximizingResults(res []Result, filter func(Result) bool, f func(Result) int) []Result {
	results := make([]Result, 0)
	maxValue := 0
	for _, r := range res {
		if nil == filter || filter(r) {
			currentValue := f(r)
			if currentValue > maxValue {
				maxValue = currentValue
				results = make([]Result, 0)
				results = append(results, r)
			} else if currentValue == maxValue {
				results = append(results, r)
			}
		}
	}
	return results
}
