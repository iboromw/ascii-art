package main

import (
	"os/exec"
	"strings"
	"testing"
)

// The 8 lines of art for "Hello" in the standard banner.
var helloRows = []string{
	" _    _          _   _          ",
	"| |  | |        | | | |         ",
	"| |__| |   ___  | | | |   ___   ",
	"|  __  |  / _ \\ | | | |  / _ \\  ",
	"| |  | | |  __/ | | | | | (_) | ",
	"|_|  |_|  \\___| |_| |_|  \\___/  ",
	"                                ",
	"                                ",
}

// The 8 lines of art for "There" in the standard banner.
var thereRows = []string{
	" _______   _                           ",
	"|__   __| | |                          ",
	"   | |    | |__     ___   _ __    ___  ",
	"   | |    |  _ \\   / _ \\ | '__|  / _ \\ ",
	"   | |    | | | | |  __/ | |    |  __/ ",
	"   |_|    |_| |_|  \\___| |_|     \\___| ",
	"                                       ",
	"                                       ",
}

// join puts every row on its own line, each ending with a newline.
func join(rows []string) string {
	return strings.Join(rows, "\n") + "\n"
}

// runProgram runs: go run . <args...> and returns what it printed to stdout.
func runProgram(t *testing.T, args ...string) string {
	t.Helper()
	cmdArgs := append([]string{"run", "."}, args...)
	cmd := exec.Command("go", cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("go %v failed: %v", cmdArgs, err)
	}
	return string(out)
}

func TestProgramOutput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", "", ""},
		{"only a newline", "\\n", "\n"},
		{"multiple newlines", "\\n\\n", "\n\n"},
		{"Hello with newline", "Hello\\n", join(helloRows) + "\n"},
		{"Hello and There", "Hello\\nThere", join(helloRows) + join(thereRows)},
		{"blank line between", "Hello\\n\\nThere", join(helloRows) + "\n" + join(thereRows)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runProgram(t, tt.input)
			if got != tt.want {
				t.Errorf("input %q\nwant: %q\ngot:  %q", tt.input, tt.want, got)
			}
		})
	}
}

func TestBannerSelectionCLI(t *testing.T) {
	banners := []string{"standard", "shadow", "thinkertoy"}
	for _, b := range banners {
		t.Run(b, func(t *testing.T) {
			out := runProgram(t, "Hello", b)
			if len(out) == 0 {
				t.Errorf("expected non-empty output for banner %s", b)
			}
		})
	}
}

func TestInvalidCharacterCLI(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello é")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("expected non-zero exit for non-ASCII input, got output: %s", string(out))
	}
}

func TestInvalidBannerCLI(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello", "does_not_exist")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("expected non-zero exit for missing banner, got output: %s", string(out))
	}
}
