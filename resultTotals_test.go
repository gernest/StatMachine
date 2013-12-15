package statmachine

import "testing"

func TestGamesWon(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,true},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
	}
	
	count := gamesWon(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games won, got %v, expected 3", count)
	}
}

func TestGamesLost(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,true},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
	}
	
	count := gamesLost(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of games lost, got %v, expected 1", count)
	}
}

func TestGamesDrawn(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,2,0,0,true},
		Result{2,3,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
		Result{6,0,1,0,0,true},
	}
	
	count := gamesDrawn(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of games drawn, got %v, expected 4", count)
	}
}

func TestGamesScoredIn(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,0,2,0,0,true},
		Result{2,3,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
		Result{6,0,1,0,0,true},
	}
	
	count := gamesScoredIn(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of games scored in, got %v, expected 5", count)
	}
}

func TestCleanSheets(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,0,2,0,0,true},
		Result{2,3,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
		Result{6,0,1,0,0,true},
	}
	
	count := cleanSheets(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of clean sheets, got %v, expected 1", count)
	}
}