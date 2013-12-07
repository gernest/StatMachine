package statmachine

import "testing"

func TestAtHome(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,2,0,false},
		Result{2,0,3,true},
		Result{3,1,1,false},
		Result{4,3,1,false},
		Result{5,1,1,true},
	}
	
	homeResults := atHome(allResults)
	
	if 3!=len(homeResults){
		t.Errorf("Didnt get correct number of home games, got %v, expected 3", len(homeResults))
	}
}	

func TestAtHomeWhenNoHomeGames(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,false},
		Result{1,2,0,false},
		Result{2,0,3,false},
		Result{3,1,1,false},
		Result{4,3,1,false},
		Result{5,1,1,false},
	}
	
	homeResults := atHome(allResults)
	
	if 0!=len(homeResults){
		t.Errorf("Didnt get correct number of home games, got %v, expected 0", len(homeResults))
	}
}	

func TestAtAway(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,2,0,false},
		Result{2,0,3,false},
		Result{3,1,1,false},
		Result{4,3,1,false},
		Result{5,1,1,true},
	}
	
	awayResults := atAway(allResults)
	
	if 4!=len(awayResults){
		t.Errorf("Didnt get correct number of away games, got %v, expected 4", len(awayResults))
	}
}	

func TestAtAwayWhenNoAwayGames(t *testing.T){
	allResults  := []Result{
		Result{0,1,0,true},
		Result{1,2,0,true},
		Result{2,0,3,true},
		Result{3,1,1,true},
		Result{4,3,1,true},
		Result{5,1,1,true},
	}
	
	awayResults := atAway(allResults)
	
	if 0!=len(awayResults){
		t.Errorf("Didnt get correct number of away games, got %v, expected 0", len(awayResults))
	}
}	