package statmachine

import "testing"

func TestGamesWon(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,1,0,0,true,0),
	}
	
	count := gamesWon(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games won, got %v, expected 3", count)
	}
}

func TestGamesLost(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,1,0,0,true,0),
	}
	
	count := gamesLost(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of games lost, got %v, expected 1", count)
	}
}

func TestGamesDrawn(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,2,2,0,0,true,0),
		NewResult(2,0,3,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,1,0,0,true,0),
		NewResult(6,0,0,1,0,0,true,0),
	}
	
	count := gamesDrawn(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of games drawn, got %v, expected 4", count)
	}
}

func TestGamesScoredIn(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,0,2,0,0,true,0),
		NewResult(2,0,3,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,1,0,0,true,0),
		NewResult(6,0,0,1,0,0,true,0),
	}
	
	count := gamesScoredIn(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of games scored in, got %v, expected 5", count)
	}
}

func TestCleanSheets(t *testing.T){
	allResults  := []Result{
		NewResult(0,1,0,0,0,0,true,0),
		NewResult(1,0,0,2,0,0,true,0),
		NewResult(2,3,0,3,0,0,true,0),
		NewResult(3,1,0,1,0,0,true,0),
		NewResult(4,3,0,1,0,0,true,0),
		NewResult(5,1,0,1,0,0,true,0),
		NewResult(6,0,0,1,0,0,true,0),
	}
	
	count := cleanSheets(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of clean sheets, got %v, expected 1", count)
	}
}


func TestTotalGoalsScored(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,0,2,0,0,true,0),
		NewResult(2,0,3,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,1,0,0,true,0),
		NewResult(6,0,0,0,0,0,true,0),
	}
	
	count := totalGoalsScored(allResults)
	
	if 9!=count{
		t.Errorf("Didnt get correct number of total goals scored , got %v, expected 9", count)
	}
}

func TestTotalGoalsConceded(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0),
		NewResult(1,0,0,2,0,0,true,0),
		NewResult(2,0,3,3,0,0,true,0),
		NewResult(3,0,1,1,0,0,true,0),
		NewResult(4,0,3,1,0,0,true,0),
		NewResult(5,0,1,0,0,0,true,0),
		NewResult(6,0,0,0,0,0,true,0),
	}
	
	count := totalGoalsConceded(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of total goals conceded, got %v, expected 7", count)
	}
}