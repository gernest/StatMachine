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
		t.Errorf("Didnt get correct number of games won in a row, got %v, expected 2", count)
	}
}

func TestGamesNotWonInARow(t *testing.T){
	allResults  := []Result{
		Result{0,0,1},
		Result{1,1,1},
		Result{2,0,3},
		Result{3,0,1},
		Result{4,0,0},
		Result{5,1,0},
	}
	
	count := gamesNotWonInARow(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of games not won in a row, got %v, expected 5", count)
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
		t.Errorf("Didnt get correct number of games lost in a row, got %v, expected 3", count)
	}
}

func TestGamesNotLostInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,0},
		Result{1,1,1},
		Result{2,3,0},
		Result{3,4,1},
		Result{4,0,0},
		Result{5,1,0},
		Result{6,1,0},
		Result{5,0,1},
	}
	
	count := gamesNotLostInARow(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of games not lost in a row, got %v, expected 7", count)
	}
}