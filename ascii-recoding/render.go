package main

import (
	"fmt"
	"strings"
)

// ─── Render ──────────────────────────────────────────────────────────────────

// Render prints ASCII art to stdout.
//
// Parameters
//
//	chunks    — the input text already split on the literal "\\n" marker.
//	            e.g. "Hello\\nWorld" → ["Hello", "World"]
//	bannerMap — the map returned by LoadBanner
//
// Behaviour
//
//   - Each non-empty chunk is printed as 8 rows of ASCII art.
//   - An empty chunk (produced by "\\n") prints a blank line.
//   - The very last element is NOT followed by an extra newline if it is empty,
//     preventing a double blank line at the end of output.
//   - Characters outside ASCII 32–126 are silently skipped.
func Render(chunks []string, bannerMap map[rune][]string) {
	lastIdx := len(chunks) - 1

	for i, chunk := range chunks {
		if chunk == "" {
			// "\\n" in input → blank line, but skip if it's the trailing element
			if i < lastIdx {
				fmt.Println()
			}
			continue
		}

		// Keep only printable characters (ASCII 32–126)
		printable := filterPrintable(chunk)
		if len(printable) == 0 {
			continue // nothing drawable in this chunk
		}

		// Print all 8 rows for this chunk
		for row := 0; row < rowsPerChar; row++ {
			fmt.Println(BuildRow(printable, bannerMap, row))
		}
	}
}

// ─── BuildRow ────────────────────────────────────────────────────────────────

// BuildRow returns one horizontal row of ASCII art for a slice of characters.
//
// Example: BuildRow([]rune("Hi"), bannerMap, 0)
// returns the top row of 'H' concatenated with the top row of 'i'.
//
// This is exported so generate.go and tests can use it directly.
func BuildRow(chars []rune, bannerMap map[rune][]string, row int) string {
	var sb strings.Builder
	for _, ch := range chars {
		rows, ok := bannerMap[ch]
		if !ok || row >= len(rows) {
			continue // character not in map or row out of range — skip
		}
		sb.WriteString(rows[row])
	}
	return sb.String()
}

// ─── filterPrintable ─────────────────────────────────────────────────────────

// filterPrintable returns only the runes in s that fall within
// the printable ASCII range (32 … 126).
// Characters outside this range are silently dropped.
func filterPrintable(s string) []rune {
	var result []rune
	for _, r := range s {
		if HasChar(r) {
			result = append(result, r)
		}
	}
	return result
}
