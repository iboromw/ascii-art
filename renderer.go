package main

import (
	"strings"
)

// RenderLine converts a single string of text into 8 horizontal ASCII-art lines using the given banner.
func RenderLine(line string, banner Banner) []string {
	result := make([]string, 8)
	for _, r := range line {
		art, ok := banner[r]
		if !ok {
			continue
		}
		for i := 0; i < 8; i++ {
			result[i] += art[i]
		}
	}
	return result
}

// RenderLines renders a slice of logical input lines into the complete ASCII art string.
func RenderLines(lines []string, banner Banner) string {
	if len(lines) == 0 {
		return ""
	}

	if IsOnlyNewlines(lines) {
		return strings.Repeat("\n", len(lines)-1)
	}

	var sb strings.Builder
	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
		} else {
			artRows := RenderLine(line, banner)
			for _, row := range artRows {
				sb.WriteString(row)
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

// Convert parses the input string and renders it using the standard banner.
func Convert(original string) string {
	if original == "" {
		return ""
	}
	banner, err := LoadBanner("standard")
	if err != nil {
		return ""
	}
	lines := ParseInput(original)
	return RenderLines(lines, banner)
}
