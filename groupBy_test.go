package statmachine

import (
	"strconv"
	"testing"
	"time"
)

func TestGroupBySeasonId(t *testing.T) {
	allResults := []Result{
		NewResult(0, 0, 1, 0, 0, 0, true, 2013, time.Now()),
		NewResult(1, 0, 2, 0, 0, 0, false, 2013, time.Now()),
		NewResult(2, 0, 0, 3, 0, 0, true, 2014, time.Now()),
		NewResult(3, 0, 1, 1, 0, 0, false, 2014, time.Now()),
		NewResult(4, 0, 3, 1, 0, 0, false, 2012, time.Now()),
		NewResult(5, 0, 1, 1, 0, 0, true, 2011, time.Now()),
		NewResult(6, 0, 1, 1, 0, 0, true, 2014, time.Now()),
	}

	groupedResults := GroupBy(allResults, func(r Result) string { return strconv.Itoa(r.SeasonId) })

	if 4 != len(groupedResults) {
		t.Errorf("Didnt get correct number of grouped results, got %v, expected 4", len(groupedResults))
	}

	for key, value := range groupedResults {
		if "2011" == key {
			if 1 != len(value) {
				t.Errorf("Didnt get correct number of grouped results for year %v, got %v, expected 1", key, len(value))
			}
		}
		if "2012" == key {
			if 1 != len(value) {
				t.Errorf("Didnt get correct number of grouped results for year %v, got %v, expected 1", key, len(value))
			}
		}
		if "2013" == key {
			if 2 != len(value) {
				t.Errorf("Didnt get correct number of grouped results for year %v, got %v, expected 2", key, len(value))
			}
		}
		if "2014" == key {
			if 3 != len(value) {
				t.Errorf("Didnt get correct number of grouped results for year %v, got %v, expected 3", key, len(value))
			}
		}
	}
}
