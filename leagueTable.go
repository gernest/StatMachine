package statmachine

import "sort"

type LeagueTable struct {
	Positions []*LeaguePosition
}

type LeaguePosition struct {
	Place        uint8
	TeamName     string
	Points       uint8
	Wins         uint8
	Draws        uint8
	Losses       uint8
	GoalsFor     uint8
	GoalsAgainst uint8
}

func (table *LeagueTable) findByTeamName(name string) *LeaguePosition {
	for i, p := range table.Positions {
		if name == p.TeamName {
			return table.Positions[i]
		}
	}
	return nil
}

func GenerateLeagueTable(league League, seasonId int) LeagueTable {
	return GenerateLeagueTableWithFilter(league, seasonId, func(r Result) bool {return true})
}

func GenerateLeagueTableWithFilter(league League, seasonId int, f func(Result) bool) LeagueTable {

	table := LeagueTable{Positions: []*LeaguePosition{}}

	for _, t := range league.Teams {
		for _, r := range t.Results {
			if seasonId == r.SeasonId && f(r){
				pos := table.findByTeamName(t.Name)
				if nil == pos {
					pos = new(LeaguePosition)
					pos.Place = uint8(len(table.Positions) + 1)
					pos.TeamName = t.Name
					table.Positions = append(table.Positions, pos)
				}
				pos.GoalsFor = pos.GoalsFor + uint8(r.Goals)
				pos.GoalsAgainst = pos.GoalsAgainst + uint8(r.OpponentGoals)
				if r.IsWin() {
					pos.Points = pos.Points + 3
					pos.Wins = pos.Wins + 1
				} else if r.IsDraw() {
					pos.Points = pos.Points + 1
					pos.Draws = pos.Draws + 1
				} else {
					pos.Losses = pos.Losses + 1
				}
			}
		}
	}

	sort.Sort(ByPointsThenGoalDifference(table.Positions))

	for i, _ := range table.Positions {
		table.Positions[i].Place = uint8(i + 1)
	}
	return table
}

type ByPointsThenGoalDifference []*LeaguePosition

func (p ByPointsThenGoalDifference) Len() int {
	return len(p)
}

func (p ByPointsThenGoalDifference) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPointsThenGoalDifference) Less(i int, j int) bool {
	return p[i].Points > p[j].Points ||
		(p[i].Points == p[j].Points && (p[i].GoalsFor-p[i].GoalsAgainst) > (p[j].GoalsFor-p[j].GoalsAgainst))
}
