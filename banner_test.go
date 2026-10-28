package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBannerSuccess(t *testing.T) {
	banners := []string{"standard", "shadow", "thinkertoy"}

	for _, name := range banners {
		t.Run(name, func(t *testing.T) {
			banner, err := LoadBanner(name)
			if err != nil {
				t.Fatalf("LoadBanner(%q) unexpected error: %v", name, err)
			}

			// Validate all 95 printable ASCII characters (32..126)
			for ch := ' '; ch <= '~'; ch++ {
				art, exists := banner[ch]
				if !exists {
					t.Fatalf("character %q (ASCII %d) missing from banner %q", ch, ch, name)
				}
				if len(art) != 8 {
					t.Fatalf("character %q in %q has %d rows, want 8", ch, name, len(art))
				}
			}
		})
	}
}

func TestLoadBannerNotFound(t *testing.T) {
	_, err := LoadBanner("non_existent_banner")
	if err == nil {
		t.Fatal("expected error for non-existent banner, got nil")
	}
	if !errors.Is(err, ErrBannerNotFound) {
		t.Errorf("expected ErrBannerNotFound, got %v", err)
	}
}

func TestLoadBannerInvalidFormat(t *testing.T) {
	tmpDir := t.TempDir()
	invalidFile := filepath.Join(tmpDir, "invalid.txt")
	if err := os.WriteFile(invalidFile, []byte("short content\nline2"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	_, err := LoadBanner(invalidFile)
	if err == nil {
		t.Fatal("expected error for invalid banner file, got nil")
	}
	if !errors.Is(err, ErrInvalidBanner) {
		t.Errorf("expected ErrInvalidBanner, got %v", err)
	}
}
