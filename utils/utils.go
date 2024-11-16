package utils

import "strings"

func DrawTheGameBox(rows int, cols int, character string) string {
	// Ensure that the character has exactly one symbol
	if len(character) != 1 || rows <= 0 || cols <= 0 {
		return ""
	}

	var box strings.Builder

	// Draw the top border
	box.WriteString(strings.Repeat(character, cols) + "\n")

	// Draw the sides of the box (rows - 2 lines in between)
	for i := 0; i < rows-2; i++ {
		box.WriteString(character + strings.Repeat(" ", cols-2) + character + "\n")
	}

	// Draw the bottom border if there's more than 1 row
	if rows > 1 {
		box.WriteString(strings.Repeat(character, cols) + "\n")
	}

	return box.String()
}
