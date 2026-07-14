package main

import (
	"strings"
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Setup: Load the banner once for all tests.
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Setup failed: could not load standard.txt: %v", err)
	}

	// Define test scenarios based on project requirements.
	tests := []struct {
		name     string
		input    string
		validate func(t *testing.T, output string)
	}{
		{
			name:  "Empty Input produces nothing",
			input: "",
			validate: func(t *testing.T, output string) {
				if output != "" {
					t.Errorf("Expected empty string, got %q", output)
				}
			},
		},
		{
			name:  "Single newline literal produces one blank line",
			input: `\n`,
			validate: func(t *testing.T, output string) {
				if output != "\n" {
					t.Errorf("Expected exactly one newline, got %q", output)
				}
			},
		},
		{
			name:  "Multiple newline literals produce multiple blank lines",
			input: `\n\n\n`,
			validate: func(t *testing.T, output string) {
				if output != "\n\n\n" {
					t.Errorf("Expected three newlines, got %d", len(output))
				}
			},
		},
		{
			name:  "Mixed text and newlines 'A\\nB'",
			input: `A\nB`,
			validate: func(t *testing.T, output string) {
				// Should have 16 lines total (8 for A, 8 for B)
				lines := strings.Split(output, "\n")
				// strings.Split adds an extra empty element if the string ends in \n
				if len(lines) != 17 {
					t.Errorf("Expected 16 lines of art, found %d", len(lines)-1)
				}
			},
		},
		{
			name:  "Intermediate newline gap 'A\\n\\nB'",
			input: `A\n\nB`,
			validate: func(t *testing.T, output string) {
				// Should have 8 lines (A) + 1 blank line + 8 lines (B) = 17 total printed lines
				lines := strings.Split(output, "\n")
				if len(lines) != 18 {
					t.Errorf("Expected 17 lines total, got %d", len(lines)-1)
				}
				// The 9th line (index 8) should be the empty gap
				if lines[8] != "" {
					t.Errorf("Expected a blank line gap at index 8, got %q", lines[8])
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GenerateArt(tc.input, banner)
			tc.validate(t, result)
		})
	}
}



