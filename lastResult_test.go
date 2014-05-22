package statmachine

import "testing"
import "time"

func TestLastGameWon(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 0, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, true, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 4, 1, 0, 0, true, 0, time.Now()),
	}

	win := lastGameWon(allResults)

	if 1 != win.Id {
		t.Errorf("Didnt get correct last game won, got result with id %v", win.Id)
	}
}

func TestLastGameLost(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 0, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, true, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 4, 1, 0, 0, true, 0, time.Now()),
	}

	lost := lastGameLost(allResults)

	if 2 != lost.Id {
		t.Errorf("Didnt get correct last game lost, got result with id %v", lost.Id)
	}
}

func TestLastGameDrawn(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 0, 0, 0, 0, true, 0, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, true, 0, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 0, time.Now()),
		NewResult(3, 0, 4, 1, 0, 0, true, 0, time.Now()),
	}
	drawn := lastGameDrawn(allResults)

	if 0 != drawn.Id {
		t.Errorf("Didnt get correct last game drawn, got result with id %v", drawn.Id)
	}
}
