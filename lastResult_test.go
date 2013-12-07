package statmachine

import "testing"

func TestLastGameWon(t *testing.T){
	allResults  := []Result{
		Result{0,0,0,true},
		Result{1,2,0,true},
		Result{2,0,3,true},
		Result{3,4,1,true},
		}
	
	win := lastGameWon(allResults)
	
	if 1!=win.id{
		t.Errorf("Didnt get correct last game won, got result with id %v", win.id)
	}
}

func TestLastGameLost(t *testing.T){
	allResults  := []Result{
		Result{0,0,0,true},
		Result{1,2,0,true},
		Result{2,0,3,true},
		Result{3,4,1,true},
		}
	
	lost := lastGameLost(allResults)
	
	if 2!=lost.id{
		t.Errorf("Didnt get correct last game lost, got result with id %v", lost.id)
	}
}

func TestLastGameDrawn(t *testing.T){
	allResults  := []Result{
		Result{0,0,0,true},
		Result{1,2,0,true},
		Result{2,0,3,true},
		Result{3,4,1,true},
		}
	drawn := lastGameDrawn(allResults)
	
	if 0!=drawn.id{
		t.Errorf("Didnt get correct last game drawn, got result with id %v", drawn.id)
	}
}