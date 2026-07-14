package main

import (
	"os"
	"strings"
	"testing"
)

// ─── Test helper ─────────────────────────────────────────────────────────────

// makeMockBannerFile writes a valid mock banner file to a temp location.
// Every character's art is filled with its own rune repeated, so we can
// assert on the content without needing the real font files.
//
// Format produced:
//
//	(blank line)
//	row1 row2 … row8   ← 8 identical lines per character
//	(blank line)
//	…  (repeated for all 95 characters)
func makeMockBannerFile(t *testing.T) string {
	t.Helper()

	var sb strings.Builder
	for i := 0; i < totalChars; i++ {
		ch := rune(i + firstASCII)
		artLine := strings.Repeat(string(ch), 4) // e.g. "AAAA" for 'A'
		sb.WriteString("\n")                     // blank separator line
		for row := 0; row < rowsPerChar; row++ {
			sb.WriteString(artLine + "\n")
		}
	}

	f, err := os.CreateTemp(t.TempDir(), "banner_*.txt")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(sb.String()); err != nil {
		t.Fatalf("could not write temp file: %v", err)
	}
	return f.Name()
}

// makeTruncatedBannerFile creates an intentionally short (corrupt) banner file.
func makeTruncatedBannerFile(t *testing.T) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "short_banner_*.txt")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	defer f.Close()
	f.WriteString("only a few lines\nnot enough\n")
	return f.Name()
}

// ─── LoadBanner tests ────────────────────────────────────────────────────────

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := LoadBanner("notfound.txt")
	if err == nil {
		t.Error("expected an error for missing file, got nil")
	}
}

func TestLoadBanner_TruncatedFile(t *testing.T) {
	path := makeTruncatedBannerFile(t)
	_, err := LoadBanner(path)
	if err == nil {
		t.Error("expected an error for truncated file, got nil")
	}
}

func TestLoadBanner_Returns95Characters(t *testing.T) {
	path := makeMockBannerFile(t)
	bannerMap, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := len(bannerMap); got != totalChars {
		t.Errorf("expected %d characters in map, got %d", totalChars, got)
	}
}

func TestLoadBanner_EachCharHas8Rows(t *testing.T) {
	path := makeMockBannerFile(t)
	bannerMap, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for ch := rune(firstASCII); ch <= rune(lastASCII); ch++ {
		rows, ok := bannerMap[ch]
		if !ok {
			t.Errorf("character %q (%d) missing from map", ch, ch)
			continue
		}
		if len(rows) != rowsPerChar {
			t.Errorf("character %q: expected %d rows, got %d", ch, rowsPerChar, len(rows))
		}
	}
}

func TestLoadBanner_SpaceCharExists(t *testing.T) {
	path := makeMockBannerFile(t)
	bannerMap, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	rows, ok := bannerMap[' ']
	if !ok {
		t.Fatal("space character ' ' missing from banner map")
	}
	if len(rows) != rowsPerChar {
		t.Errorf("space has %d rows, expected %d", len(rows), rowsPerChar)
	}
}

func TestLoadBanner_TildeCharExists(t *testing.T) {
	path := makeMockBannerFile(t)
	bannerMap, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := bannerMap['~']; !ok {
		t.Error("tilde character '~' missing from banner map")
	}
}

func TestLoadBanner_CorrectArtContent(t *testing.T) {
	// In our mock file, every art row for char 'A' is "AAAA"
	path := makeMockBannerFile(t)
	bannerMap, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	rows := bannerMap['A']
	for i, row := range rows {
		if row != "AAAA" {
			t.Errorf("row %d of 'A': got %q, want %q", i, row, "AAAA")
		}
	}
}

func TestLoadBanner_WindowsLineEndings(t *testing.T) {
	// Write a file with \r\n line endings and ensure it parses correctly.
	var sb strings.Builder
	for i := 0; i < totalChars; i++ {
		ch := rune(i + firstASCII)
		artLine := strings.Repeat(string(ch), 4)
		sb.WriteString("\r\n") // Windows blank separator
		for row := 0; row < rowsPerChar; row++ {
			sb.WriteString(artLine + "\r\n")
		}
	}
	f, err := os.CreateTemp(t.TempDir(), "win_banner_*.txt")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	f.WriteString(sb.String())
	f.Close()

	bannerMap, err := LoadBanner(f.Name())
	if err != nil {
		t.Fatalf("unexpected error on Windows line endings: %v", err)
	}
	if len(bannerMap) != totalChars {
		t.Errorf("expected %d chars, got %d", totalChars, len(bannerMap))
	}
}

// TestLoadBanner_RealFile skips if the file is not present.
// Run from the project root where the .txt files live.
func TestLoadBanner_RealFile(t *testing.T) {
	if _, err := os.Stat("standard.txt"); os.IsNotExist(err) {
		t.Skip("standard.txt not found — skipping real-file test")
	}
	bannerMap, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(bannerMap) != totalChars {
		t.Errorf("expected %d chars, got %d", totalChars, len(bannerMap))
	}
}
