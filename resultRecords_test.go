package statmachine

import (
	"testing"
	"time"
)

func TestBiggestWin(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 4, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 5, 2, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	wins := BiggestWins(allResults)
	if 1 != len(wins) {
		t.Errorf("Didnt get correct # of biggest wins, expected 1 but got %v", wins)
	}
	if !(4 == wins[0].Goals && 0 == wins[0].OpponentGoals) {
		t.Errorf("Didnt get correct biggest win, got %v", wins[0])
	}
}

func TestBiggestWinIfNoWin(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 0, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 0, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 0, 2, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	wins := BiggestWins(allResults)
	if 0 != len(wins) {
		t.Errorf("Didnt get correct biggest wins, expected not to find any but found %v", len(wins))
	}
}

func TestBiggestWinIfMultiple(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 3, 0, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 4, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 0, 0, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 3, 0, 1, 0, true, 0, time.Now()),
	}

	wins := BiggestWins(allResults)
	if 3 != len(wins) {
		t.Errorf("Didnt get correct # of biggest wins, expected 3 but got %v", wins)
	}

	for _, r := range wins {
		if r.Goals-r.OpponentGoals != 3 {
			t.Errorf("Got a result with unexpected goal difference in TestBiggestWinIfMultiple, got %v, expected 3", r.Goals-r.OpponentGoals)
		}
	}
}

func TestBiggestLosses(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 4, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 5, 2, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	losses := BiggestLosses(allResults)
	if 1 != len(losses) {
		t.Errorf("Didnt get correct # of biggest losses, expected 1 but got %v", losses)
	}
	if !(0 == losses[0].Goals && 3 == losses[0].OpponentGoals) {
		t.Errorf("Didnt get correct biggest loss, got %v", losses[0])
	}
}

func TestBiggestLossesIfNone(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
	}

	losses := BiggestLosses(allResults)
	if 0 != len(losses) {
		t.Errorf("Didnt get correct biggest losses results, expected not to find any but found %v", len(losses))
	}
}
func TestBiggestLossesIfMultiple(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 0, 1, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 0, 1, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 2, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 3, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 0, 0, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 0, 0, 0, 0, true, 0, time.Now()),
	}

	losses := BiggestLosses(allResults)
	if 2 != len(losses) {
		t.Errorf("Didnt get correct # of biggest losses, expected 2 but got %v", losses)
	}

	for _, r := range losses {
		if r.Goals-r.OpponentGoals != -2 {
			t.Errorf("Got a result with unexpected goal difference in TestBiggestLossesIfMultiple, got %v, expected -2", r.Goals-r.OpponentGoals)
		}
	}
}

func TestHighestScoring(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 4, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 5, 2, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	highestScoring := HighestScoring(allResults)
	if 1 != len(highestScoring) {
		t.Errorf("Didnt get correct # of biggest scoring results, expected 1 but got %v", highestScoring)
	}
	if !(5 == highestScoring[0].Goals && 2 == highestScoring[0].OpponentGoals) {
		t.Errorf("Didnt get correct biggest scoring, got %v", highestScoring[0])
	}
}

func TestBiggestScoringResultsIfMultiple(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 4, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 2, 2, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 3, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 1, 3, 0, 0, false, 0, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 0, time.Now()),
	}

	highestScoring := HighestScoring(allResults)
	if 4 != len(highestScoring) {
		t.Errorf("Didnt get correct # of biggest scoring results, expected 4 but got %v=> %", len(highestScoring), highestScoring)
	}

	for _, r := range highestScoring {
		if r.Goals+r.OpponentGoals != 4 {
			t.Errorf("Got a result with unexpected goal sum in TestBiggestScoringResults, got %v, expected 4", r.Goals+r.OpponentGoals)
		}
	}
}
