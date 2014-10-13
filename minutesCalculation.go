package statmachine

func MinutesLeading(res []Result) (int, int, int) {
	leading := 0
	tied := 0
	trailing := 0
	minute := 0
	scored := 0
	conceded := 0
	for _,r := range res {
		minute = 0
		scored = 0
		conceded = 0
		for _,g := range r.GoalsInfo {
			if(scored == conceded) {
				tied = tied + (int(g.Minute) - minute)
			} else if (scored < conceded){
				trailing = trailing + (int(g.Minute) - minute)
			}else {
				leading = leading + (int(g.Minute) - minute)
			}
			minute = int(g.Minute)
			if( g.TeamId == r.Id ){
				scored = scored + 1
			}else {
				conceded = conceded + 1
			}

		}
		if(scored == conceded){
			tied = tied + (90-minute)
		} else if (scored < conceded){
			trailing = trailing + (90 - minute)
		}else {
			leading = leading + (90 - minute)
		}
	}

	return leading, tied, trailing
}
