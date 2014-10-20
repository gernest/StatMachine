package statmachine

import "sort"

func MinutesSinceLastScored(res []Result) int {
	mins := 0
	sort.Sort(ByDateDesc(res))
	for _, r := range res {
		if r.Goals > 0 {
			lastGoalInMatch := 0
			for _, g := range r.GoalsInfo {
				if r.OpponentId != g.TeamId && int(g.Minute) > lastGoalInMatch {
					lastGoalInMatch = int(g.Minute)
				}
			}
			mins = 90 - lastGoalInMatch
		} else {
			mins = mins + 90
		}
	}
	return mins
}

func MinutesSinceLastConceded(res []Result) int {
	mins := 0
	sort.Sort(ByDateDesc(res))
	for _, r := range res {
		if r.OpponentGoals > 0 {
			lastGoalInMatch := 0
			for _, g := range r.GoalsInfo {
				if r.OpponentId == g.TeamId && int(g.Minute) > lastGoalInMatch {
					lastGoalInMatch = int(g.Minute)
				}
			}
			mins = 90 - lastGoalInMatch
		} else {
			mins = mins + 90
		}
	}
	return mins
}