package statmachine

type Team struct{
	id int
	Name string
	Results []Result
}

func NewTeam(teamId int, teamName string) Team{
	return Team{teamId, teamName, make([]Result, 0,0)}
}