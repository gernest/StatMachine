package statmachine

func BiggestWins(res []Result) []Result{
	return findMaximizingResults(res, isWin, func(r Result) int{return goalsScored(r)-goalsConceded(r)})
}

func BiggestLosses(res []Result) []Result{
	return findMaximizingResults(res, isLoss, func(r Result) int {return goalsConceded(r)-goalsScored(r)})
}

func HighestScoring(res []Result) []Result{
	return findMaximizingResults(res, nil, func(r Result) int {return goalsConceded(r)+goalsScored(r)})
}

func findMaximizingResults(res []Result, filter func(Result) bool, f func(Result) int) []Result{
	results :=make([]Result, 0)
	maxValue :=0
	for _,r := range res{
		if(nil==filter || filter(r)){
			currentValue :=f(r)
			if(currentValue>maxValue){
				maxValue = currentValue
				results = make([]Result, 0)
				results = append(results, r)
			}else if(currentValue == maxValue){
				results = append(results, r)
			}
		}
	}
	return results
}