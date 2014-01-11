package statmachine

import (
		"testing"
		"time"
		)

func TestGamesWonInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,true,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,true,0, time.Now()),
	}
	
	count := gamesWonInARow(allResults)
	
	if 2!=count{
		t.Errorf("Didnt get correct number of games won in a row, got %v, expected 2", count)
	}
}

func TestGamesNotWonInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,1,0,0,true,0, time.Now()),
		NewResult(1,0,1,1,0,0,true,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,0,1,0,0,true,0, time.Now()),
		NewResult(4,0,0,0,0,0,true,0, time.Now()),
		NewResult(5,0,1,0,0,0,true,0, time.Now()),
	}
	
	count := gamesNotWonInARow(allResults)
	
	if 5!=count{
		t.Errorf("Didnt get correct number of games not won in a row, got %v, expected 5", count)
	}
}

func TestGamesLostInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,1,0,0,true,0, time.Now()),
		NewResult(1,0,0,4,0,0,true,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,true,0, time.Now()),
	}
	
	count := gamesLostInARow(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of games lost in a row, got %v, expected 3", count)
	}
}

func TestGamesNotLostInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,1,1,0,0,true,0, time.Now()),
		NewResult(2,0,3,0,0,0,true,0, time.Now()),
		NewResult(3,0,4,1,0,0,true,0, time.Now()),
		NewResult(4,0,0,0,0,0,true,0, time.Now()),
		NewResult(5,0,1,0,0,0,true,0, time.Now()),
		NewResult(6,0,1,0,0,0,true,0, time.Now()),
		NewResult(5,0,0,1,0,0,true,0, time.Now()),
	}
	
	count := gamesNotLostInARow(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of games not lost in a row, got %v, expected 7", count)
	}
}

func TestGamesScoredInInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,1,1,0,0,true,0, time.Now()),
		NewResult(2,0,3,0,0,0,true,0, time.Now()),
		NewResult(3,0,4,1,0,0,true,0, time.Now()),
		NewResult(4,0,0,0,0,0,true,0, time.Now()),
		NewResult(5,0,1,0,0,0,true,0, time.Now()),
		NewResult(6,0,1,0,0,0,true,0, time.Now()),
		NewResult(5,0,0,1,0,0,true,0, time.Now()),
	}
	
	count := gamesScoredInInARow(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of games scored in in a row, got %v, expected 4", count)
	}
}


func TestGamesConcededInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,1,0,0,true,0, time.Now()),
		NewResult(1,0,1,1,0,0,true,0, time.Now()),
		NewResult(2,0,3,3,0,0,true,0, time.Now()),
		NewResult(3,0,4,1,0,0,true,0, time.Now()),
		NewResult(4,0,0,4,0,0,true,0, time.Now()),
		NewResult(5,0,1,5,0,0,true,0, time.Now()),
		NewResult(6,0,1,1,0,0,true,0, time.Now()),
		NewResult(7,0,0,0,0,0,true,0, time.Now()),
	}
	
	count := gamesConcededInARow(allResults)
	
	if 7!=count{
		t.Errorf("Didnt get correct number of games conceded a goal in a row, got %v, expected 7", count)
	}
}