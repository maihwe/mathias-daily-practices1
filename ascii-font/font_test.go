package asciifont

import "testing"

func TestGenerateFont(t *testing.T) {
    font := GenerateFont()

    if len(font) != 95 {
        t.Errorf("Expected 95 printable ASCII characters, got %d", len(font))
    }

    // Check that each character has 8 lines
    for ch, art := range font {
        if len(art) != 8 {
            t.Errorf("Character %q does not have 8 lines", ch)
        }
        for _, line := range art {
            if len(line) != 8 {
                t.Errorf("Character %q has a line with length %d instead of 8", ch, len(line))
            }
        }
    }
}
