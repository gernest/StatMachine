package statmachine

import (
		"testing"
		"time"
		)

func TestFirstHalfsScoredInInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,1,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,true,0, time.Now()),
		NewResult(2,0,1,3,1,0,true,0, time.Now()),
		NewResult(3,0,1,1,1,0,true,0, time.Now()),
	}
	
	count := firstHalfsScoredInInARow(allResults)
	
	if 2!=count{
		t.Errorf("Didnt get correct number of first halfs scored in in a row, got %v, expected 2", count)
	}
}


func TestSecondHalfsScoredInInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,1,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,true,0, time.Now()),
		NewResult(2,0,1,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,1,0,true,0, time.Now()),
	}
	
	count := secondHalfsScoredInInARow(allResults)
	
	if 3!=count{
		t.Errorf("Didnt get correct number of second halfs scored in in a row, got %v, expected 3", count)
	}
}

func TestFirstHalfsConcededInInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,1,0,true,0, time.Now()),
		NewResult(1,0,2,1,0,1,true,0, time.Now()),
		NewResult(2,0,1,3,1,2,true,0, time.Now()),
		NewResult(3,0,1,1,1,0,true,0, time.Now()),
	}
	
	count := firstHalfsConcededInInARow(allResults)
	
	if 2!=count{
		t.Errorf("Didnt get correct number of first halfs conceded in in a row, got %v, expected 2", count)
	}
}


func TestSecondHalfsConcededInInARow(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,1,1,true,0, time.Now()),
		NewResult(1,0,2,1,0,0,true,0, time.Now()),
		NewResult(2,0,1,3,1,2,true,0, time.Now()),
		NewResult(3,0,1,1,1,0,true,0, time.Now()),
		NewResult(2,0,1,3,1,2,true,0, time.Now()),
		NewResult(3,0,1,4,1,4,true,0, time.Now()),
	}
	
	count := secondHalfsConcededInInARow(allResults)
	
	if 4!=count{
		t.Errorf("Didnt get correct number of second halfs conceded in in a row, got %v, expected 4", count)
	}
}