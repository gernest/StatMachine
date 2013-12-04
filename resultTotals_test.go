package statmachine

import "testing"

func TestGamesWon(t *testing.T){
	allResults  := []Result{
		Result{0,1,0},
		Result{1,2,0},
		Result{2,0,3},
		Result{3,1,1},
		Result{4,3,1},
		Result{5,1,1},
	}
	
	count := gamesWon(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games won, got %v, expected 3", count)
	}
}

func TestGamesLost(t *testing.T){
	allResults  := []Result{
		Result{0,1,0},
		Result{1,2,0},
		Result{2,0,3},
		Result{3,1,1},
		Result{4,3,1},
		Result{5,1,1},
	}
	
	count := gamesLost(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of games lost, got %v, expected 1", count)
	}
}

func TestGamesDrawn(t *testing.T){
	allResults  := []Result{
		Result{0,1,0},
		Result{1,2,2},
		Result{2,3,3},
		Result{3,1,1},
		Result{4,3,1},
		Result{5,1,1},
		Result{6,0,1},
	}
	
	count := gamesDrawn(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of games drawn, got %v, expected 4", count)
	}
}