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
	
	win := biggestWin(allResults)
	if(!(4==goalsScored(win) && 0==goalsConceded(win))){
		t.Errorf("Didnt get correct biggest win, got %v-%v", goalsScored(win), goalsConceded(win))
	}
}