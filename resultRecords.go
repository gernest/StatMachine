package statmachine

func biggestWin(res []Result) Result{
	win := Result{}
	for _,r :=range res{
		if((goalsScored(r)-goalsConceded(r))>goalsScored(win)-goalsConceded(win)){
			win = r
		}
	}
	return win
}