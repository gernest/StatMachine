package statmachine

type Team struct {
	id      int
	Name    string
	Results []Result
}

func (t Team) Id() int {
	return t.id
}

func NewTeam(teamId int, teamName string) *Team {
	return &Team{teamId, teamName, make([]Result, 0, 0)}
}

type ByName []Team

func (t ByName) Len() int {
	return len(t.Name)
}

func (t ByName) Swap(i, j) {
	t[i], t[j] = t[j], t[i]
}

func (t ByName) Less(i, j) {
	return t[i].Name < t[j].Name
}
