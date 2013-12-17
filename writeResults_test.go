package statmachine

import "testing"

func TestWriteResults(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,false},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,0,false},
		Result{4,3,1,0,0,false},
		Result{5,1,1,0,0,true},
	}
	
	desc := writeResults(allResults)
	if(desc!="Total (W-D-L): 3-2-1"){
		t.Errorf("Didnt get correct result set string, got %v", desc)
	}
}