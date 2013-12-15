package statmachine

func homeResults(res []Result) []Result {
	home := make([]Result,0)
	
	for i,r := range res{
		if(r.isHomeGame){
			home = append(home, res[i])
		}
	}
	return home
}

func awayResults(res []Result) []Result {
	away := make([]Result,0)
	
	for i,r := range res{
		if(!r.isHomeGame){
			away = append(away, res[i])
		}
	}
	return away
}

func leadingAtHalfTime(res []Result) []Result {
	leading := make([]Result,0)
	
	for i,r := range res{
		if(r.goalsAtHalfTime>r.opponentGoalsAtHalfTime){
			leading = append(leading, res[i])
		}
	}
	return leading
}