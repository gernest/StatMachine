package statmachine

import (
	"testing"
	"time"
)

func TestFirstHalfsScoredInInARow(t *testing.T) {
	allResults := []Result{
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 1, 3, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := firstHalfsScoredInInARow(allResults)

	if 2 != count {
		t.Errorf("Didnt get correct number of first halfs scored in in a row, got %v, expected 2", count)
	}
}

func TestSecondHalfsScoredInInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 1, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := secondHalfsScoredInInARow(allResults)

	if 3 != count {
		t.Errorf("Didnt get correct number of second halfs scored in in a row, got %v, expected 3", count)
	}
}

func TestFirstHalfsConcededInInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 1, 0, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 1, 3, 1, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := firstHalfsConcededInInARow(allResults)

	if 2 != count {
		t.Errorf("Didnt get correct number of first halfs conceded in in a row, got %v, expected 2", count)
	}
}

func TestSecondHalfsConcededInInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 1, 3, 1, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 1, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 1, 3, 1, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 4, 1, 4, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := secondHalfsConcededInInARow(allResults)

	if 4 != count {
		t.Errorf("Didnt get correct number of second halfs conceded in in a row, got %v, expected 4", count)
	}
}

func TestHalfsScoredInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 2, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 2, 3, 1, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 3, 1, 2, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 1, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 1, 4, 1, 4, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := halfsScoredInInARow(allResults)

	if 6 != count {
		t.Errorf("Didnt get correct number of halfs socred in in a row, got %v, expected 6", count)
	}
}

func TestHalfsScoredInInARowWhenNoResults(t *testing.T) {
	allResults := []Result{}

	count := halfsScoredInInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of halfs socred in in a row, got %v, expected 0", count)
	}
}

func TestHalfsScoredInInARowWhenOnlyOne(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 1, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //scored in first half
		Result{1, 0, 1, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //scores in second half
		Result{2, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //does not score
		Result{3, 0, 3, 0, 3, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //scores in second half
		Result{2, 0, 0, 3, 0, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //does not score
		Result{3, 0, 1, 4, 0, 3, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}}, //scores in scond half
	}

	count := halfsScoredInInARow(allResults)

	if 1 != count {
		t.Errorf("Didnt get correct number of halfs socred in in a row, got %v, expected 1", count)
	}
}

func TestHalfsConcededInInARow(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 3, 0, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 3, 0, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 1, 0, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 2, 0, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 3, 0, 2, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 3, 0, 3, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := halfsConcededInInARow(allResults)

	if 5 != count {
		t.Errorf("Didnt get correct number of halfs conceded in in a row, got %v exepcted 5", count)
	}
}

func TestHalfsConcededInInARowWhenNoResults(t *testing.T) {
	allResults := []Result{}

	count := halfsConcededInInARow(allResults)

	if 0 != count {
		t.Errorf("Didnt get correct number of halfs conceded in in a row, got %v exepcted 0", count)
	}
}

func TestHalfsConcededInInARowWhenOnlyOne(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 1, 0, 1, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{1, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{2, 0, 0, 2, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
		Result{3, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}, []GoalInfo{}},
	}

	count := halfsConcededInInARow(allResults)

	if 1 != count {
		t.Errorf("Didnt get correct number of halfs conceded in in a row, got %v exepcted 1", count)
	}
}
