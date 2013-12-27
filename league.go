package statmachine

import "errors"

type League struct{
	name string
	Teams []Team
}

func NewLeague(name string) League{
	return League{name, make([]Team, 0,0)}
}

func FindTeamByName(league League, name string) (Team, error){
	
	for _,t := range league.Teams{
		if(t.Name == name){
			return t, nil
		}
	}
	return Team{}, errors.New("Couldnt find a team with the give name")
}