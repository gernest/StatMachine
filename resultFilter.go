package statmachine

func atHome(res []Result) []Result {
	home := make([]Result,0)
	
	for _,r := range res{
		if(r.isHomeGame){
			home = append(home, res[1])
		}
	}
	return home
}

func atAway(res []Result) []Result {
	away := make([]Result,0)
	
	for _,r := range res{
		if(!r.isHomeGame){
			away = append(away, res[1])
		}
	}
	return away
}