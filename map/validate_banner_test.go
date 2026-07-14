package main

import (
    "testing"
)

func TestValidateBannerShort(t *testing.T) {
    if err := ValidateBanner(nil); err == nil {
        t.Errorf("Expected error for nil banner, got nil")
    }

    ValidBanner := make(map[rune][]string)
    for r := ' '; r <= '~'; r++ {
        ValidBanner[r] = make([]string, 8)
    }

    if err := ValidateBanner(ValidBanner); err != nil {
        t.Errorf("Expected banner to pass got erro: %v", err)
    }
}

