package maihwe

import (
	"testing"
)

func TestMergeBanners(t *testing.T) {
	base := map[rune][]string{
		'A': {"line1"},
		'B': {"base_b"},
	}
	priority := map[rune][]string{
		'A': {"override"},
		'C': {"new"},
	}

	result := MergeBanners(base, priority)

	// priority should win for 'A'
	if result['A'][0] != "override" {
		t.Errorf("expected 'override', got %s", result['A'][0])
	}

	// 'B' from base should still be there
	if result['B'][0] != "base_b" {
		t.Errorf("expected 'base_b', got %s", result['B'][0])
	}

	// 'C' from priority should be there
	if result['C'][0] != "new" {
		t.Errorf("expected 'new', got %s", result['C'][0])
	}

	// base and priority must not be modified
	if len(base) != 2 {
		t.Errorf("base was modified")
	}
}