/*

package main

import (
    "strings"
    "testing"
)

// ---------------------------------------------------------------------------
// Minimal banner fixture
// ---------------------------------------------------------------------------
// Each character is represented by exactly 8 rows, each row ending with a
// space so that RenderLine can concatenate them and then append "\n".
// We use single-character rows so the expected output stays readable.

func makeTestBanner() map[rune][]string {
    // Craft rows for a handful of printable characters.
    // Format mirrors a real banner file: 8 lines per character, each line is
    // the "body" of that row (no trailing newline — GenerateArt adds those).
    charRows := func(ch string) []string {
        rows := make([]string, 8)
        for i := range rows {
            rows[i] = ch // one-char wide cell, easy to spot in diffs
        }
        return rows
    }

    banner := make(map[rune][]string)
    for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 !?" {
        banner[r] = charRows(string(r))
    }
    return banner
}

// expectedArt builds the expected 8-line block for a non-empty segment, each
// line terminated with \n.  It assumes the test banner puts the same single
// character in every row.
func expectedArt(segment string, banner map[rune][]string) string {
    var sb strings.Builder
    for row := 0; row < 8; row++ {
        for _, r := range segment {
            if rows, ok := banner[r]; ok {
                sb.WriteString(rows[row])
            }
        }
        sb.WriteByte('\n')
    }
    return sb.String()
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

func TestGenerateArt_EmptyInput(t *testing.T) {
    // An empty overall input must produce no output at all — not even a newline.
    banner := makeTestBanner()
    got := GenerateArt("", banner)
    if got != "" {
        t.Errorf("empty input: want %q, got %q", "", got)
    }
}

func TestGenerateArt_SingleNewline(t *testing.T) {
    // "\n" means one blank line — exactly one "\n" in the output.
    banner := makeTestBanner()
    got := GenerateArt("\\n", banner)
    want := "\n"
    if got != want {
        t.Errorf("single newline: want %q, got %q", want, got)
    }
}

func TestGenerateArt_SingleWord(t *testing.T) {
    // "Hi" → 8 art lines, each ending with \n, nothing more.
    banner := makeTestBanner()
    got := GenerateArt("Hi", banner)
    want := expectedArt("Hi", banner)
    if got != want {
        t.Errorf("single word:\nwant: %q\ngot:  %q", want, got)
    }
    // Must be exactly 8 lines.
    lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
    if len(lines) != 8 {
        t.Errorf("single word: expected 8 lines, got %d", len(lines))
    }
}

func TestGenerateArt_TwoSegments_NoBlank(t *testing.T) {
    // "A\nB" → 8 lines for A immediately followed by 8 lines for B, no blank
    // line separating them.
    banner := makeTestBanner()
    got := GenerateArt("A\\nB", banner)
    want := expectedArt("A", banner) + expectedArt("B", banner)
    if got != want {
        t.Errorf("A\\nB:\nwant: %q\ngot:  %q", want, got)
    }
    lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
    if len(lines) != 16 {
        t.Errorf("A\\nB: expected 16 lines, got %d", len(lines))
    }
}

func TestGenerateArt_TwoSegments_WithBlank(t *testing.T) {
    // "A\n\nB" → 8 lines for A + exactly 1 blank line + 8 lines for B.
    banner := makeTestBanner()
    got := GenerateArt("A\\n\\nB", banner)
    want := expectedArt("A", banner) + "\n" + expectedArt("B", banner)
    if got != want {
        t.Errorf("A\\n\\nB:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_EmptySegmentProducesOneBlankLine_NotEight(t *testing.T) {
    // This is called out as the most commonly failed case.
    // An empty segment (between two \n\n) must produce exactly ONE blank line,
    // NOT 8 blank lines.
    banner := makeTestBanner()
    got := GenerateArt("A\\n\\nB", banner)

    // Count consecutive blank lines in the output.
    rawLines := strings.Split(got, "\n")
    maxConsecBlanks := 0
    cur := 0
    for _, l := range rawLines {
        if l == "" {
            cur++
            if cur > maxConsecBlanks {
                maxConsecBlanks = cur
            }
        } else {
            cur = 0
        }
    }
    // There will be one genuine blank line from the empty segment, plus the
    // trailing newline of the last art line creates an empty final element
    // after Split — so ≤ 2 consecutive empty strings is acceptable.
    // But 8+ consecutive empty strings means 8 blank lines were emitted.
    if maxConsecBlanks >= 8 {
        t.Errorf("empty segment produced %d consecutive blank lines; want exactly 1", maxConsecBlanks)
    }
}

func TestGenerateArt_MultipleConsecutiveNewlines(t *testing.T) {
    // "A\n\n\nB" has two empty segments between A and B → two blank lines.
    banner := makeTestBanner()
    got := GenerateArt("A\\n\\n\\nB", banner)
    want := expectedArt("A", banner) + "\n" + "\n" + expectedArt("B", banner)
    if got != want {
        t.Errorf("A\\n\\n\\nB:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_SingleCharacter(t *testing.T) {
    banner := makeTestBanner()
    got := GenerateArt("A", banner)
    want := expectedArt("A", banner)
    if got != want {
        t.Errorf("single char A:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_MultipleWords(t *testing.T) {
    // "Hello" is a single segment; should produce exactly 8 art lines.
    banner := makeTestBanner()
    got := GenerateArt("Hello", banner)
    want := expectedArt("Hello", banner)
    if got != want {
        t.Errorf("Hello:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_SpaceCharacter(t *testing.T) {
    // A space is a valid printable character and must render art, not be skipped.
    banner := makeTestBanner()
    got := GenerateArt(" ", banner)
    want := expectedArt(" ", banner)
    if got != want {
        t.Errorf("space char:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_OutputAlwaysEndsWithNewline(t *testing.T) {
    // Every non-empty, non-blank output must end with \n.
    banner := makeTestBanner()
    cases := []string{"A", "AB", "AB", "Hello World"}
    for _, input := range cases {
        got := GenerateArt(input, banner)
        if !strings.HasSuffix(got, "\n") {
            t.Errorf("input %q: output does not end with newline; got %q", input, got)
        }
    }
}

func TestGenerateArt_EachArtLineEndsWithNewline(t *testing.T) {
    // Every one of the 8 art lines for a segment must end with \n.
    banner := makeTestBanner()
    got := GenerateArt("AB", banner)
    lines := strings.Split(got, "\n")
    // After splitting on \n, the last element will be "" (trailing newline).
    // All other elements should be non-empty art lines.
    for i, l := range lines[:len(lines)-1] {
        if l == "" {
            t.Errorf("line %d is unexpectedly blank in output for \"AB\"", i)
        }
    }
}

func TestGenerateArt_InputWithOnlyNewlines(t *testing.T) {
    // "\n\n" → two blank lines.
    banner := makeTestBanner()
    got := GenerateArt("\\n\\n", banner)
    want := "\n\n"
    if got != want {
        t.Errorf("\\n\\n: want %q, got %q", want, got)
    }
}

func TestGenerateArt_LeadingNewline(t *testing.T) {
    // "\nA" → one blank line then 8 art lines for A.
    banner := makeTestBanner()
    got := GenerateArt("\\nA", banner)
    want := "\n" + expectedArt("A", banner)
    if got != want {
        t.Errorf("\\nA:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_TrailingNewline(t *testing.T) {
    // "A\n" → 8 art lines for A then one blank line.
    banner := makeTestBanner()
    got := GenerateArt("A\\n", banner)
    want := expectedArt("A", banner) + "\n"
    if got != want {
        t.Errorf("A\\n:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_ThreeSegments(t *testing.T) {
    // "A\nB\nC" → art for A + art for B + art for C, no blank lines.
    banner := makeTestBanner()
    got := GenerateArt("A\\nB\\nC", banner)
    want := expectedArt("A", banner) + expectedArt("B", banner) + expectedArt("C", banner)
    if got != want {
        t.Errorf("A\\nB\\nC:\nwant: %q\ngot:  %q", want, got)
    }
}

func TestGenerateArt_LiteralBackslashN_NotTreated_As_Newline(t *testing.T) {
    // The spec says the input "may contain literal \n sequences" — meaning the
    // caller passes the real newline rune (\n), not the two-char sequence
    // backslash-n.  This test confirms that a real newline in the Go string
    // acts as a segment separator.
    banner := makeTestBanner()
    input := "A" + "\\n" + "B" // real newline, same as "A\nB"
    got := GenerateArt(input, banner)
    want := expectedArt("A", banner) + expectedArt("B", banner)
    if got != want {
        t.Errorf("real newline separator:\nwant: %q\ngot:  %q", want, got)
    }
}

*/
