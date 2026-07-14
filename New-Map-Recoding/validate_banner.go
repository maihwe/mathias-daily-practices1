package main

import (
	"errors"
)

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return errors.New("baaner is nil")
	}
	if len(banner) == 0 {
		return errors.New("empty bannerfile")
	}
	if len(banner) != 95 {
		return errors.New("Invalid file")
	}
	for r, art := range banner {
		if len(art) != 8 {
			return errors.New("'A', got: Invalid content of art")
		}
		if r < ' ' || r > '~' {
			return errors.New("Invalid file 'A'")
		}
	}
	return nil
}
