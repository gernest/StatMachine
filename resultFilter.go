package statmachine

func HomeResults(res []Result) []Result {
	return filterResults(res, func(r Result) bool { return r.isHomeGame })
}

func AwayResults(res []Result) []Result {
	return filterResults(res, func(r Result) bool { return !r.isHomeGame })
}

func leadingAtHalfTime(res []Result) []Result {
	return filterResults(res, func(r Result) bool { return r.GoalsAtHalfTime > r.opponentGoalsAtHalfTime })
}

func trailingAtHalfTime(res []Result) []Result {
	return filterResults(res, func(r Result) bool { return r.GoalsAtHalfTime < r.opponentGoalsAtHalfTime })
}

func ResultsBySeason(res []Result, seasonId int) []Result {
	return filterResults(res, func(r Result) bool { return seasonId == r.seasonId })
}

func ResultsByOpponent(res []Result, opponentId int) []Result {
	return filterResults(res, func(r Result) bool { return opponentId == r.opponentId })
}

func ResultsByRounds(res []Result, rounds []int) []Result {
	filtered := make([]Result, 0)

	for i, r := range res {
		if oneOfGiveRounds(r, rounds) {
			filtered = append(filtered, res[i])
		}
	}
	return filtered
}

func oneOfGiveRounds(res Result, rounds []int) bool {
	for _, round := range rounds {
		if round == res.Round() {
			return true
		}
	}
	return false
}

func filterResults(res []Result, f func(Result) bool) []Result {
	filtered := make([]Result, 0)

	for i, r := range res {
		if f(r) {
			filtered = append(filtered, res[i])
		}
	}
	return filtered
}
