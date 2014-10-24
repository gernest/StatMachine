package statmachine

import (
	"testing"
	"time"
	"fmt"
)

func TestAtHome(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	homeResults := HomeResults(allResults)

	if 3 != len(homeResults) {
		t.Errorf("Didnt get correct number of home games, got %v, expected 3", len(homeResults))
	}
}

func TestAtHomeWhenNoHomeGames(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	homeResults := HomeResults(allResults)

	if 0 != len(homeResults) {
		t.Errorf("Didnt get correct number of home games, got %v, expected 0", len(homeResults))
	}
}

func TestAtAway(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	awayResults := AwayResults(allResults)

	if 4 != len(awayResults) {
		t.Errorf("Didnt get correct number of away games, got %v, expected 4", len(awayResults))
	}
}

func TestAtAwayWhenNoAwayGames(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	awayResults := AwayResults(allResults)

	if 0 != len(awayResults) {
		t.Errorf("Didnt get correct number of away games, got %v, expected 0", len(awayResults))
	}
}

func TestLeadingAtHalfTime(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 2, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	leadingAtHalfTimeResults := leadingAtHalfTime(allResults)

	if 2 != len(leadingAtHalfTimeResults) {
		t.Errorf("Didnt get correct number of results leading at half time, got %v, expected 2", len(leadingAtHalfTimeResults))
	}
}

func TestTrailingAtHalfTime(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 1, 2, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	results := trailingAtHalfTime(allResults)

	if 3 != len(results) {
		t.Errorf("Didnt get correct number of results trailing at half time, got %v, expected 3", len(results))
	}
}

func TestResultsBySeason(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 1, 2, false, 1, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 1, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 1, false, 1, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 1, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	results := ResultsBySeason(allResults, 1)
	if 4 != len(results) {
		t.Errorf("Didnt get correct numer of results when filtering by season, got %v, expected 4", len(results))
	}
}

func TestResultsByOpponent(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 1, 2, 0, 1, 2, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 1, 1, 1, 0, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	results := ResultsByOpponent(allResults, 1)
	if 2 != len(results) {
		t.Errorf("Didnt get correct numer of results when filtering by oppoent, got %v, expected 2", len(results))
	}
}

func TestResultsByRounds(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 1, 2, 0, 1, 2, false, 0, time.Now(), 2, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 3, CardInfo{}, []GoalInfo{}},
		Result{3, 1, 1, 1, 0, 1, false, 0, time.Now(), 4, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 0, time.Now(), 5, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 6, CardInfo{}, []GoalInfo{}},
	}
	results := ResultsByRounds(allResults, []int{1, 2, 3})

	if 3 != len(results) {
		t.Errorf("Didnt get correct number of results when filtering by round, got %v, expected 3", len(results))
	}
}

func TestResultsByRoundsWithOneRound(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 1, 2, 0, 1, 2, false, 0, time.Now(), 2, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 3, CardInfo{}, []GoalInfo{}},
		Result{3, 1, 1, 1, 0, 1, false, 0, time.Now(), 4, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 0, time.Now(), 5, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 6, CardInfo{}, []GoalInfo{}},
	}
	results := ResultsByRounds(allResults, []int{6})

	if 1 != len(results) {
		t.Errorf("Didnt get correct number of results when filtering by round, got %v, expected 1", len(results))
	}
}

func TestResultsBeforeDate(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Date(2014, 1, 1, 0,0,0,0, time.UTC), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 1, 2, 0, 1, 2, false, 0, time.Date(2013, 12, 20, 0,0,0,0, time.UTC), 2, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Date(2013, 11, 21, 0,0,0,0, time.UTC), 3, CardInfo{}, []GoalInfo{}},
		Result{3, 1, 1, 1, 0, 1, false, 0, time.Date(2013, 11, 15, 0,0,0,0, time.UTC), 4, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 1, false, 0, time.Date(2013, 11, 14, 0,0,0,0, time.UTC), 5, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Date(2013, 10, 21, 0,0,0,0, time.UTC), 6, CardInfo{}, []GoalInfo{}},
	}
	d := time.Date(2013, 11, 15, 0,0,0,0, time.UTC)
	results := ResultsBeforeDate(allResults, d)
	fmt.Printf("%v\n", results[0])
	fmt.Printf("%v\n", results[1])
	if 2 != len(results) {
		t.Errorf("Didnt get correct number of results when filtering by date, got %v, expected 2", len(results))
	}
}
