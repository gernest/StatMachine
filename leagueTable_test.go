package statmachine

import "testing"
import "time"
import "fmt"

func TestCreatingLeagueTableWhenNoResults(t *testing.T) {

	league := NewLeague("Test League")
	league.Teams = append(league.Teams, NewTeam(1, "Liverpool"))
	league.Teams = append(league.Teams, NewTeam(2, "Chelsea"))
	league.Teams = append(league.Teams, NewTeam(2, "Arsenal"))

	leagueTable := GenerateLeagueTable(league, 2014)

	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 0, 0, 0, 0, 0, 0, "Arsenal")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 0, 0, 0, 0, 0, 0, "Chelsea")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 0, 0, 0, "Liverpool")
}

func TestCreatingLeagueTableWhenEachResultIsOnlyRecoredOnce(t *testing.T) {

	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	result := Result{1, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, NewTeam(2, "Chelsea"))
	league.Teams = append(league.Teams, NewTeam(3, "Arsenal"))

	leagueTable := GenerateLeagueTable(league, 2014)

	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 3, 1, 0, 0, 4, 0, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 0, 0, 0, 0, 0, 0, "Arsenal")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 0, 0, 0, "Chelsea")
}

func TestCreatingLeagueTableWhenAllResultsAreDuplicate(t *testing.T) {
	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	teamChelsea := NewTeam(2, "Chelsea")
	teamArsenal := NewTeam(3, "Arsenal")

	//liverpool 4 - arsenal 0
	result := Result{1001, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1001, 1, 0, 4, 0, 2, false, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	//liverpool 1 - chelsea 1
	result = Result{1002, 2, 1, 1, 0, 0, true, 2014, time.Now(), 2, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1002, 1, 1, 1, 0, 0, false, 2014, time.Now(), 2, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, teamArsenal)

	leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 4, 1, 1, 0, 5, 1, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 1, 0, 1, 0, 1, 1, "Chelsea")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 1, 0, 4, "Arsenal")
}

