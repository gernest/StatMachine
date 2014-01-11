package statmachine

import "testing"

func TestFindExistingTeamInLeague(t *testing.T){
	league := NewLeague("Test League")
	league.Teams = append(league.Teams, NewTeam(1, "Liverpool"))
	league.Teams = append(league.Teams, NewTeam(2, "Arsenal"))
	
	team,err := FindTeamByName(league, "Liverpool")
	if(nil!=err){
		t.Error("Got error from find by team when I should not have")
	}
	if("Liverpool"!=team.Name){
		t.Errorf("Didint find the team I expected, found team with Name %v", team.Name)
	}
}

func TestFindNoneExistingTeamInLeague(t *testing.T){
	league := NewLeague("Test League")
	league.Teams = append(league.Teams, NewTeam(1, "Liverpool"))
	league.Teams = append(league.Teams, NewTeam(2, "Arsenal"))
	
	team,err := FindTeamByName(league, "Hull")
	if(nil==err){
		t.Error("Did not get an error from find by team when I should have")
	}
	if(nil!=team){
		t.Error("Should have gotten a nil reference, got an team instance")
	}
}

func TestNumberOfGames(t *testing.T){
	league :=NewLeague("Test League")
	teamA :=NewTeam(1, "Liverpool")
	teamA.Results = []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
	}
	league.Teams = append(league.Teams, teamA)
	teamB :=NewTeam(1, "Arsenal")
	teamB.Results = []Result{
		NewResult(0,0,1,0,0,0,true,0),
	}
	league.Teams = append(league.Teams, teamB)
	
	numResults :=TotalNumberOfResults(league)
	if(4!=numResults){
		t.Errorf("expected 4 results, got %v", numResults)
	}
}