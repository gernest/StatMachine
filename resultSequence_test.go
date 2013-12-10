package statmachine

import "testing"

func TestGamesWonInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,2,0,true},
		Result{2,0,3,true},
		Result{3,1,1,true},
	}
	
	count := gamesWonInARow(allResults)
	
	if 2!=count{
		t.Errorf("Didnt get correct number of games won in a row, got %v, expected 2", count)
	}
}

func TestGamesNotWonInARow(t *testing.T){
	allResults  := []Result{
		Result{0,0,1,true},
		Result{1,1,1,true},
		Result{2,0,3,true},
		Result{3,0,1,true},
		Result{4,0,0,true},
		Result{5,1,0,true},
	}
	
	count := gamesNotWonInARow(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of games not won in a row, got %v, expected 5", count)
	}
}

func TestGamesLostInARow(t *testing.T){
	allResults  := []Result{
		Result{0,0,1,true},
		Result{1,0,4,true},
		Result{2,0,3,true},
		Result{3,1,1,true},
	}
	
	count := gamesLostInARow(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games lost in a row, got %v, expected 3", count)
	}
}

func TestGamesNotLostInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,1,1,true},
		Result{2,3,0,true},
		Result{3,4,1,true},
		Result{4,0,0,true},
		Result{5,1,0,true},
		Result{6,1,0,true},
		Result{5,0,1,true},
	}
	
	count := gamesNotLostInARow(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of games not lost in a row, got %v, expected 7", count)
	}
}

func TestGamesScoredInInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,1,1,true},
		Result{2,3,0,true},
		Result{3,4,1,true},
		Result{4,0,0,true},
		Result{5,1,0,true},
		Result{6,1,0,true},
		Result{5,0,1,true},
	}
	
	count := gamesScoredInInARow(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of games scored in in a row, got %v, expected 4", count)
	}
}


func TestGamesConcededInARow(t *testing.T){
	allResults  := []Result{
		Result{0,1,1,true},
		Result{1,1,1,true},
		Result{2,3,3,true},
		Result{3,4,1,true},
		Result{4,0,4,true},
		Result{5,1,5,true},
		Result{6,1,1,true},
		Result{7,0,0,true},
	}
	
	count := gamesConcededInARow(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of games conceded a goal in a row, got %v, expected 7", count)
	}
}