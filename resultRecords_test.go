package statmachine

import (
		"testing"
		"time"
		)

func TestBiggestWin(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,4,0,0,0,false,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,false,0, time.Now()),
		NewResult(4,0,5,2,0,0,false,0, time.Now()),
		NewResult(5,0,1,1,0,0,true,0, time.Now()),
	}
	
	wins := BiggestWins(allResults)
	if(1!=len(wins)){
		t.Errorf("Didnt get correct # of biggest wins, expected 1 but got %v", wins)
	}
	if(!(4==goalsScored(wins[0]) && 0==goalsConceded(wins[0]))){
		t.Errorf("Didnt get correct biggest win, got %v", wins[0])
	}
}

func TestBiggestWinIfNoWin(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,0,true,0, time.Now()),
		NewResult(1,0,0,0,0,0,false,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,false,0, time.Now()),
		NewResult(4,0,0,2,0,0,false,0, time.Now()),
		NewResult(5,0,1,1,0,0,true,0, time.Now()),
	}
	
	wins := BiggestWins(allResults)
	if(0!=len(wins)){
		t.Errorf("Didnt get correct biggest wins, expected to not find any but found %v", len(wins))
	}
}

func TestBiggestWinIfMultiple(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,false,0, time.Now()),
		NewResult(2,0,3,0,0,0,true,0, time.Now()),
		NewResult(3,0,4,1,0,0,false,0, time.Now()),
		NewResult(4,0,0,0,0,0,false,0, time.Now()),
		NewResult(5,0,3,0,1,0,true,0, time.Now()),
	}
	
	wins := BiggestWins(allResults)
	if(3!=len(wins)){
		t.Errorf("Didnt get correct # of biggest wins, expected 3 but got %v", wins)
	}
}