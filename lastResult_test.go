package statmachine

import "testing"

func TestLastGameWon(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,0,true,0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,0,4,1,0,0,true,0),
		}
	
	win := lastGameWon(allResults)
	
	if 1!=win.id{
		t.Errorf("Didnt get correct last game won, got result with id %v", win.id)
	}
}


func TestLastGameLost(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,0,true, 0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,0,4,1,0,0,true,0),
		}
	
	lost := lastGameLost(allResults)
	
	if 2!=lost.id{
		t.Errorf("Didnt get correct last game lost, got result with id %v", lost.id)
	}
}

func TestLastGameDrawn(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,0,0,0,0,true,0),
		NewResult(1,0,2,0,0,0,true,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,0,4,1,0,0,true,0),
		}
	drawn := lastGameDrawn(allResults)
	
	if 0!=drawn.id{
		t.Errorf("Didnt get correct last game drawn, got result with id %v", drawn.id)
	}
}