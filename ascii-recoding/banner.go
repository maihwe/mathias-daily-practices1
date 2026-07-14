package main

import (
	"fmt"
	"os"
	"strings"
)

// ─── Constants ───────────────────────────────────────────────────────────────

const (
	firstASCII   = 32  // Space ' '  — first printable ASCII character
	lastASCII    = 126 // Tilde '~'  — last printable ASCII character
	totalChars   = 95  // lastASCII - firstASCII + 1
	rowsPerChar  = 8   // every character is drawn across 8 lines
	linesPerChar = 9   // 8 art lines + 1 blank separator line
)

// ─── LoadBanner ──────────────────────────────────────────────────────────────

// LoadBanner reads a banner file and returns a map where:
//
//	key   = a rune (character), e.g. 'A', ' ', '!'
//	value = []string of exactly 8 lines representing that character's ASCII art
//
// Banner file format
//
//	The file stores 95 characters (ASCII 32 … 126).
//	Each character occupies 9 lines in the file:
//	  line 0  → blank separator  (ignored)
//	  line 1–8 → the 8 rows of ASCII art
//
//	So character index i starts at line i*9, and its art is at i*9+1 … i*9+8.
//
// Example usage
//
//	bannerMap, err := LoadBanner("standard.txt")
//	spaceRows := bannerMap[' ']   // 8 empty/space strings
//	aRows     := bannerMap['A']   // 8 strings drawing the letter A
//
// Errors returned
//
//   - file not found or not readable
//   - file has fewer lines than expected  (corrupt / wrong file)
func LoadBanner(filename string) (map[rune][]string, error) {
	// Read the entire file at once — banner files are small (< 100 KB)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("LoadBanner: cannot read %q: %w", filename, err)
	}

	// Normalise Windows line endings (\r\n → \n) so the parser
	// works correctly on any operating system.
	content := strings.ReplaceAll(string(raw), "\r\n", "\n")

	// Split into individual lines.
	// Note: strings.Split always produces at least one element, even for "".
	lines := strings.Split(content, "\n")

	// Validate that the file is long enough.
	// A valid banner file must have at least totalChars * linesPerChar lines.
	minRequired := totalChars * linesPerChar // 95 * 9 = 855
	if len(lines) < minRequired {
		return nil, fmt.Errorf(
			"LoadBanner: file %q has %d lines, need at least %d (file may be corrupt)",
			filename, len(lines), minRequired,
		)
	}

	// Build the map: rune → 8-line art slice
	bannerMap := make(map[rune][]string, totalChars)

	for i := 0; i < totalChars; i++ {
		char := rune(i + firstASCII) // i=0 → ' '(32), i=1 → '!'(33), …

		startLine := i * linesPerChar
		// lines[startLine]        → blank separator  (skip)
		// lines[startLine+1 … +8] → the 8 art rows
		artStart := startLine + 1
		artEnd := startLine + linesPerChar // exclusive upper bound

		// Deep-copy the 8 rows into a fresh slice so that changes to
		// the lines slice later cannot corrupt the map.
		rows := make([]string, rowsPerChar)
		copy(rows, lines[artStart:artEnd])

		bannerMap[char] = rows
	}

	return bannerMap, nil
}

// HasChar reports whether ch is a printable ASCII character
// that exists in every standard banner map.
func HasChar(ch rune) bool {
	return ch >= firstASCII && ch <= lastASCII
}
