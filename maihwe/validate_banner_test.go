package maihwe

import (
	"testing"
)

func TestValidateBanner(t *testing.T) {
	// nil check
	if err := ValidateBanner(nil); err == nil {
		t.Error("expected error for nil banner")
	}

	// wrong entry count
	small := map[rune][]string{'A': {"l1", "l2", "l3", "l4", "l5", "l6", "l7", "l8"}}
	if err := ValidateBanner(small); err == nil {
		t.Error("expected error for wrong entry count")
	}

	// build a valid banner (95 entries, 8 lines each)
	good := make(map[rune][]string)
	for r := rune(32); r <= 126; r++ {
		good[r] = []string{"l1", "l2", "l3", "l4", "l5", "l6", "l7", "l8"}
	}
	if err := ValidateBanner(good); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
