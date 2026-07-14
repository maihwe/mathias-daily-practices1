package maihwe

import (
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
    if banner == nil {
        return fmt.Errorf("banner is nil")
    }

    if len(banner) != 95 {
        return fmt.Errorf("banner has %d entries, expected 95", len(banner))
    }

    for r := rune(32); r <= 126; r++ {
        lines, ok := banner[r]
        if !ok {
            return fmt.Errorf("missing character: %q (ASCII %d)", r, r)
        }
        if len(lines) != 8 {
            return fmt.Errorf("character %q has %d lines, expected 8", r, len(lines))
        }
    }

    return nil
}