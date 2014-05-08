package statmachine

import (
	"testing"
	"time"
)

func TestAtHome(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	homeResults := HomeResults(allResults)

	if 3 != len(homeResults) {
		t.Errorf("Didnt get correct number of home games, got %v, expected 3", len(homeResults))
	}
}

func TestAtHomeWhenNoHomeGames(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, false, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, false, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, false, 0, time.Now()),
	}

	homeResults := HomeResults(allResults)

	if 0 != len(homeResults) {
		t.Errorf("Didnt get correct number of home games, got %v, expected 0", len(homeResults))
	}
}

func TestAtAway(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, false, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	awayResults := AwayResults(allResults)

	if 4 != len(awayResults) {
		t.Errorf("Didnt get correct number of away games, got %v, expected 4", len(awayResults))
	}
}

func TestAtAwayWhenNoAwayGames(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, true, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, true, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, true, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	awayResults := AwayResults(allResults)

	if 0 != len(awayResults) {
		t.Errorf("Didnt get correct number of away games, got %v, expected 0", len(awayResults))
	}
}

func TestLeadingAtHalfTime(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 1, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 1, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 2, 1, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	leadingAtHalfTimeResults := leadingAtHalfTime(allResults)

	if 2 != len(leadingAtHalfTimeResults) {
		t.Errorf("Didnt get correct number of results leading at half time, got %v, expected 2", len(leadingAtHalfTimeResults))
	}
}

func TestTrailingAtHalfTime(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 1, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 1, 2, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 1, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 1, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	results := trailingAtHalfTime(allResults)

	if 3 != len(results) {
		t.Errorf("Didnt get correct number of results trailing at half time, got %v, expected 3", len(results))
	}
}

func TestResultsBySeason(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 1, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 1, 2, false, 1, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 1, time.Now()),
		NewResult(3, 0, 1, 1, 0, 1, false, 1, time.Now()),
		NewResult(4, 0, 3, 1, 0, 1, false, 1, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	results := ResultsBySeason(allResults, 1)
	if 4 != len(results) {
		t.Errorf("Didnt get correct numer of results when filtering by season, got %v, expected 4", len(results))
	}
}

func TestResultsByOpponent(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 1, 0, true, 0, time.Now()),
		NewResult(1, 1, 2, 0, 1, 2, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 1, 1, 1, 0, 1, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 1, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	results := ResultsByOpponent(allResults, 1)
	if 2 != len(results) {
		t.Errorf("Didnt get correct numer of results when filtering by oppoent, got %v, expected 2", len(results))
	}
}

func TestResultsByRounds(t *testing.T) {
	allResults := []Result{
		NewResultWithRound(0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1),
		NewResultWithRound(1, 1, 2, 0, 1, 2, false, 0, time.Now(), 2),
		NewResultWithRound(2, 0, 0, 3, 0, 0, true, 0, time.Now(), 3),
		NewResultWithRound(3, 1, 1, 1, 0, 1, false, 0, time.Now(), 4),
		NewResultWithRound(4, 0, 3, 1, 0, 1, false, 0, time.Now(), 5),
		NewResultWithRound(5, 0, 1, 1, 0, 0, true, 0, time.Now(), 6),
	}
	results := ResultsByRounds(allResults, []int{1, 2, 3})

	if 3 != len(results) {
		t.Errorf("Didnt get correct number of results when filtering by round, got %v, expected 3", len(results))
	}
}

func TestResultsByRoundsWithOneRound(t *testing.T) {
	allResults := []Result{
		NewResultWithRound(0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1),
		NewResultWithRound(1, 1, 2, 0, 1, 2, false, 0, time.Now(), 2),
		NewResultWithRound(2, 0, 0, 3, 0, 0, true, 0, time.Now(), 3),
		NewResultWithRound(3, 1, 1, 1, 0, 1, false, 0, time.Now(), 4),
		NewResultWithRound(4, 0, 3, 1, 0, 1, false, 0, time.Now(), 5),
		NewResultWithRound(5, 0, 1, 1, 0, 0, true, 0, time.Now(), 6),
	}
	results := ResultsByRounds(allResults, []int{6})

	if 1 != len(results) {
		t.Errorf("Didnt get correct number of results when filtering by round, got %v, expected 1", len(results))
	}
}
