package statmachine

func mostGamesWonInARow(res []Result) int {
	return longestSequence(res, isWin)
}

func mostGamesWithoutALossInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return isWin(r) || isDraw(r)})
}

func mostGamesLostInARow(res []Result) int{
	return longestSequence(res, isLoss)
}

func mostGamesWithoutAWinInARow(res []Result) int {
	return longestSequence(res, func(r Result) bool{return isLoss(r) || isDraw(r)})
}

func longestSequence(res []Result, f func(Result) bool) int{
	seq :=0
	longestSeq :=0
	for _,e := range res{
		if(f(e)){
			seq++
		}else{
			if(seq>longestSeq){
				longestSeq = seq
				seq = 0
			}
		}
	}
	if(seq>longestSeq){
		longestSeq = seq
	}
			
	return longestSeq
}