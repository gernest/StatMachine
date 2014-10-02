package statmachine

import (
	"testing"
	"time"
)

func TestNoResultsRetrunsNoMatchingResults(t *testing.T) {
	thisSeasonResults := []Result{}
	compareToSeason := []Result{}
	matchingResults := FindMatchingResults(thisSeasonResults, compareToSeason)
	if 0!=len(matchingResults) {
		t.Errorf("Didnt get correct number of matching results, got %v, expected 0", len(matchingResults))
	}
}

func TestNoResultsReturnedWhenNoResultsInTargetSeason(t *testing.T) {
	thisSeasonResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}
	compareToSeason := []Result{}
	matchingResults := FindMatchingResults(thisSeasonResults, compareToSeason)
	if 0!=len(matchingResults) {
		t.Errorf("Didnt get correct number of matching results, got %v, expected 0", len(matchingResults))
	}
}

func TestFindsMultipleMatchingResults(t *testing.T) {
	thisSeasonResults := []Result{
		Result{0, 10, 1, 0, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 11, 2, 0, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 12, 0, 3, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 13, 1, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 14, 3, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 15, 1, 1, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	compareToSeasonResults := []Result{
		Result{6, 10, 1, 1, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{7, 11, 2, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{8, 12, 0, 0, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{9, 13, 1, 1, 0, 0, false, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{10, 14, 3, 1, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{11, 20, 1, 1, 0, 0, true, 2015, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	matchingResults := FindMatchingResults(thisSeasonResults, compareToSeasonResults)
	if 3!=len(matchingResults) {
		t.Errorf("Didnt get correct number of matching results, got %v, expected 3", len(matchingResults))
	}

	for _,rp := range matchingResults {
		if rp.A.Id == 0 {
			if rp.B.Id != 6 {
				t.Errorf("Didnt get correct matching result, got result with Id %v, expected 6", rp.B.Id)			
			}
		}

		if rp.A.Id == 1 {
			if rp.B.Id != 7 {
				t.Errorf("Didnt get correct matching result, got result with Id %v, expected 7", rp.B.Id)			
			}
		}

		if rp.A.Id == 3 {
			if rp.B.Id != 9 {
				t.Errorf("Didnt get correct matching result, got result with Id %v, expected 9", rp.B.Id)			
			}
		}
	}
}