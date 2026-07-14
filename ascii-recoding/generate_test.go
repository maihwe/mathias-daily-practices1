package main

import (
	"os"
	"strings"
	"testing"
)

// ─── GenerateToString ────────────────────────────────────────────────────────

func TestGenerateToString_SimpleText(t *testing.T) {
	bm := makeSimpleBannerMap()
	// Our mock: every row of 'H' is "H", every row of 'i' is "i"
	result := GenerateToString("Hi", bm)

	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != rowsPerChar {
		t.Errorf("expected %d lines, got %d", rowsPerChar, len(lines))
	}
	for i, line := range lines {
		if line != "Hi" {
			t.Errorf("line %d: got %q, want %q", i, line, "Hi")
		}
	}
}

func TestGenerateToString_NewlineMarker(t *testing.T) {
	bm := makeSimpleBannerMap()
	// "A\nB" should produce: 8 rows for A, blank line, 8 rows for B
	result := GenerateToString(`A\nB`, bm)

	lines := strings.Split(result, "\n")
	// 8 rows for A + blank line + 8 rows for B + trailing newline = 18 elements
	// after Split: ["A","A","A","A","A","A","A","A","","B","B","B","B","B","B","B","B",""]
	// = 18 elements
	if len(lines) != 18 {
		t.Errorf("expected 18 lines after split, got %d\nresult:\n%s", len(lines), result)
	}
	// The 9th element (index 8) should be the blank line
	if lines[8] != "" {
		t.Errorf("expected blank line at index 8, got %q", lines[8])
	}
}

func TestGenerateToString_EmptyInput(t *testing.T) {
	bm := makeSimpleBannerMap()
	result := GenerateToString("", bm)
	// Empty input → empty output (nothing to draw)
	if result != "" {
		t.Errorf("expected empty output for empty input, got %q", result)
	}
}

func TestGenerateToString_OnlyNewlineMarkers(t *testing.T) {
	bm := makeSimpleBannerMap()
	// \n\n → two blank lines between nothing
	result := GenerateToString(`\n\n`, bm)
	// chunks = ["", "", ""] → two empty chunks produce blank lines, last skipped
	// Expected: "\n\n" (two newlines)
	if result != "\n\n" {
		t.Errorf("got %q, want %q", result, "\n\n")
	}
}

func TestGenerateToString_NonPrintableCharsSkipped(t *testing.T) {
	bm := makeSimpleBannerMap()
	// Tab and newline (real \n, not the marker) should be skipped
	result := GenerateToString("A\tB", bm)
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	for i, line := range lines {
		// Tab is not printable, so output should be "AB" not "A\tB"
		if line != "AB" {
			t.Errorf("line %d: got %q, want %q", i, line, "AB")
		}
	}
}

func TestGenerateToString_SingleSpace(t *testing.T) {
	bm := makeSimpleBannerMap()
	result := GenerateToString(" ", bm)
	if result == "" {
		t.Error("expected output for space character, got empty string")
	}
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != rowsPerChar {
		t.Errorf("expected %d lines for space, got %d", rowsPerChar, len(lines))
	}
}

// ─── GenerateLines ───────────────────────────────────────────────────────────

func TestGenerateLines_ReturnsSlice(t *testing.T) {
	bm := makeSimpleBannerMap()
	lines := GenerateLines("Hi", bm)
	if len(lines) != rowsPerChar {
		t.Errorf("expected %d lines, got %d", rowsPerChar, len(lines))
	}
}

func TestGenerateLines_EmptyInput(t *testing.T) {
	bm := makeSimpleBannerMap()
	lines := GenerateLines("", bm)
	if lines != nil {
		t.Errorf("expected nil for empty input, got %v", lines)
	}
}

func TestGenerateLines_NoTrailingEmptyLine(t *testing.T) {
	bm := makeSimpleBannerMap()
	lines := GenerateLines("A", bm)
	last := lines[len(lines)-1]
	if last == "" {
		t.Error("GenerateLines should not have a trailing empty string")
	}
}

// ─── Generate (integration — requires font files) ────────────────────────────

func TestGenerate_FileNotFound(t *testing.T) {
	err := Generate("Hello", "notfound.txt")
	if err == nil {
		t.Error("expected error for missing font file, got nil")
	}
}

func TestGenerate_RealFile(t *testing.T) {
	if _, err := os.Stat("standard.txt"); os.IsNotExist(err) {
		t.Skip("standard.txt not found — skipping real-file integration test")
	}
	err := Generate("Hello", "standard.txt")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}