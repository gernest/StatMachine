package statmachine

import (
	"testing"
	"time"
)

func TestMinutesLeadingWhenNoResults(t *testing.T) {
	allResults := []Result{
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 0 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 0", leading)
	}
	if tied != 0 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 0", tied)
	}
	if trailing != 0 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 0", trailing)
	}
}

func TestMinutesLeadingWhenScorelessDraw(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 0 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 0", leading)
	}
	if tied != 90 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 90", tied)
	}
	if trailing != 0 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 0", trailing)
	}
}

func TestMinutesLeadingWhenMultipleScorelessDraw(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{101, 103, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{101, 104, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 0 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 00", leading)
	}
	if tied != 270 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 270", tied)
	}
	if trailing != 0 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 0", trailing)
	}
}

func TestMinutesLeadingWhenOnlyLead(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:12, TeamId:101}}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 78 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 78", leading)
	}
	if tied != 12 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 12", tied)
	}
	if trailing != 0 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 0", trailing)
	}
}

func TestMinutesLeadingWhenLeadingAndTrailingInOneMatch(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:12, TeamId:101}, GoalInfo{Minute:23, TeamId:102}, GoalInfo{Minute:60, TeamId:102}}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 11 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 11", leading)
	}
	if tied != 49 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 49", tied)
	}
	if trailing != 30 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 30", trailing)
	}
}

func TestMinutesLeadingWhenLeadingAndTrailingInOneMatchAndScoringAgainAfterTakingTheLead(t *testing.T) {
	allResults := []Result{
		Result{101, 104, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:28, TeamId:104}, GoalInfo{Minute:61, TeamId:101}, GoalInfo{Minute:69, TeamId:101},GoalInfo{Minute:82, TeamId:101}}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 21 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 21", leading)
	}
	if tied != 28+8 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 36", tied)
	}
	if trailing != 33 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 28", trailing)
	}
}

func TestMinutesLeadingWhenLeadingAndTrailingInMultipleMatches(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:1, TeamId:102}, GoalInfo{Minute:8, TeamId:101}, GoalInfo{Minute:55, TeamId:102}}},
		Result{101, 103, 3, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:12, TeamId:101}, GoalInfo{Minute:20, TeamId:103}, GoalInfo{Minute:24, TeamId:101}, GoalInfo{Minute:47, TeamId:103}, GoalInfo{Minute:87, TeamId:101}}},
		Result{101, 104, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{Minute:28, TeamId:104}, GoalInfo{Minute:61, TeamId:101}, GoalInfo{Minute:69, TeamId:101},GoalInfo{Minute:82, TeamId:101}}},
	}

	leading, tied, trailing := MinutesLeading(allResults)
	if leading != 55 {
		t.Errorf("Didnt get correct minutes leading, got %v, expected 55", leading)
	}
	if tied != 140 {
		t.Errorf("Didnt get correct minutes tied, got %v, expected 140", tied)
	}
	if trailing != 75 {
		t.Errorf("Didnt get correct minutes trailing, got %v, expected 75", trailing)
	}
}