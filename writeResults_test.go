package statmachine

import (
		"testing"
		"time"
		)

func TestWriteResults(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,false,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,false,0, time.Now()),
		NewResult(4,0,3,1,0,0,false,0, time.Now()),
		NewResult(5,0,1,1,0,0,true,0, time.Now()),
	}
	
	desc := ResultInfoString(allResults)
	if(desc!="Total (W-D-L): 3-2-1"){
		t.Errorf("Didnt get correct result set string, got %v", desc)
	}
}

func TestWriteSequenceString(t *testing.T){
	allResults  := []Result{
		NewResult(0,0,1,0,0,0,true,0, time.Now()),
		NewResult(1,0,2,0,0,0,false,0, time.Now()),
		NewResult(2,0,0,3,0,0,true,0, time.Now()),
		NewResult(3,0,1,1,0,0,false,0, time.Now()),
		NewResult(4,0,3,1,0,0,false,0, time.Now()),
		NewResult(5,0,1,1,0,0,true,0, time.Now()),
	}
	
	desc := ResultSequenceString(allResults)
	if(desc!="WWLDWD"){
		t.Errorf("Didnt get correct result sequence string, got %v", desc)
	}

}