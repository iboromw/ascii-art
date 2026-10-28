package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		return
	}

	rawInput := os.Args[1]
	if rawInput == "" {
		return
	}

	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	if err := ValidateInput(rawInput); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	banner, err := LoadBanner(bannerName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	lines := ParseInput(rawInput)
	if len(lines) == 0 {
		return
	}

	output := RenderLines(lines, banner)
	fmt.Print(output)
}
