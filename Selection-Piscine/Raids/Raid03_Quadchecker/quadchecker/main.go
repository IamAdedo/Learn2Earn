package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func makeQuad(x, y int, tl, tr, bl, br, edge, vert rune) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var sb strings.Builder

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			switch {
			case row == 0 && col == 0:
				sb.WriteRune(tl)
			case row == 0 && col == x-1:
				sb.WriteRune(tr)
			case row == y-1 && col == 0:
				sb.WriteRune(bl)
			case row == y-1 && col == x-1:
				sb.WriteRune(br)
			case row == 0 || row == y-1:
				sb.WriteRune(edge)
			case col == 0 || col == x-1:
				sb.WriteRune(vert)
			default:
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func getDimensions(s string) (int, int) {
	if len(s) == 0 {
		return 0, 0
	}

	// Split into lines
	lines := strings.Split(s, "\n")

	// makeQuad always ends with \n, so Split produces a trailing empty string
	// e.g. "a\nb\n" -> ["a", "b", ""]
	// Strip exactly one trailing empty string if present
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	height := len(lines)
	if height == 0 {
		return 0, 0
	}

	width := len(lines[0])
	if width == 0 {
		return 0, 0
	}

	// All lines must have the same width
	for _, line := range lines {
		if len(line) != width {
			return 0, 0
		}
	}

	return width, height
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}

	input := string(data)

	x, y := getDimensions(input)

	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	matches := []string{}

	if input == makeQuad(x, y, 'o', 'o', 'o', 'o', '-', '|') {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if input == makeQuad(x, y, '/', '\\', '\\', '/', '*', '*') {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if input == makeQuad(x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if input == makeQuad(x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if input == makeQuad(x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}
