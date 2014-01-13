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
		t.Errorf("Didnt get correct biggest wins, expected not to find any but found %v", len(wins))
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
	
	for _,r := range wins{
		if(goalsScored(r)-goalsConceded(r)!=3){
			t.Errorf("Got a result with unexpected goal difference in TestBiggestWinIfMultiple, got %v, expected 3",goalsScored(r)-goalsConceded(r)) 
		}
	}
}

func TestBiggestLosses(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,4,0,0,0,false,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,false,0, time.Now()),
		NewResult(4,0,5,2,0,0,false,0, time.Now()),
		NewResult(5,0,1,1,0,0,true,0, time.Now()),
	}
	
	losses := BiggestLosses(allResults)
	if(1!=len(losses)){
		t.Errorf("Didnt get correct # of biggest losses, expected 1 but got %v", losses)
	}
	if(!(0==goalsScored(losses[0]) && 3==goalsConceded(losses[0]))){
		t.Errorf("Didnt get correct biggest loss, got %v", losses[0])
	}
}

func TestBiggestLossesIfNone(t *testing.T){
	allResults :=[]Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
	}
	
	losses :=BiggestLosses(allResults)
	if(0!=len(losses)){
		t.Errorf("Didnt get correct biggest losses results, expected not to find any but found %v", len(losses))
	}
}
func TestBiggestLossesIfMultiple(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,1,0,0,true,0, time.Now()),
		NewResult(1,0,0,1,0,0,false,0, time.Now()),
		NewResult(2,0,0,2,0,0,true,0, time.Now()),
		NewResult(3,0,1,3,0,0,false,0, time.Now()),
		NewResult(4,0,0,0,0,0,false,0, time.Now()),
		NewResult(5,0,0,0,0,0,true,0, time.Now()),
	}
	
	losses :=BiggestLosses(allResults)
	if(2!=len(losses)){
		t.Errorf("Didnt get correct # of biggest losses, expected 2 but got %v", losses)
	}
	
	for _,r := range losses{
		if(goalsScored(r)-goalsConceded(r)!=-2){
			t.Errorf("Got a result with unexpected goal difference in TestBiggestLossesIfMultiple, got %v, expected -2",goalsScored(r)-goalsConceded(r)) 
		}
	}
}