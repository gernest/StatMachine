package statmachine

func firstHalfsScoredInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool { return r.ScoredInFirstHalf() })
}

func secondHalfsScoredInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool { return r.ScoredInSecondHalf() })
}

func firstHalfsConcededInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool { return r.ConcededInFirstHalf() })
}

func secondHalfsConcededInInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool { return concededInSecondHalf(r) })
}

func halfsScoredInInARow(res []Result) int {
	return findLongestSequencOfHalfs(res, func(r Result) bool { return r.ScoredInFirstHalf() }, func(r Result) bool { return r.ScoredInSecondHalf() })
}

func halfsConcededInInARow(res []Result) int {
	return findLongestSequencOfHalfs(res, func(r Result) bool {return r.ConcededInFirstHalf() }, concededInSecondHalf)
}

func findLongestSequencOfHalfs(res []Result, firstHalf func(Result) bool, secondHalf func(Result) bool) int {
	seq := 0
	longestSeq := 0
	for _, r := range res {
		if firstHalf(r) {
			seq++
		} else {
			if seq > longestSeq {
				longestSeq = seq
			}
			seq = 0
		}
		if secondHalf(r) {
			seq++
		} else {
			if seq > longestSeq {
				longestSeq = seq
			}
			seq = 0
		}
	}
	return longestSeq
}
