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