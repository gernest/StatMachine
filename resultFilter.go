package statmachine

func homeResults(res []Result) []Result {
	return filterResults(res, func(r Result) bool{return r.isHomeGame})
}

func awayResults(res []Result) []Result {
	return filterResults(res, func(r Result) bool{return !r.isHomeGame})
}

func leadingAtHalfTime(res []Result) []Result {
	return filterResults(res, func(r Result) bool{return r.goalsAtHalfTime>r.opponentGoalsAtHalfTime})
}

func trailingAtHalfTime(res []Result) []Result {
	return filterResults(res, func(r Result) bool{return r.goalsAtHalfTime<r.opponentGoalsAtHalfTime})
}

func filterResults(res []Result, f func(Result) bool) []Result{
	filtered := make([]Result,0)
	
	for i,r := range res{
		if(f(r)){
			filtered = append(filtered, res[i])
		}
	}
	return filtered
}