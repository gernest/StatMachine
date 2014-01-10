package statmachine

import "errors"

type Season struct{
	name string
	Teams []*Team
}

func NewSeason(name string) Season{
	return Season{name, make([]*Team, 0,0)}
}

func FindTeamByName(season Season, name string) (*Team, error){
	
	for _,t := range season.Teams{
		if(t.Name == name){
			return t, nil
		}
	}
	return nil, errors.New("Couldnt find a team with the give name")
}