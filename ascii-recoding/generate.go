package main

import (
	"fmt"
	"strings"
)

// ─── Generate ────────────────────────────────────────────────────────────────

// Generate is the top-level function that ties everything together.
// It loads the banner file, splits the input on "\\n", and prints
// the resulting ASCII art to stdout.
//
// Parameters
//
//	input    — raw input string, e.g. "Hello\\nWorld"
//	fontPath — path to the banner file, e.g. "standard.txt"
//
// Returns an error if the banner file cannot be loaded.
func Generate(input, fontPath string) error {
	// Step 1: load the banner map (rune → 8 art rows)
	bannerMap, err := LoadBanner(fontPath)
	if err != nil {
		return fmt.Errorf("Generate: %w", err)
	}

	// Step 2: split input on the literal two-character sequence \n
	// (NOT a real newline — the user types backslash + n)
	chunks := strings.Split(input, `\n`)

	// Step 3: print ASCII art
	Render(chunks, bannerMap)
	return nil
}

// ─── GenerateToString ────────────────────────────────────────────────────────

// GenerateToString returns the ASCII art as a plain string instead of
// printing it. This is mainly useful for tests and for capturing output.
//
// Parameters
//
//	input     — raw input string
//	bannerMap — map returned by LoadBanner (passed in so the caller
//	            controls which font is used without touching the filesystem)
//
// Returns the full ASCII art block as a single string with embedded newlines.
func GenerateToString(input string, bannerMap map[rune][]string) string {
	var sb strings.Builder
	chunks := strings.Split(input, `\n`)
	lastIdx := len(chunks) - 1

	for i, chunk := range chunks {
		if chunk == "" {
			if i < lastIdx {
				sb.WriteString("\n")
			}
			continue
		}

		printable := filterPrintable(chunk)
		if len(printable) == 0 {
			continue
		}

		for row := 0; row < rowsPerChar; row++ {
			sb.WriteString(BuildRow(printable, bannerMap, row))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// ─── GenerateLines ───────────────────────────────────────────────────────────

// GenerateLines returns the ASCII art as a []string (one entry per output line).
// Useful when you need to post-process or align the output.
func GenerateLines(input string, bannerMap map[rune][]string) []string {
	full := GenerateToString(input, bannerMap)
	if full == "" {
		return nil
	}
	// Trim the trailing newline before splitting so we don't get a spurious
	// empty string at the end of the slice.
	trimmed := strings.TrimSuffix(full, "\n")
	return strings.Split(trimmed, "\n")
}
