package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Banner represents a mapping from printable ASCII characters (32..126) to 8 lines of ASCII art.
type Banner map[rune][]string

var (
	ErrBannerNotFound = errors.New("banner file not found")
	ErrInvalidBanner  = errors.New("invalid banner file format")
)

// ResolveBannerPath returns the path to the banner file, checking common locations.
func ResolveBannerPath(name string) string {
	if !strings.HasSuffix(name, ".txt") {
		name += ".txt"
	}
	candidates := []string{
		name,
		filepath.Join(".", name),
		filepath.Join("renderer", name),
	}
	for _, p := range candidates {
		if info, err := os.Stat(p); err == nil && !info.IsDir() {
			return p
		}
	}
	return name
}

// LoadBanner reads and parses a banner file by name or path.
func LoadBanner(nameOrPath string) (Banner, error) {
	filePath := ResolveBannerPath(nameOrPath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBannerNotFound, nameOrPath)
	}

	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	lines := strings.Split(text, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) < 95*9 {
		return nil, fmt.Errorf("%w: %s has %d lines, expected at least %d", ErrInvalidBanner, filePath, len(lines), 95*9)
	}

	banner := make(Banner, 95)
	for i := 0; i < 95; i++ {
		ch := rune(32 + i)
		start := i*9 + 1
		charLines := make([]string, 8)
		copy(charLines, lines[start:start+8])
		banner[ch] = charLines
	}

	return banner, nil
}