func TestCreatingLeagueTableWhenNotAllResultsAreFromSameSeason(t *testing.T) {
	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	teamChelsea := NewTeam(2, "Chelsea")
	teamArsenal := NewTeam(3, "Arsenal")

	//liverpool 4 - arsenal 0
	result := Result{1001, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1001, 1, 0, 4, 0, 2, false, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	//liverpool 1 - chelsea 1
	result = Result{1002, 2, 1, 1, 0, 0, true, 2014, time.Now(), 2, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1002, 1, 1, 1, 0, 0, false, 2014, time.Now(), 2, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	//arsenal 3 - chelsea 0
	result = Result{1003, 2, 3, 0, 0, 0, true, 2014, time.Now(), 3, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)
	result = Result{1003, 3, 0, 3, 0, 0, false, 2014, time.Now(), 3, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	//liverpool 7 - arsenal 0, season 2013
	result = Result{1004, 3, 7, 0, 2, 0, true, 2013, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1004, 1, 0, 7, 0, 2, false, 2013, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, teamArsenal)

	leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 4, 1, 1, 0, 5, 1, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 3, 1, 0, 1, 3, 4, "Arsenal")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 1, 0, 1, 1, 1, 4, "Chelsea")
}

func TestCreatingLeagueTableWhenNotAllResultsAreFromSameSeasonOnlyReturnsTeamsFromRequestedSeason(t *testing.T) {
	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	teamChelsea := NewTeam(2, "Chelsea")
	teamArsenal := NewTeam(3, "Arsenal")

	//liverpool 4 - arsenal 0
	result := Result{1001, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1001, 1, 0, 4, 0, 2, false, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	//liverpool 1 - chelsea 1
	result = Result{1002, 2, 1, 1, 0, 0, true, 2014, time.Now(), 2, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1002, 1, 1, 1, 0, 0, false, 2014, time.Now(), 2, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	//arsenal 3 - chelsea 0
	result = Result{1003, 2, 3, 0, 0, 0, true, 2014, time.Now(), 3, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)
	result = Result{1003, 3, 0, 3, 0, 0, false, 2014, time.Now(), 3, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	//liverpool 7 - arsenal 0, season 2013
	result = Result{1004, 3, 7, 1, 2, 0, true, 2013, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{1004, 1, 1, 7, 0, 2, false, 2013, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, teamArsenal)

	leagueTable := GenerateLeagueTable(league, 2013)
	//printLeagueTable(leagueTable)
	if 3 != len(leagueTable.Positions) {
		t.Errorf("Didnt get expected number of positions %v, got %v", 3, len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 3, 1, 0, 0, 7, 1, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 0, 0, 0, 0, 0, 0, "Chelsea")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 1, 1, 7, "Arsenal")
}

func verifyLeaguePosition(t *testing.T, pos *LeaguePosition, expectedPlace uint8, expectedPoints uint8, expectedWins uint8, expectedDraws uint8, expectedLosses uint8, expectedGoalsFor uint8, expectedGoalsAgainst uint8, expectedTeamName string) {
	if expectedPlace != pos.Place {
		t.Errorf("Didnt get expected place %v, got %v", expectedPlace, pos.Place)
	}

	if expectedPoints != pos.Points {
		t.Errorf("Didnt get expected points %v, got %v", expectedPoints, pos.Points)
	}

	if expectedWins != pos.Wins {
		t.Errorf("Didnt get expected wins %v, got %v", expectedWins, pos.Wins)
	}

	if expectedDraws != pos.Draws {
		t.Errorf("Didnt get expected draws %v, got %v", expectedDraws, pos.Draws)
	}

	if expectedLosses != pos.Losses {
		t.Errorf("Didnt get expected losses %v, got %v", expectedLosses, pos.Losses)
	}

	if expectedTeamName != pos.TeamName {
		t.Errorf("Didnt get expected team Name, %v, got %v", expectedTeamName, pos.TeamName)
	}

	if expectedGoalsFor != pos.GoalsFor {
		t.Errorf("Didnt get expected goals for, %v, got %v", expectedGoalsFor, pos.GoalsFor)
	}

	if expectedGoalsAgainst != pos.GoalsAgainst {
		t.Errorf("Didnt get expected goals against, %v, got %v", expectedGoalsAgainst, pos.GoalsAgainst)
	}
}

func TestLeagueTableIsSortedByPointsThenGoalDifference(t *testing.T) {

	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	result := Result{1, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)

	teamChelsea := NewTeam(2, "Chelsea")
	result = Result{2, 3, 5, 0, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, NewTeam(3, "Arsenal"))

	leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 3, 1, 0, 0, 5, 0, "Chelsea")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 3, 1, 0, 0, 4, 0, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 0, 0, 0, "Arsenal")
}

func TestLeagueTableIsSortedByGoalsScoredWhenGoalDifferenceIsEqual(t *testing.T) {

	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	result := Result{1, 3, 5, 1, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)

	teamChelsea := NewTeam(2, "Chelsea")
	result = Result{2, 3, 4, 0, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, NewTeam(3, "Arsenal"))

	leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	if 3 != len(leagueTable.Positions) {
		t.Errorf("didnt get 3 league positions as expected, got %v", len(leagueTable.Positions))
	}
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 3, 1, 0, 0, 5, 1, "Liverpool")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 3, 1, 0, 0, 4, 0, "Chelsea")
	verifyLeaguePosition(t, leagueTable.Positions[2], 3, 0, 0, 0, 0, 0, 0, "Arsenal")
}

func TestTwoTeamLeagueWhenTheyAreOnlySeperatedByGoalDifference(t *testing.T) {

	league := NewLeague("Test League")
	teamArsenal := NewTeam(1, "Arsenal")
	teamSouthampton := NewTeam(1, "Southampton")

	result := Result{3, 4, 4, 0, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)
	result = Result{3, 4, 0, 4, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)
	
	result = Result{4, 3, 0, 3, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamSouthampton.Results = append(teamSouthampton.Results, result)
	result = Result{4, 3, 2, 1, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamSouthampton.Results = append(teamSouthampton.Results, result)

	league.Teams = append(league.Teams, teamArsenal)
	league.Teams = append(league.Teams, teamSouthampton)
	leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	verifyLeaguePosition(t, leagueTable.Positions[0], 1, 3, 1, 0, 1, 4, 4, "Arsenal")
	verifyLeaguePosition(t, leagueTable.Positions[1], 2, 3, 1, 0, 1, 2, 4, "Southampton")

}
func TestLeaguePositionByRoundAfterOneGame(t *testing.T) {

	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	result := Result{1, 3, 4, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)

	teamChelsea := NewTeam(2, "Chelsea")
	result = Result{2, 3, 1, 1, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)

	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, NewTeam(3, "Arsenal"))

	leaguePositions := FindLeaguePositionsByRound(league, 2014)

	//printLeagueTable(leagueTable)
	if 3 != len(leaguePositions) {
	 	t.Errorf("didnt get 3 league positions as expected, got %v", len(leaguePositions))
	}
	if(leaguePositions["Liverpool"]["1"] != 1){
		t.Errorf("expected Livepool to be in place 1 after 1 round, got place %v", leaguePositions["Liverpool"]["1"])
	}
	if(leaguePositions["Chelsea"]["1"] != 2){
		t.Errorf("expected Chelsea to be in place 2 after 1 round, got place %v", leaguePositions["Chelsea"]["1"])
	}
	if(leaguePositions["Arsenal"]["1"] != 3){
		t.Errorf("expected Arsenal to be in place 3 after 1 round, got place %v", leaguePositions["Arsenal"]["1"])
	}
}

func TestLeaguePositionByRoundAfterTwoRounds(t *testing.T) {

	league := NewLeague("Test League")
	teamLiverpool := NewTeam(1, "Liverpool")
	teamChelsea := NewTeam(2, "Chelsea")
	teamArsenal := NewTeam(3, "Arsenal")
	teamSouthampton := NewTeam(4, "Southampton")
	
	result := Result{1, 2, 2, 0, 2, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{2, 1, 0,2, 0, 2, false, 2014, time.Now(), 1, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)


	result = Result{3, 4, 4, 1, 1, 0, true, 2014, time.Now(), 1, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)
	result = Result{4, 3, 1, 4, 1, 0, false, 2014, time.Now(), 1, CardInfo{}}
	teamSouthampton.Results = append(teamSouthampton.Results, result)

	result = Result{1, 3, 3, 0, 1, 0, false, 2014, time.Now(), 2, CardInfo{}}
	teamLiverpool.Results = append(teamLiverpool.Results, result)
	result = Result{3, 1, 0,3, 0, 1, true, 2014, time.Now(), 2, CardInfo{}}
	teamArsenal.Results = append(teamArsenal.Results, result)

	result = Result{2, 4, 0, 1, 0, 0, false, 2014, time.Now(), 2, CardInfo{}}
	teamChelsea.Results = append(teamChelsea.Results, result)
	result = Result{4, 2, 1, 0, 0, 0, true, 2014, time.Now(), 2, CardInfo{}}
	teamSouthampton.Results = append(teamSouthampton.Results, result)


	league.Teams = append(league.Teams, teamLiverpool)
	league.Teams = append(league.Teams, teamChelsea)
	league.Teams = append(league.Teams, teamArsenal)
	league.Teams = append(league.Teams, teamSouthampton)
	//leagueTable := GenerateLeagueTable(league, 2014)
	//printLeagueTable(leagueTable)
	leaguePositions := FindLeaguePositionsByRound(league, 2014)

	if 4 != len(leaguePositions) {
	 	t.Errorf("didnt get 4 league positions as expected, got %v", len(leaguePositions))
	}
	if(leaguePositions["Arsenal"]["1"] != 1){
		t.Errorf("expected Arsenal to be in place 1 after 1 round, got place %v", leaguePositions["Arsenal"]["1"])
	}
	if(leaguePositions["Liverpool"]["1"] != 2){
		t.Errorf("expected Livepool to be in place 2 after 1 round, got place %v", leaguePositions["Liverpool"]["1"])
	}
	if(leaguePositions["Chelsea"]["1"] != 3){
		t.Errorf("expected Chelsea to be in place 3 after 1 round, got place %v", leaguePositions["Chelsea"]["1"])
	}
	if(leaguePositions["Southampton"]["1"] != 4){
		t.Errorf("expected Southampton to be in place 4 after 1 round, got place %v", leaguePositions["Southampton"]["1"])
	}

	if(leaguePositions["Liverpool"]["2"] != 1){
		t.Errorf("expected Livepool to be in place 1 after 2 round, got place %v", leaguePositions["Liverpool"]["2"])
	}
	if(leaguePositions["Arsenal"]["2"] != 2){
		t.Errorf("expected Arsenal to be in place 2 after 2 round, got place %v", leaguePositions["Arsenal"]["2"])
	}
	if(leaguePositions["Southampton"]["2"] != 3){
		t.Errorf("expected Southampton to be in place 3 after 2 round, got place %v", leaguePositions["Southampton"]["2"])
	}
	if(leaguePositions["Chelsea"]["2"] != 4){
		t.Errorf("expected Chelsea to be in place 4 after 2 round, got place %v", leaguePositions["Chelsea"]["2"])
	}
}

func printLeagueTable(table LeagueTable) {
	fmt.Println("Team Points - Place, w-d-l Goals")
	for _, p := range table.Positions {
		fmt.Printf("%v %v - %v, %v-%v-%v %v-%v", p.TeamName, p.Points, p.Place, p.Wins, p.Draws, p.Losses, p.GoalsFor, p.GoalsAgainst)
		fmt.Println()
	}
}
