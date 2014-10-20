package statmachine

import (
	"testing"
	"time"
)

func TestMinutesSinceLastScoredWhenNoGoal(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	mins := MinutesSinceLastScored(allResults)
	if mins != 90 {
		t.Errorf("Didnt get correct minutes since last scored, got %v, expected 90", mins)
	}
}

func TestMinutesSinceLastScoredWhenOnlyGoalInOnlyMatch(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{13, 1}}},
	}

	mins := MinutesSinceLastScored(allResults)
	if mins != 77 {
		t.Errorf("Didnt get correct minutes since last scored, got %v, expected 77", mins)
	}
}

func TestMinutesSinceLastScoredWhenBothTeamsScoredInOnlyMatch(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{22, 101}, GoalInfo{40, 102}}},
	}

	mins := MinutesSinceLastScored(allResults)
	if mins != 68 {
		t.Errorf("Didnt get correct minutes since last scored, got %v, expected 68", mins)
	}
}

func TestMinutesSinceLastScoredWhenTwoResults(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 0, 2, 0, 0, true, 0, time.Date(2014, 9, 28, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}, GoalInfo{40, 102}}},
		Result{101, 102, 1, 3, 0, 0, true, 0, time.Date(2014, 9, 27, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{62, 101}, GoalInfo{40, 102}, GoalInfo{45, 102}, GoalInfo{80, 102}}},
	}

	mins := MinutesSinceLastScored(allResults)
	if mins != 118 {
		t.Errorf("Didnt get correct minutes since last scored, got %v, expected 118", mins)
	}
}

func TestMinutesSinceLastScoredWhenMultipleResults(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 0, 2, 0, 0, true, 0, time.Date(2014, 9, 28, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}, GoalInfo{40, 102}}},
		Result{101, 102, 1, 3, 0, 0, true, 0, time.Date(2014, 9, 27, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 101}, GoalInfo{2, 102}, GoalInfo{10, 102}, GoalInfo{83, 102}}},
		Result{101, 102, 0, 2, 0, 0, true, 0, time.Date(2014, 9, 26, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{67, 102}, GoalInfo{46, 102}}},
		Result{101, 102, 4, 2, 0, 0, true, 0, time.Date(2014, 9, 25, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{4, 101}, GoalInfo{23, 101}, GoalInfo{37, 101}, GoalInfo{81, 101}, GoalInfo{66, 102}, GoalInfo{76, 102}}},
		Result{101, 102, 0, 1, 0, 0, true, 0, time.Date(2014, 9, 24, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}}},
	}

	mins := MinutesSinceLastScored(allResults)
	if mins != 179 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 179", mins)
	}
}


func TestMinutesSinceLastConcededWhenNoGoal(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	mins := MinutesSinceLastConceded(allResults)
	if mins != 90 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 90", mins)
	}
}

func TestMinutesSinceLastConcededWhenOnlyGoalInOnlyMatch(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{22, 102}}},
	}

	mins := MinutesSinceLastConceded(allResults)
	if mins != 68 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 68", mins)
	}
}

func TestMinutesSinceLastConcededWhenBothTeamsConcededInOnlyMatch(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{GoalInfo{22, 101}, GoalInfo{40, 102}}},
	}

	mins := MinutesSinceLastConceded(allResults)
	if mins != 50 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 50", mins)
	}
}

func TestMinutesSinceLastConcededWhenTwoResults(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 2, 0, 0, 0, true, 0, time.Date(2014, 9, 28, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}, GoalInfo{40, 102}}},
		Result{101, 102, 1, 3, 0, 0, true, 0, time.Date(2014, 9, 27, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{62, 101}, GoalInfo{40, 102}, GoalInfo{45, 102}, GoalInfo{81, 102}}},
	}

	mins := MinutesSinceLastConceded(allResults)
	if mins != 99 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 99", mins)
	}
}

func TestMinutesSinceLastConcededWhenMultipleResults(t *testing.T) {
	allResults := []Result{
		Result{101, 102, 2, 0, 0, 0, true, 0, time.Date(2014, 9, 28, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}, GoalInfo{40, 102}}},
		Result{101, 102, 3, 1, 0, 0, true, 0, time.Date(2014, 9, 27, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 101}, GoalInfo{2, 102}, GoalInfo{10, 102}, GoalInfo{83, 102}}},
		Result{101, 102, 2, 0, 0, 0, true, 0, time.Date(2014, 9, 26, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{67, 102}, GoalInfo{46, 102}}},
		Result{101, 102, 4, 2, 0, 0, true, 0, time.Date(2014, 9, 25, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{4, 101}, GoalInfo{23, 101}, GoalInfo{37, 101}, GoalInfo{81, 101}, GoalInfo{66, 102}, GoalInfo{76, 102}}},
		Result{101, 102, 1, 0, 0, 0, true, 0, time.Date(2014, 9, 24, 00, 00, 0, 0, time.UTC), 1, CardInfo{}, []GoalInfo{GoalInfo{1, 102}}},
	}

	mins := MinutesSinceLastConceded(allResults)
	if mins != 97 {
		t.Errorf("Didnt get correct minutes since last conceded, got %v, expected 97", mins)
	}
}
