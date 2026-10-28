package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Single line string",
			input:    "Hello",
			expected: []string{"Hello"},
		},
		{
			name:     "Multi-line string with \\n",
			input:    "Hello\\nWorld",
			expected: []string{"Hello", "World"},
		},
		{
			name:     "String with consecutive \\n",
			input:    "Hello\\n\\nWorld",
			expected: []string{"Hello", "", "World"},
		},
		{
			name:     "Leading \\n",
			input:    "\\nHello",
			expected: []string{"", "Hello"},
		},
		{
			name:     "Trailing \\n",
			input:    "Hello\\n",
			expected: []string{"Hello", ""},
		},
		{
			name:     "Only \\n sequences",
			input:    "\\n\\n",
			expected: []string{"", "", ""},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Space padded string",
			input:    "  Hello World  ",
			expected: []string{"  Hello World  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseInput(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseInput(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid ASCII string",
			input:   "Hello World 123!",
			wantErr: false,
		},
		{
			name:    "Valid ASCII with \\n escape sequence",
			input:   "Hello\\nWorld",
			wantErr: false,
		},
		{
			name:    "Invalid non-ASCII characters",
			input:   "Bonjour éà",
			wantErr: true,
		},
		{
			name:    "Control character (bell)",
			input:   "Hello\x07World",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestIsOnlyNewlines(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected bool
	}{
		{
			name:     "All empty lines",
			lines:    []string{"", "", ""},
			expected: true,
		},
		{
			name:     "Mixed empty and text lines",
			lines:    []string{"Hello", "", "World"},
			expected: false,
		},
		{
			name:     "All text lines",
			lines:    []string{"Hello", "World"},
			expected: false,
		},
		{
			name:     "Empty slice",
			lines:    []string{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsOnlyNewlines(tt.lines)
			if got != tt.expected {
				t.Errorf("IsOnlyNewlines(%v) = %v; want %v", tt.lines, got, tt.expected)
			}
		})
	}
}
