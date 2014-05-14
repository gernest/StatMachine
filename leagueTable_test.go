package statmachine

import "testing"

func TestCreatingLeagueTableWhenNoResults(t *testing.T) {

	league := NewLeague("Test League")
	league.Teams = append(league.Teams, NewTeam(1, "Liverpool"))
	league.Teams = append(league.Teams, NewTeam(2, "Arsenal"))

	leagueTable := GenerateLeagueTable(league, 2014)

	if 2 != len(leagueTable.LeaguePositions) {
		t.Errorf("didnt get 2 league positions as expected, got %v", len(leagueTable.LeaguePositions))
	}

	verifyLeaguePosition(t, leagueTable.LeaguePositions[0], 1, 0, 0, 0, 0)
	verifyLeaguePosition(t, leagueTable.LeaguePositions[1], 2, 0, 0, 0, 0)
}

func TestCreatingLeagueTableWhenEachResultIsOnlyRecoredOnce(t *testing.T) {

}

func TestCreatingLeagueTableWhenAllResultsAreDuplicate(t *testing.T) {

}

func TestCreatingLeagueTableWhenAllResultsAreFromSameSeason(t *testing.T) {
	// allResults := []Result{
	// 	NewResult(0, 0, 1, 0, 0, 0, true, 2014, time.Now()),
	// 	NewResult(1, 0, 2, 0, 0, 0, false, 2014, time.Now()),
	// 	NewResult(2, 0, 0, 3, 0, 0, true, 2014, time.Now()),
	// 	NewResult(3, 0, 1, 1, 0, 0, false, 2014, time.Now()),
	// 	NewResult(4, 0, 3, 1, 0, 0, false, 2014, time.Now()),
	// 	NewResult(5, 0, 1, 1, 0, 0, true, 2014, time.Now()),
	// 	NewResult(6, 0, 1, 1, 0, 0, true, 2014, time.Now()),
	// }
}

func verifyLeaguePosition(t *testing.T, pos LeaguePosition, expectedPlace uint8, expectedPoints uint8, expectedWins uint8, expectedDraws uint8, expectedLosses uint8) {
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
}
