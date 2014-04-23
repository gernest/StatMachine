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

func halfsScoredInInARow(res []Result) int {
	seq := 0
	longestSeq := 0
	for _,r := range res{
		if scoredInFirstHalf(r){
			seq++
		}else{
			if(seq>longestSeq){
				longestSeq = seq
			}
			seq = 0
		}
		if(scoredInSecondHalf(r)){
			seq ++
		}else{
			if(seq>longestSeq){
				longestSeq = seq;
			}
			seq = 0
		}
	}
	return longestSeq
}