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
	if(""!=team.Name){
		t.Errorf("Should have gotten an empty name, got %v", team.Name)
	}
}