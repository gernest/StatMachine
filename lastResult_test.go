package statmachine

import "testing"

func TestLastWin(t *testing.T){
	allResults  := []Result{
		Result{0,0,0},
		Result{1,2,0},
		Result{2,0,3},
		Result{3,4,1},
		}
	
	win := lastWin(allResults)
	
	if 1!=win.id{
		t.Errorf("Didnt get correct win, got result with id%v", win.id)
	}
}