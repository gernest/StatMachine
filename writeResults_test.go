package statmachine

import "testing"

func TestWriteResults(t *testing.T){
	allResults  := []Result{
		NewResult(0,1,0,0,0,0,true,0),
		NewResult(1,2,0,0,0,0,false,0),
		NewResult(2,0,0,3,0,0,true,0),
		NewResult(3,1,0,1,0,0,false,0),
		NewResult(4,3,0,1,0,0,false,0),
		NewResult(5,1,0,1,0,0,true,0),
	}
	
	desc := ResultInfoString(allResults)
	if(desc!="Total (W-D-L): 3-2-1"){
		t.Errorf("Didnt get correct result set string, got %v", desc)
	}
}