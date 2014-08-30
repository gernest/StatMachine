package statmachine

import (
	"testing"
	"time"
)

func TestCalculatingNumberOfPoints(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 1, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{1, 0, 2, 0, 0, 0, false, 0, time.Now(), 1, CardInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{3, 0, 1, 1, 0, 0, false, 0, time.Now(),1, CardInfo{}},
		Result{4, 0, 3, 1, 0, 0, false, 0, time.Now(),1, CardInfo{}},
	}

	totalPoints := Points(allResults)

	if 10 != totalPoints {
		t.Errorf("Didnt get correct total number of points, got %v, expected 10", totalPoints)
	}
}
