package statmachine

import "errors"

type League struct{
	name string
	Teams []*Team
}

func NewLeague(name string) League{
	return League{name, make([]*Team, 0,0)}
}

func FindTeamByName(league League, name string) (*Team, error){

	for _,t := range league.Teams{
		if(t.Name == name){
			return t, nil
		}
	}
	return nil, errors.New("Couldnt find a team with the give name")
}

func FindTeamById(league League, id int) (*Team, error){

	for _,t := range league.Teams{
		if(t.Id() == id){
			return t, nil
		}
	}
	return nil, errors.New("Couldnt find a team with the give id")
}

func TotalNumberOfResults(league League) int{
	numResults:=0
	for _,t := range league.Teams{
		numResults+=len(t.Results)
	}
	return numResults
}

func AllResults(league League) []Result{
	res :=make([]Result, 0)
	for _,t := range league.Teams{
		res = append(res, t.Results...)
	}
	return res
}

func AllGoals(league League) []GoalInfo{
	goals := []GoalInfo{}
	for _,r := range AllResults(league) {
		if r.IsHomeGame {
			for _,g := range r.GoalsInfo {
				goals = append(goals, g)
			}
		}
	}
	return goals
}
