package statmachine

import "sort"

type LeagueTable struct {
	Positions []LeaguePosition
}

type LeaguePosition struct {
	Place    uint8
	Points   uint8
	Wins     uint8
	Draws    uint8
	Losses   uint8
	TeamName string
}

func GenerateLeagueTable(league League, seasonId int) LeagueTable {

	table := LeagueTable{Positions: []LeaguePosition{}}

	sort.Sort(ByName(league.Teams))
	for i, t := range league.Teams {
		table.Positions = append(table.Positions, LeaguePosition{Place: uint8(i+1), Points: 0, Wins: 0, Draws: 0, Losses: 0, TeamName: t.Name})
	}
	return table
}
