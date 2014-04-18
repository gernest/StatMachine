package statmachine

func firstHalfsScoredInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return scoredInFirstHalf(r)})
}

func secondHalfsScoredInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return scoredInSecondHalf(r)})
}

func firstHalfsConcededInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return concededInFirstHalf(r)})
}

func secondHalfsConcededInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return concededInSecondHalf(r)})
}