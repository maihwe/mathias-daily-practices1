package main

import "testing"

func TestLoadBanner_ValidFile(t *testing.T) {
    banner, err := LoadBanner("standard.txt")
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if len(banner) != 95 {
        t.Errorf("expected 95 characters got %d", len(banner))
    }
}

func TestLoadBanner_EachCharHas8Lines(t *testing.T) {
    banner, _ := LoadBanner("standard.txt")
    for char, lines := range banner {
        if len(lines) != 8 {
            t.Errorf("character '%c' has %d lines, expected 8", char, len(lines))
        }
    }
}

func TestLoadBanner_FileNotFound(t *testing.T) {
    _, err := LoadBanner("notfound.txt")
    if err == nil {
        t.Error("expected error for missing file got nil")
    }
}

func TestLoadBanner_SpaceChar(t *testing.T) {
    banner, _ := LoadBanner("standard.txt")
    lines, ok := banner[' ']
    if !ok {
        t.Error("expected space character in banner")
    }
    if len(lines) != 8 {
        t.Errorf("expected 8 lines for space got %d", len(lines))
    }
}
