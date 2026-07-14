package main

import (
	"reflect"
	"testing"
)

// makeArt builds a simple 8-line art slice with an identifying label.
func makeArt(label string) []string {
	art := make([]string, 8)
	for i := range art {
		art[i] = label
	}
	return art
}

func TestMergeBanners_BaseOnlyEntries(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['A'], makeArt("base-A")) {
		t.Errorf("entry only in base must appear in result")
	}
}

func TestMergeBanners_PriorityOnlyEntries(t *testing.T) {
	base := map[rune][]string{}
	priority := map[rune][]string{'B': makeArt("priority-B")}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['B'], makeArt("priority-B")) {
		t.Errorf("entry only in priority must appear in result")
	}
}

func TestMergeBanners_PriorityWinsOnConflict(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['A'], makeArt("priority-A")) {
		t.Errorf("priority entry must overwrite base entry for same rune")
	}
}

func TestMergeBanners_BothContributeDistinctKeys(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'B': makeArt("priority-B")}
	result := MergeBanners(base, priority)
	if _, ok := result['A']; !ok {
		t.Error("'A' from base is missing in result")
	}
	if _, ok := result['B']; !ok {
		t.Error("'B' from priority is missing in result")
	}
}

func TestMergeBanners_DoesNotModifyBase(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	MergeBanners(base, priority)
	if !reflect.DeepEqual(base['A'], makeArt("base-A")) {
		t.Error("MergeBanners must not modify the base map")
	}
}

func TestMergeBanners_DoesNotModifyPriority(t *testing.T) {
	base := map[rune][]string{}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	MergeBanners(base, priority)
	if !reflect.DeepEqual(priority['A'], makeArt("priority-A")) {
		t.Error("MergeBanners must not modify the priority map")
	}
}

func TestMergeBanners_ResultIsNewMap(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{}
	result := MergeBanners(base, priority)
	// Mutating result must not affect base
	result['A'] = makeArt("mutated")
	if reflect.DeepEqual(base['A'], makeArt("mutated")) {
		t.Error("mutating the result must not affect the base map")
	}
}

func TestMergeBanners_BothEmpty(t *testing.T) {
	result := MergeBanners(map[rune][]string{}, map[rune][]string{})
	if result == nil {
		t.Error("result must not be nil even when both inputs are empty")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d entries", len(result))
	}
}

func TestMergeBanners_ResultLength(t *testing.T) {
	base := map[rune][]string{
		'A': makeArt("A"),
		'B': makeArt("B"),
	}
	priority := map[rune][]string{
		'B': makeArt("B-override"),
		'C': makeArt("C"),
	}
	result := MergeBanners(base, priority)
	// Unique keys: A, B, C = 3
	if len(result) != 3 {
		t.Errorf("expected 3 entries in merged result, got %d", len(result))
	}
}

func TestMergeBanners_NilBaseActsAsEmpty(t *testing.T) {
	priority := map[rune][]string{'A': makeArt("A")}
	// Should not panic on nil base
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MergeBanners panicked on nil base: %v", r)
		}
	}()
	result := MergeBanners(nil, priority)
	if result == nil {
		t.Error("result must not be nil")
	}
}
