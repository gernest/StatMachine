package statmachine

import "testing"
import "time"

func TestLastGameWon(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{3, 0, 4, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
	}

	win := lastGameWon(allResults)

	if 1 != win.Id {
		t.Errorf("Didnt get correct last game won, got result with id %v", win.Id)
	}
}

func TestLastGameLost(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{3, 0, 4, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
	}

	lost := lastGameLost(allResults)

	if 2 != lost.Id {
		t.Errorf("Didnt get correct last game lost, got result with id %v", lost.Id)
	}
}

func TestLastGameDrawn(t *testing.T) {
	allResults := []Result{
		Result{0, 0, 0, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{1, 0, 2, 0, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{2, 0, 0, 3, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
		Result{3, 0, 4, 1, 0, 0, true, 0, time.Now(), 1, CardInfo{}},
	}
	drawn := lastGameDrawn(allResults)

	if 0 != drawn.Id {
		t.Errorf("Didnt get correct last game drawn, got result with id %v", drawn.Id)
	}
}
