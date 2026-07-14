package main

import (
    "reflect"
    "testing"
)

// TestSplitInput uses a table-driven approach to cover all meaningful cases
// for the SplitInput function in one place.
func TestSplitInput(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  []string
    }{
        {
            name:  "no newline sequence",
            input: "Hello",
            want:  []string{"Hello"},
        },
        {
            name:  "single newline between two words",
            input: `Hello\nThere`,
            want:  []string{"Hello", "There"},
        },
        {
            name:  "double newline produces empty middle segment",
            input: `Hello\n\nThere`,
            want:  []string{"Hello", "", "There"},
        },
        {
            name:  "leading newline",
            input: `\nHello`,
            want:  []string{"", "Hello"},
        },
        {
            name:  "trailing newline",
            input: `Hello\n`,
            want:  []string{"Hello", ""},
        },
        {
            name:  "newline only",
            input: `\n`,
            want:  []string{"", ""},
        },
        {
            name:  "double newline only",
            input: `\n\n`,
            want:  []string{"", "", ""},
        },
        {
            name:  "empty string",
            input: "",
            want:  []string{""},
        },
        {
            name:  "three words separated by newlines",
            input: `A\nB\nC`,
            want:  []string{"A", "B", "C"},
        },
        {
            name:  "spaces around newline",
            input: `Hello \n World`,
            want:  []string{"Hello ", " World"},
        },
        {
            name:  "real newline byte is NOT a separator",
            input: "Hello\nThere", // actual newline byte, not backslash-n
            want:  []string{"Hello\nThere"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SplitInput(tt.input)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SplitInput(%q)\n  got:  %q\n  want: %q", tt.input, got, tt.want)
            }
        })
    }
}

// TestSplitInput_NeverReturnsNil ensures the function always returns a
// non-nil slice, even for empty input.
func TestSplitInput_NeverReturnsNil(t *testing.T) {
    result := SplitInput("")
    if result == nil {
        t.Error("SplitInput(\"\") returned nil, expected non-nil slice")
    }
}

// TestSplitInput_SegmentCount verifies the number of segments produced is
// always (number of \\n sequences) + 1.
func TestSplitInput_SegmentCount(t *testing.T) {
    cases := []struct {
        input         string
        wantSegments  int
    }{
        {`A`, 1},
        {`A\nB`, 2},
        {`A\nB\nC`, 3},
        {`\n\n\n`, 4},
    }
    for _, tc := range cases {
        got := SplitInput(tc.input)
        if len(got) != tc.wantSegments {
            t.Errorf("SplitInput(%q): expected %d segments, got %d",
                tc.input, tc.wantSegments, len(got))
        }
    }
}