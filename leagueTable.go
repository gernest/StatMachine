package statmachine

type LeagueTable struct {
	LeaguePositions []LeaguePosition
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
	return LeagueTable{LeaguePositions: []LeaguePosition{
		LeaguePosition{Place: 1, Points: 0, Wins: 0, Draws: 0, Losses: 0},
		LeaguePosition{Place: 2, Points: 0, Wins: 0, Draws: 0, Losses: 0}}}
}
