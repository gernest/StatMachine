package statmachine

import "testing"

func TestGamesWonInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,0},
		Result{1,2,0},
		Result{2,0,3},
		Result{3,1,1},
	}
	
	count := gamesWonInARow(allResults)
	
	if 2!=count{
		t.Errorf("Didnt get correct number of games won in a row, got %v games won in a row, expected 2", count)
	}
}

func TestGamesLostInARow(t *testing.T){
	allResults  := []Result{
		Result{0,0,1},
		Result{1,0,4},
		Result{2,0,3},
		Result{3,1,1},
	}
	
	count := gamesLostInARow(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games lost in a row, got %v games won in a row, expected 3", count)
	}
}