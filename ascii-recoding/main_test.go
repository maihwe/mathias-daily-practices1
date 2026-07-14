package main

import (
	"os"
	"strings"
	"testing"
)

// ─── Full pipeline integration tests ─────────────────────────────────────────
// These tests exercise validate → LoadBanner → Generate together.
// Tests that need real font files are skipped when files are absent,
// so the test suite always passes in CI without font files.

// TestPipeline_ValidInputDefaultFont tests the full flow with a mock font.
func TestPipeline_ValidInputDefaultFont(t *testing.T) {
	bm := makeSimpleBannerMap()

	input := "ABC"
	if err := ValidateInput(input); err != nil {
		t.Fatalf("ValidateInput failed unexpectedly: %v", err)
	}

	result := GenerateToString(input, bm)
	if result == "" {
		t.Error("expected non-empty output for valid input")
	}

	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != rowsPerChar {
		t.Errorf("expected %d lines, got %d", rowsPerChar, len(lines))
	}
}

// TestPipeline_MultiLineInput tests input with \n markers.
func TestPipeline_MultiLineInput(t *testing.T) {
	bm := makeSimpleBannerMap()

	input := `Hello\nWorld`
	if err := ValidateInput(input); err != nil {
		t.Fatalf("ValidateInput failed unexpectedly: %v", err)
	}

	result := GenerateToString(input, bm)
	lines := strings.Split(result, "\n")

	// 8 rows Hello + blank + 8 rows World + trailing newline
	// split gives 18 elements
	if len(lines) != 18 {
		t.Errorf("expected 18 lines, got %d\noutput:\n%s", len(lines), result)
	}
}

// TestPipeline_EmptyInputRejected ensures empty inputs are caught at validate.
func TestPipeline_EmptyInputRejected(t *testing.T) {
	if err := ValidateInput(""); err == nil {
		t.Error("expected error for empty input, got nil")
	}
}

// TestPipeline_OnlyNewlinesRejected ensures \n-only inputs are caught.
func TestPipeline_OnlyNewlinesRejected(t *testing.T) {
	if err := ValidateInput(`\n\n`); err == nil {
		t.Error(`expected error for "\n\n" input, got nil`)
	}
}

// TestPipeline_UnknownFontRejected ensures bad font names are caught.
func TestPipeline_UnknownFontRejected(t *testing.T) {
	_, err := ValidateFontName("wingdings")
	if err == nil {
		t.Error("expected error for unknown font, got nil")
	}
}

// TestPipeline_AllFontsKnown verifies each supported font name resolves.
func TestPipeline_AllFontsKnown(t *testing.T) {
	fonts := []string{"standard", "shadow", "thinkertoy"}
	for _, f := range fonts {
		path, err := ValidateFontName(f)
		if err != nil {
			t.Errorf("font %q should be valid, got error: %v", f, err)
		}
		if path == "" {
			t.Errorf("font %q returned empty path", f)
		}
	}
}

// TestPipeline_NonPrintableInputHandled ensures non-ASCII input does not panic.
func TestPipeline_NonPrintableInputHandled(t *testing.T) {
	bm := makeSimpleBannerMap()
	// Unicode characters outside ASCII 32–126 should be silently skipped.
	// The function should not panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("GenerateToString panicked with non-ASCII input: %v", r)
		}
	}()
	GenerateToString("こんにちは", bm)
}

// TestPipeline_RealStandardFont is a live integration test.
// It is skipped if the font file is absent.
func TestPipeline_RealStandardFont(t *testing.T) {
	if _, err := os.Stat("standard.txt"); os.IsNotExist(err) {
		t.Skip("standard.txt not present — skipping live integration test")
	}

	bm, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner error: %v", err)
	}

	result := GenerateToString("Hello", bm)
	if result == "" {
		t.Error("expected non-empty output for 'Hello' with standard font")
	}
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != rowsPerChar {
		t.Errorf("expected %d lines, got %d", rowsPerChar, len(lines))
	}
}

// TestPipeline_SpaceOnly ensures a single space renders without error.
func TestPipeline_SpaceOnly(t *testing.T) {
	bm := makeSimpleBannerMap()
	result := GenerateToString(" ", bm)
	if result == "" {
		t.Error("expected output for space input")
	}
}

// TestPipeline_TildeChar ensures tilde (last printable ASCII) works correctly.
func TestPipeline_TildeChar(t *testing.T) {
	bm := makeSimpleBannerMap()
	result := GenerateToString("~", bm)
	if result == "" {
		t.Error("expected output for tilde input")
	}
}