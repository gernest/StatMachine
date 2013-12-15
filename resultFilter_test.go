package statmachine

import "testing"

func TestAtHome(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,false},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,0,false},
		Result{4,3,1,0,0,false},
		Result{5,1,1,0,0,true},
	}
	
	homeResults := homeResults(allResults)
	
	if 3!=len(homeResults){
		t.Errorf("Didnt get correct number of home games, got %v, expected 3", len(homeResults))
	}
}	

func TestAtHomeWhenNoHomeGames(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,false},
		Result{1,2,0,0,0,false},
		Result{2,0,3,0,0,false},
		Result{3,1,1,0,0,false},
		Result{4,3,1,0,0,false},
		Result{5,1,1,0,0,false},
	}
	
	homeResults := homeResults(allResults)
	
	if 0!=len(homeResults){
		t.Errorf("Didnt get correct number of home games, got %v, expected 0", len(homeResults))
	}
}	

func TestAtAway(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,false},
		Result{2,0,3,0,0,false},
		Result{3,1,1,0,0,false},
		Result{4,3,1,0,0,false},
		Result{5,1,1,0,0,true},
	}
	
	awayResults := awayResults(allResults)
	
	if 4!=len(awayResults){
		t.Errorf("Didnt get correct number of away games, got %v, expected 4", len(awayResults))
	}
}	

func TestAtAwayWhenNoAwayGames(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,0,0,true},
		Result{1,2,0,0,0,true},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,0,true},
		Result{4,3,1,0,0,true},
		Result{5,1,1,0,0,true},
	}
	
	awayResults := awayResults(allResults)
	
	if 0!=len(awayResults){
		t.Errorf("Didnt get correct number of away games, got %v, expected 0", len(awayResults))
	}
}

func TestLeadingAtHalfTime(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,1,0,true},
		Result{1,2,0,0,0,false},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,1,false},
		Result{4,3,1,2,1,false},
		Result{5,1,1,0,0,true},
	}
	
	leadingAtHalfTimeResults := leadingAtHalfTime(allResults)
	
	if 2!=len(leadingAtHalfTimeResults){
		t.Errorf("Didnt get correct number of results leading at half time, got %v, expected 2", len(leadingAtHalfTimeResults))
	}
}	

func TestTrailingAtHalfTime(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,1,0,true},
		Result{1,2,0,1,2,false},
		Result{2,0,3,0,0,true},
		Result{3,1,1,0,1,false},
		Result{4,3,1,0,1,false},
		Result{5,1,1,0,0,true},
	}
	
	results := trailingAtHalfTime(allResults)
	
	if 3!=len(results){
		t.Errorf("Didnt get correct number of results trailing at half time, got %v, expected 3", len(results))
	}
}	