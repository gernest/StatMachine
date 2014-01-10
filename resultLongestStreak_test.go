package statmachine

import "testing"

func TestMostGamesWonInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,2,0,0,0,true,0),
		NewResult(2,0,3,0,0,true,0),
		NewResult(3,1,0,0,0,true,0),
		NewResult(4,2,0,0,0,true,0),
		NewResult(5,3,1,0,0,true,0),
		NewResult(6,1,1,0,0,true,0),
	}
	
	count := mostGamesWonInARow(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of most games won in a row, got %v, expected 3", count)
	}
}

func TestMostGamesWonInARowWhenThereAreNone(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,0,0,0,0,true,0),
		NewResult(2,0,3,0,0,true,0),
		NewResult(3,0,0,0,0,true,0),
		NewResult(4,0,2,0,0,true,0),
		NewResult(5,1,1,0,0,true,0),
		NewResult(6,1,1,0,0,true,0),
	}
	
	count := mostGamesWonInARow(allResults)
	
	if 0!=count{
		t.Errorf("Didnt get correct number of most games won in a row, got %v, expected 0", count)
	}
}

func TestMostGamesWithoutALossInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,2,0,0,0,true,0),
		NewResult(2,0,3,0,0,true,0),
		NewResult(3,1,0,0,0,true,0),
		NewResult(4,2,0,0,0,true,0),
		NewResult(5,3,1,0,0,true,0),
		NewResult(6,1,1,0,0,true,0),
		NewResult(7,1,1,0,0,true,0),
		NewResult(8,0,1,0,0,true,0),
	}
	
	count := mostGamesWithoutALossInARow(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of most games without a loss in a row, got %v, expected 5", count)
	}
}

func TestMostGamesWithoutALossInARowWhenThereAreNone(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,true,0),
		NewResult(1,0,1,0,0,true,0),
		NewResult(2,0,1,0,0,true,0),
		NewResult(3,0,1,0,0,true,0),
		NewResult(4,2,3,0,0,true,0),
		NewResult(5,1,2,0,0,true,0),
		NewResult(6,1,2,0,0,true,0),
	}
	
	count := mostGamesWithoutALossInARow(allResults)
	
	if 0!=count{
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}

func TestMostGamesLostInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,2,0,0,0,true,0),
		NewResult(2,0,3,0,0,true,0),
		NewResult(3,1,0,0,0,true,0),
		NewResult(4,2,0,0,0,true,0),
		NewResult(5,3,1,0,0,true,0),
		NewResult(6,1,1,0,0,true,0),
	}
	
	count := mostGamesLostInARow(allResults)
	
	if 1!=count{
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 1", count)
	}
}

func TestMostGamesLostInARowWhenThereAreNone(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,0,0,0,0,true,0),
		NewResult(2,0,0,0,0,true,0),
		NewResult(3,0,0,0,0,true,0),
		NewResult(4,2,0,0,0,true,0),
		NewResult(5,1,1,0,0,true,0),
		NewResult(6,1,1,0,0,true,0),
	}
	
	count := mostGamesLostInARow(allResults)
	
	if 0!=count{
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}


func TestMostGamesWithoutAWinInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,true,0),
		NewResult(1,0,0,0,0,true,0),
		NewResult(2,0,0,0,0,true,0),
		NewResult(3,0,0,0,0,true,0),
		NewResult(4,0,0,0,0,true,0),
		NewResult(5,0,0,0,0,true,0),
		NewResult(6,0,0,0,0,true,0),
		NewResult(7,0,0,0,0,true,0),
		NewResult(8,1,0,0,0,true,0),
	}
	
	count := mostGamesWithoutAWinInARow(allResults)
	
	if 8!=count{
		t.Errorf("Didnt get correct number of most games without a win in a row, got %v, expected 8", count)
	}
}

func TestMostGamesWithoutAWinInARowWhenThereAreNone(t *testing.T){
	allResults  := []Result{
		NewResult(0,2,1,0,0,true,0),
		NewResult(1,2,1,0,0,true,0),
		NewResult(2,2,1,0,0,true,0),
		NewResult(3,2,1,0,0,true,0),
		NewResult(4,4,3,0,0,true,0),
		NewResult(5,4,2,0,0,true,0),
		NewResult(6,4,2,0,0,true,0),
	}
	
	count := mostGamesWithoutAWinInARow(allResults)
	
	if 0!=count{
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}