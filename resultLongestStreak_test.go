package statmachine

import (
	"testing"
	"time"
)

func TestMostGamesWonInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWonInARow(allResults)

	if 3 != count {
		t.Errorf("Didnt get correct number of most games won in a row, got %v, expected 3", count)
	}
}

func TestMostGamesWonInARowWhenThereAreNone(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWonInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of most games won in a row, got %v, expected 0", count)
	}
}

func TestMostGamesWithoutALossInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{7, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{8, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWithoutALossInARow(allResults)

	if 5 != count {
		t.Errorf("Didnt get correct number of most games without a loss in a row, got %v, expected 5", count)
	}
}

func TestMostGamesWithoutALossInARowWhenThereAreNone(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 2, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWithoutALossInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}

func TestMostGamesLostInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 3, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesLostInARow(allResults)

	if 1 != count {
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 1", count)
	}
}

func TestMostGamesLostInARowWhenThereAreNone(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesLostInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}

func TestMostGamesWithoutAWinInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{7, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{8, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWithoutAWinInARow(allResults)

	if 8 != count {
		t.Errorf("Didnt get correct number of most games without a win in a row, got %v, expected 8", count)
	}
}

func TestMostGamesWithoutAWinInARowWhenThereAreNone(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{4, 0, 4, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{5, 0, 4, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{6, 0, 4, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := mostGamesWithoutAWinInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of most games lost in a row, got %v, expected 0", count)
	}
}
