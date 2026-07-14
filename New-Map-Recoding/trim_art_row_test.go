package main

import (
	"strings"
	"testing"
)

func TestTrimArtRows_LengthUnchanged(t *testing.T) {
	input := []string{"a  ", "b", "c   ", "d", "e  ", "f", "g   ", "h"}
	result := TrimArtRows(input)
	if len(result) != len(input) {
		t.Errorf("expected %d rows, got %d", len(input), len(result))
	}
}

func TestTrimArtRows_AllEmptyRows(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := TrimArtRows(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: expected empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_AllSpaceRows(t *testing.T) {
	input := []string{"   ", "  ", " ", "    ", "     ", "  ", "   ", " "}
	result := TrimArtRows(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: all-space row should become empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_NoTrailingSpaces(t *testing.T) {
	// Nothing to trim — rows should come back identical
	input := []string{"_", "| |", "|_|", "", " _", "| |", "|_|", ""}
	result := TrimArtRows(input)
	for i := range input {
		if result[i] != input[i] {
			t.Errorf("row %d: no trailing spaces, must be unchanged. got %q, want %q",
				i, result[i], input[i])
		}
	}
}

func TestTrimArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi   ", "there  ", "a   ", "b  ", "c ", "d  ", "e  ", "f   "}
	originals := make([]string, len(input))
	copy(originals, input)
	TrimArtRows(input)
	for i := range input {
		if input[i] != originals[i] {
			t.Errorf("row %d: input was modified — must return new slice. got %q", i, input[i])
		}
	}
}

func TestTrimArtRows_ReturnsNewSlice(t *testing.T) {
	input := []string{"a  ", "b  ", "c  ", "d  ", "e  ", "f  ", "g  ", "h  "}
	result := TrimArtRows(input)
	// Mutate result — original must be unchanged
	result[0] = "MUTATED"
	if input[0] == "MUTATED" {
		t.Error("TrimArtRows must return a new slice, not a reference to the input")
	}
}

func TestTrimArtRows_MidRowSpacesPreserved(t *testing.T) {
	// Internal spaces in art must never be touched
	input := []string{"| _ |   ", "| | |  ", "|___|", "", "", "", "", ""}
	result := TrimArtRows(input)
	if !strings.Contains(result[0], "| _ |") {
		t.Errorf("row 0: internal spaces removed — must be preserved. got %q", result[0])
	}
	if !strings.Contains(result[1], "| | |") {
		t.Errorf("row 1: internal spaces removed — must be preserved. got %q", result[1])
	}
}

func TestTrimArtRows_EmptySlice(t *testing.T) {
	result := TrimArtRows([]string{})
	if result == nil {
		t.Error("must not return nil for empty input slice")
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got length %d", len(result))
	}
}
