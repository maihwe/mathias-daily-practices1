package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// ─── Sentinel errors ────────────────────────────────────────────────────────
// Sentinel errors let callers check the *kind* of error with errors.Is()
// without having to compare strings.

var (
	ErrTooFewArgs   = errors.New("too few arguments")
	ErrTooManyArgs  = errors.New("too many arguments")
	ErrEmptyInput   = errors.New("input string is empty")
	ErrOnlyNewlines = errors.New("input contains only newline markers")
	ErrUnknownFont  = errors.New("unknown font name")
)

// allowedFonts maps the short font name the user types → the actual filename.
// Add new fonts here and they become available automatically everywhere.
var allowedFonts = map[string]string{
	"standard":   "standard.txt",
	"shadow":     "shadow.txt",
	"thinkertoy": "thinkertoy.txt",
}

// ─── Argument validation ─────────────────────────────────────────────────────

// ValidateArgCount checks that os.Args has 2 or 3 elements:
//
//	go run . "text"           → 2 args  ✓
//	go run . "text" shadow    → 3 args  ✓
//	go run .                  → 1 arg   ✗
//	go run . a b c            → 4 args  ✗
func ValidateArgCount() error {
	n := len(os.Args)
	switch {
	case n < 2:
		return fmt.Errorf("%w: %s", ErrTooFewArgs, Usage())
	case n > 3:
		return fmt.Errorf("%w: %s", ErrTooManyArgs, Usage())
	}
	return nil
}

// ─── Input validation ────────────────────────────────────────────────────────

// ValidateInput checks that the input string is usable:
//   - Must not be empty
//   - Must contain at least one non-newline character
//     (e.g. "\\n\\n" is rejected — nothing to draw)
func ValidateInput(input string) error {
	if input == "" {
		return ErrEmptyInput
	}
	// Remove all literal "\n" markers and see if anything is left
	stripped := strings.ReplaceAll(input, `\n`, "")
	if stripped == "" {
		return fmt.Errorf("%w: %q", ErrOnlyNewlines, input)
	}
	return nil
}

// ─── Font validation ─────────────────────────────────────────────────────────

// ValidateFontName checks that the requested font exists in allowedFonts
// and returns the corresponding filename.
//
// Example:
//
//	ValidateFontName("shadow")     → "shadow.txt", nil
//	ValidateFontName("comic")      → "", ErrUnknownFont
func ValidateFontName(name string) (string, error) {
	path, ok := allowedFonts[name]
	if !ok {
		return "", fmt.Errorf(
			"%w: %q — available fonts: standard, shadow, thinkertoy",
			ErrUnknownFont, name,
		)
	}
	return path, nil
}

// ─── Convenience helpers used by main ────────────────────────────────────────

// ParseArgs reads os.Args and returns (inputString, fontFilePath, error).
// It validates arg count, input content, and font name in one call.
func ParseArgs() (input string, fontPath string, err error) {
	if err = ValidateArgCount(); err != nil {
		return
	}

	input = os.Args[1]
	if err = ValidateInput(input); err != nil {
		return
	}

	fontName := "standard" // default
	if len(os.Args) == 3 {
		fontName = os.Args[2]
	}

	fontPath, err = ValidateFontName(fontName)
	return
}

// Usage returns the usage string shown on errors.
func Usage() string {
	return "Usage: go run . [STRING] [BANNER]\n" +
		"       Banners: standard | shadow | thinkertoy"
}
