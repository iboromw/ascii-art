package main

import (
	"errors"
	"strings"
)

var ErrInvalidCharacter = errors.New("input contains non-printable ASCII characters")

func ValidateInput(input string) error {
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch == '\\' && i+1 < len(input) && input[i+1] == 'n' {
			i++
			continue
		}
		if ch == '\n' || ch == '\r' {
			continue
		}
		if ch < 32 || ch > 126 {
			return ErrInvalidCharacter
		}
	}
	return nil
}

func ParseInput(input string) []string {
	if input == "" {
		return []string{}
	}
	normalized := strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(normalized, "\n")
}

func IsOnlyNewlines(lines []string) bool {
	if len(lines) == 0 {
		return false
	}
	for _, line := range lines {
		if line != "" {
			return false
		}
	}
	return true
}
