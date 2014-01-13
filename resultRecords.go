package statmachine

func BiggestWins(res []Result) []Result{
	wins := make([] Result, 0)
	maxWinBy :=0
	for _,r :=range res{
		if(isWin(r)){
			scored :=goalsScored(r)
			conceded :=goalsConceded(r)
			if((scored-conceded)>maxWinBy){
				maxWinBy = goalsScored(r)-goalsConceded(r)
				wins = make([] Result, 0)
				wins = append(wins, r)
			}else if((scored-conceded)==maxWinBy){
				wins = append(wins, r)
			}
		}
	}
	return wins
}

func BiggestLosses(res []Result) []Result{
	losses :=make([]Result, 0)
	maxLossBy :=0
	for _,r :=range res{
		if(isLoss(r)){
			scored :=goalsScored(r)
			conceded :=goalsConceded(r)
			if((conceded-scored)>maxLossBy){
				maxLossBy = conceded-scored
				losses = make([]Result,0)
				losses = append(losses, r)
			}else if((conceded-scored)==maxLossBy){
				losses = append(losses, r)
			}
		}
	}
	return losses
}