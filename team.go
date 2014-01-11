package statmachine

type Team struct{
	id int
	Name string
	Results []Result
}

func (t Team) Id() int{
	return t.id
}
func NewTeam(teamId int, teamName string) *Team{
	return &Team{teamId, teamName, make([]Result, 0,0)}
}
