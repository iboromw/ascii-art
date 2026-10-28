package main

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Hello followed by newline",
			input: "Hello\n",
			want: lines(
				" _    _          _   _          ",
				"| |  | |        | | | |         ",
				"| |__| |   ___  | | | |   ___   ",
				"|  __  |  / _ \\ | | | |  / _ \\  ",
				"| |  | | |  __/ | | | | | (_) | ",
				"|_|  |_|  \\___| |_| |_|  \\___/  ",
				"                                ",
				"                                ",
				"",
			),
		},
		{
			name:  "lowercase hello",
			input: "hello",
			want: lines(
				" _              _   _          ",
				"| |            | | | |         ",
				"| |__     ___  | | | |   ___   ",
				"|  _ \\   / _ \\ | | | |  / _ \\  ",
				"| | | | |  __/ | | | | | (_) | ",
				"|_| |_|  \\___| |_| |_|  \\___/  ",
				"                               ",
				"                               ",
			),
		},
		{
			name:  "mixed case",
			input: "HeLlO",
			want: lines(
				" _    _          _        _    ____   ",
				"| |  | |        | |      | |  / __ \\  ",
				"| |__| |   ___  | |      | | | |  | | ",
				"|  __  |  / _ \\ | |      | | | |  | | ",
				"| |  | | |  __/ | |____  | | | |__| | ",
				"|_|  |_|  \\___| |______| |_|  \\____/  ",
				"                                      ",
				"                                      ",
			),
		},
		{
			name:  "Hello There",
			input: "Hello There",
			want: lines(
				" _    _          _   _                 _______   _                           ",
				"| |  | |        | | | |               |__   __| | |                          ",
				"| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  ",
				"|  __  |  / _ \\ | | | |  / _ \\           | |    |  _ \\   / _ \\ | '__|  / _ \\ ",
				"| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ ",
				"|_|  |_|  \\___| |_| |_|  \\___/           |_|    |_| |_|  \\___| |_|     \\___| ",
				"                                                                             ",
				"                                                                             ",
			),
		},
		{
			name:  "numbers and words",
			input: "1Hello 2There",
			want: lines(
				"     _    _          _   _                         _______   _                           ",
				" _  | |  | |        | | | |                ____   |__   __| | |                          ",
				"/ | | |__| |   ___  | | | |   ___         |___ \\     | |    | |__     ___   _ __    ___  ",
				"| | |  __  |  / _ \\ | | | |  / _ \\          __) |    | |    |  _ \\   / _ \\ | '__|  / _ \\ ",
				"| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ ",
				"|_| |_|  |_|  \\___| |_| |_|  \\___/        |_____|    |_|    |_| |_|  \\___| |_|     \\___| ",
				"                                                                                         ",
				"                                                                                         ",
			),
		},
		{
			name:  "braces around words",
			input: "{Hello There}",
			want: lines(
				"   __  _    _          _   _                 _______   _                           __    ",
				"  / / | |  | |        | | | |               |__   __| | |                          \\ \\   ",
				" | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  ",
				"/ /   |  __  |  / _ \\ | | | |  / _ \\           | |    |  _ \\   / _ \\ | '__|  / _ \\   \\ \\ ",
				"\\ \\   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / ",
				" | |  |_|  |_|  \\___| |_| |_|  \\___/           |_|    |_| |_|  \\___| |_|     \\___|  | |  ",
				"  \\_\\                                                                              /_/   ",
				"                                                                                         ",
			),
		},
		{
			name:  "Hello and There on separate lines",
			input: "Hello\nThere",
			want: lines(
				" _    _          _   _          ",
				"| |  | |        | | | |         ",
				"| |__| |   ___  | | | |   ___   ",
				"|  __  |  / _ \\ | | | |  / _ \\  ",
				"| |  | | |  __/ | | | | | (_) | ",
				"|_|  |_|  \\___| |_| |_|  \\___/  ",
				"                                ",
				"                                ",
				" _______   _                           ",
				"|__   __| | |                          ",
				"   | |    | |__     ___   _ __    ___  ",
				"   | |    |  _ \\   / _ \\ | '__|  / _ \\ ",
				"   | |    | | | | |  __/ | |    |  __/ ",
				"   |_|    |_| |_|  \\___| |_|     \\___| ",
				"                                       ",
				"                                       ",
			),
		},
		{
			name:  "Hello and There separated by two newlines",
			input: "Hello\n\nThere",
			want: lines(
				" _    _          _   _          ",
				"| |  | |        | | | |         ",
				"| |__| |   ___  | | | |   ___   ",
				"|  __  |  / _ \\ | | | |  / _ \\  ",
				"| |  | | |  __/ | | | | | (_) | ",
				"|_|  |_|  \\___| |_| |_|  \\___/  ",
				"                                ",
				"                                ",
				"",
				" _______   _                           ",
				"|__   __| | |                          ",
				"   | |    | |__     ___   _ __    ___  ",
				"   | |    |  _ \\   / _ \\ | '__|  / _ \\ ",
				"   | |    | | | | |  __/ | |    |  __/ ",
				"   |_|    |_| |_|  \\___| |_|     \\___| ",
				"                                       ",
				"                                       ",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert(tt.input)

			if got != tt.want {
				t.Errorf("Convert(%q) mismatch\nwant:\n%q\ngot:\n%q",
					tt.input, tt.want, got)
			}
		})
	}
}

func TestRenderLineRows(t *testing.T) {
	banner, err := LoadBanner("standard")
	if err != nil {
		t.Fatalf("failed to load standard banner: %v", err)
	}

	rows := RenderLine("A", banner)
	if len(rows) != 8 {
		t.Fatalf("RenderLine(\"A\") returned %d rows, want 8", len(rows))
	}
}

func TestRenderShadowAndThinkertoy(t *testing.T) {
	for _, bannerName := range []string{"shadow", "thinkertoy"} {
		banner, err := LoadBanner(bannerName)
		if err != nil {
			t.Fatalf("failed to load %s banner: %v", bannerName, err)
		}
		rows := RenderLine("Test", banner)
		if len(rows) != 8 {
			t.Fatalf("RenderLine(\"Test\", %s) returned %d rows, want 8", bannerName, len(rows))
		}
	}
}

// lines joins output rows with newlines and adds the final newline.
func lines(rows ...string) string {
	return strings.Join(rows, "\n") + "\n"
}
