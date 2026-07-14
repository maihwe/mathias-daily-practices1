package main

import (
	"strings"
	"testing"
)

func TestPadArtRows_PadsShortRows(t *testing.T) {
	input := []string{"hi", "there", "a", "b", "c", "d", "e", "f"}
	result := PadArtRows(input, 8)
	for i, row := range result {
		if len(row) != 8 {
			t.Errorf("row %d: expected length 8, got %d (%q)", i, len(row), row)
		}
	}
}

func TestPadArtRows_PaddingIsSpaces(t *testing.T) {
	input := []string{"ab", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)
	if result[0] != "ab   " {
		t.Errorf("expected \"ab   \", got %q", result[0])
	}
	// All padding characters must be spaces
	if strings.TrimLeft(result[0][2:], " ") != "" {
		t.Errorf("padding must be spaces only, got %q", result[0][2:])
	}
}

func TestPadArtRows_DoesNotTruncate(t *testing.T) {
	// Row wider than target width must come back unchanged
	input := []string{"hello world", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)
	if result[0] != "hello world" {
		t.Errorf("wider row must not be truncated: got %q", result[0])
	}
}

func TestPadArtRows_ExactWidthUnchanged(t *testing.T) {
	input := []string{"abcd", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)
	if result[0] != "abcd" {
		t.Errorf("row at exact width must be unchanged: got %q", result[0])
	}
}

func TestPadArtRows_EmptyRowPaddedToWidth(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)
	for i, row := range result {
		if row != "    " {
			t.Errorf("row %d: empty row should pad to 4 spaces, got %q", i, row)
		}
	}
}

func TestPadArtRows_LeadingSpacesPreserved(t *testing.T) {
	// Leading spaces are art — must not be touched
	input := []string{"  _  ", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 8)
	if !strings.HasPrefix(result[0], "  _  ") {
		t.Errorf("leading spaces removed: got %q", result[0])
	}
}

func TestPadArtRows_LengthAlwaysPreserved(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh"}
	result := PadArtRows(input, 6)
	if len(result) != len(input) {
		t.Errorf("number of rows changed: got %d, want %d", len(result), len(input))
	}
}

func TestPadArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi", "a", "b", "c", "d", "e", "f", "g"}
	originals := make([]string, len(input))
	copy(originals, input)
	PadArtRows(input, 10)
	for i := range input {
		if input[i] != originals[i] {
			t.Errorf("row %d: input was mutated — must not modify input slice", i)
		}
	}
}

func TestPadArtRows_ZeroWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PadArtRows panicked with width=0: %v", r)
		}
	}()
	input := []string{"hi", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 0)
	if len(result) != 8 {
		t.Errorf("expected 8 rows returned, got %d", len(result))
	}
}

func TestPadArtRows_NegativeWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PadArtRows panicked with width=-5: %v", r)
		}
	}()
	input := []string{"hi", "", "", "", "", "", "", ""}
	PadArtRows(input, -5)
}

func TestPadArtRows_AllRowsSameWidthAfterPadding(t *testing.T) {
	input := []string{"a", "bb", "ccc", "d", "ee", "f", "ggg", "hh"}
	result := PadArtRows(input, 10)
	for i, row := range result {
		if len(row) != 10 {
			t.Errorf("row %d: expected length 10 after padding, got %d (%q)", i, len(row), row)
		}
	}
}
