package main

import (
	"errors"
	"os"
	"testing"
)

// ─── ValidateArgCount ────────────────────────────────────────────────────────

func TestValidateArgCount(t *testing.T) {
	// Save the real os.Args and restore it after every sub-test
	original := os.Args
	defer func() { os.Args = original }()

	tests := []struct {
		name    string
		args    []string // simulated os.Args
		wantErr error    // nil means we expect success
	}{
		{
			name:    "no arguments",
			args:    []string{"prog"},
			wantErr: ErrTooFewArgs,
		},
		{
			name:    "one argument (string only)",
			args:    []string{"prog", "Hello"},
			wantErr: nil,
		},
		{
			name:    "two arguments (string + font)",
			args:    []string{"prog", "Hello", "shadow"},
			wantErr: nil,
		},
		{
			name:    "three arguments (too many)",
			args:    []string{"prog", "Hello", "shadow", "extra"},
			wantErr: ErrTooManyArgs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := ValidateArgCount()
			if tt.wantErr == nil && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got: %v", tt.wantErr, err)
			}
		})
	}
}

// ─── ValidateInput ───────────────────────────────────────────────────────────

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: ErrEmptyInput,
		},
		{
			name:    "only one newline marker",
			input:   `\n`,
			wantErr: ErrOnlyNewlines,
		},
		{
			name:    "multiple newline markers only",
			input:   `\n\n\n`,
			wantErr: ErrOnlyNewlines,
		},
		{
			name:    "normal text",
			input:   "Hello",
			wantErr: nil,
		},
		{
			name:    "text with newline marker",
			input:   `Hello\nWorld`,
			wantErr: nil,
		},
		{
			name:    "single space",
			input:   " ",
			wantErr: nil,
		},
		{
			name:    "numbers",
			input:   "12345",
			wantErr: nil,
		},
		{
			name:    "special printable chars",
			input:   "!@#$%",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input)
			if tt.wantErr == nil && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got: %v", tt.wantErr, err)
			}
		})
	}
}

// ─── ValidateFontName ────────────────────────────────────────────────────────

func TestValidateFontName(t *testing.T) {
	tests := []struct {
		name     string
		fontName string
		wantPath string
		wantErr  error
	}{
		{
			name:     "standard font",
			fontName: "standard",
			wantPath: "standard.txt",
			wantErr:  nil,
		},
		{
			name:     "shadow font",
			fontName: "shadow",
			wantPath: "shadow.txt",
			wantErr:  nil,
		},
		{
			name:     "thinkertoy font",
			fontName: "thinkertoy",
			wantPath: "thinkertoy.txt",
			wantErr:  nil,
		},
		{
			name:     "unknown font",
			fontName: "comic",
			wantPath: "",
			wantErr:  ErrUnknownFont,
		},
		{
			name:     "empty font name",
			fontName: "",
			wantPath: "",
			wantErr:  ErrUnknownFont,
		},
		{
			name:     "uppercase font name (case-sensitive)",
			fontName: "Standard",
			wantPath: "",
			wantErr:  ErrUnknownFont,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := ValidateFontName(tt.fontName)

			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
				if gotPath != tt.wantPath {
					t.Errorf("expected path %q, got %q", tt.wantPath, gotPath)
				}
			} else {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("expected error %v, got: %v", tt.wantErr, err)
				}
			}
		})
	}
}

// ─── ParseArgs ───────────────────────────────────────────────────────────────

func TestParseArgs(t *testing.T) {
	original := os.Args
	defer func() { os.Args = original }()

	tests := []struct {
		name         string
		args         []string
		wantInput    string
		wantFontPath string
		wantErr      bool
	}{
		{
			name:         "string only — default font",
			args:         []string{"prog", "Hello"},
			wantInput:    "Hello",
			wantFontPath: "standard.txt",
			wantErr:      false,
		},
		{
			name:         "string + shadow font",
			args:         []string{"prog", "Hi", "shadow"},
			wantInput:    "Hi",
			wantFontPath: "shadow.txt",
			wantErr:      false,
		},
		{
			name:    "no args",
			args:    []string{"prog"},
			wantErr: true,
		},
		{
			name:    "unknown font",
			args:    []string{"prog", "Hi", "comic"},
			wantErr: true,
		},
		{
			name:    "empty input",
			args:    []string{"prog", ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			input, fontPath, err := ParseArgs()

			if tt.wantErr {
				if err == nil {
					t.Error("expected an error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if input != tt.wantInput {
				t.Errorf("input: got %q, want %q", input, tt.wantInput)
			}
			if fontPath != tt.wantFontPath {
				t.Errorf("fontPath: got %q, want %q", fontPath, tt.wantFontPath)
			}
		})
	}
}
