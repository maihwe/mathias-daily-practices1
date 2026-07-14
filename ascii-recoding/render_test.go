package main

import (
	"testing"
)

// ─── Test helper ─────────────────────────────────────────────────────────────

// makeSimpleBannerMap creates a minimal in-memory banner map where each
// character's art rows are just the character itself repeated.
// Used by render and generate tests so we don't need real font files.
func makeSimpleBannerMap() map[rune][]string {
	m := make(map[rune][]string, totalChars)
	for i := 0; i < totalChars; i++ {
		ch := rune(i + firstASCII)
		rows := make([]string, rowsPerChar)
		for r := range rows {
			rows[r] = string(ch) // every row is just the character itself
		}
		m[ch] = rows
	}
	return m
}

// ─── filterPrintable ─────────────────────────────────────────────────────────

func TestFilterPrintable_AllPrintable(t *testing.T) {
	got := filterPrintable("Hello!")
	want := []rune("Hello!")
	if len(got) != len(want) {
		t.Fatalf("length: got %d, want %d", len(got), len(want))
	}
	for i, r := range want {
		if got[i] != r {
			t.Errorf("index %d: got %q, want %q", i, got[i], r)
		}
	}
}

func TestFilterPrintable_EmptyString(t *testing.T) {
	got := filterPrintable("")
	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestFilterPrintable_NonPrintableCharsRemoved(t *testing.T) {
	// Tab (\t = 9), newline (\n = 10), DEL (127) are all outside 32–126
	input := "A\tB\nC\x7fD"
	got := filterPrintable(input)
	want := []rune("ABCD")
	if len(got) != len(want) {
		t.Fatalf("length: got %d, want %d — result: %q", len(got), len(want), string(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFilterPrintable_SpaceIsKept(t *testing.T) {
	got := filterPrintable(" ")
	if len(got) != 1 || got[0] != ' ' {
		t.Errorf("space should be kept, got %v", got)
	}
}

func TestFilterPrintable_TildeIsKept(t *testing.T) {
	got := filterPrintable("~")
	if len(got) != 1 || got[0] != '~' {
		t.Errorf("tilde should be kept, got %v", got)
	}
}

// ─── BuildRow ────────────────────────────────────────────────────────────────

func TestBuildRow_ConcatenatesRows(t *testing.T) {
	bm := makeSimpleBannerMap()
	// In our mock map every row of 'A' is "A" and every row of 'B' is "B"
	result := BuildRow([]rune("AB"), bm, 0)
	if result != "AB" {
		t.Errorf("got %q, want %q", result, "AB")
	}
}

func TestBuildRow_EmptyInput(t *testing.T) {
	bm := makeSimpleBannerMap()
	result := BuildRow([]rune{}, bm, 0)
	if result != "" {
		t.Errorf("got %q, want empty string", result)
	}
}

func TestBuildRow_SkipsMissingCharacters(t *testing.T) {
	bm := make(map[rune][]string) // empty map — no characters
	result := BuildRow([]rune("ABC"), bm, 0)
	if result != "" {
		t.Errorf("expected empty output for missing chars, got %q", result)
	}
}

func TestBuildRow_AllRows(t *testing.T) {
	bm := makeSimpleBannerMap()
	// Check all 8 rows return the same result with our mock (rows are identical)
	for row := 0; row < rowsPerChar; row++ {
		result := BuildRow([]rune("Hi"), bm, row)
		if result != "Hi" {
			t.Errorf("row %d: got %q, want %q", row, result, "Hi")
		}
	}
}

// ─── HasChar ─────────────────────────────────────────────────────────────────

func TestHasChar(t *testing.T) {
	printable := []rune{' ', '!', 'A', 'z', '~'}
	for _, ch := range printable {
		if !HasChar(ch) {
			t.Errorf("HasChar(%q) = false, want true", ch)
		}
	}

	notPrintable := []rune{'\t', '\n', 0, 127, 200}
	for _, ch := range notPrintable {
		if HasChar(ch) {
			t.Errorf("HasChar(%q = %d) = true, want false", ch, ch)
		}
	}
}