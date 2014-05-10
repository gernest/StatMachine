package statmachine

import (
	"testing"
	"time"
)

func TestCalculatingNumberOfPoints(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 0, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, false, 0, time.Now()),
	}

	totalPoints := Points(allResults)

	if 10 != totalPoints {
		t.Errorf("Didnt get correct total number of points, got %v, expected 10", totalPoints)
	}
}