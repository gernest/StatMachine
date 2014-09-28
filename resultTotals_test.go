package statmachine

import (
	"testing"
	"time"
)

func TestGamesWon(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := gamesWon(allResults)

	if 3 != count {
		t.Errorf("Didnt get correct number of games won, got %v, expected 3", count)
	}
}

func TestGamesLost(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := gamesLost(allResults)

	if 1 != count {
		t.Errorf("Didnt get correct number of games lost, got %v, expected 1", count)
	}
}

func TestGamesDrawn(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 3, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := gamesDrawn(allResults)

	if 4 != count {
		t.Errorf("Didnt get correct number of games drawn, got %v, expected 4", count)
	}
}

func TestGamesScoredIn(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 3, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := gamesScoredIn(allResults)

	if 5 != count {
		t.Errorf("Didnt get correct number of games scored in, got %v, expected 5", count)
	}
}

func TestCleanSheets(t *testing.T) {
	allResults := []Result{
		Result{0, 1, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 3, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 1, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 3, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 1, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := cleanSheets(allResults)

	if 1 != count {
		t.Errorf("Didnt get correct number of clean sheets, got %v, expected 1", count)
	}
}

func TestTotalGoalsScored(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 3, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := totalGoalsScored(allResults)

	if 9 != count {
		t.Errorf("Didnt get correct number of total goals scored , got %v, expected 9", count)
	}
}

func TestTotalGoalsConceded(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 3, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := totalGoalsConceded(allResults)

	if 7 != count {
		t.Errorf("Didnt get correct number of total goals conceded, got %v, expected 7", count)
	}
}
