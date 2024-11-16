package state

import (
	"time"
)

type Coordinate struct {
	Row       int
	Col       int
	Character rune
}

const (
	RefreshRate     time.Duration = 900 * time.Millisecond
	InputSampleRate time.Duration = 50 * time.Millisecond
)

// game struct should be a singleton
type GameState struct {
	Content               [][]rune
	NumberOfPlayersOnline int
	RightKeyPresses       int
	LeftKeyPresses        int
	UpKeyPresses          int
	DownKeyPresses        int

	Ready bool

	// snake info
	Fruit Coordinate
	Snake []Coordinate
}

type RefreshTickMsg time.Time

var SingleGameState *GameState

const rows, cols = 40, 120

func init() {

	snakeInit := make([]Coordinate, 4)
	snakeInit = append(snakeInit, Coordinate{
		Row:       rows / 2,
		Col:       cols / 2,
		Character: '>',
	})
	snakeInit = append(snakeInit, Coordinate{
		Row:       rows / 2,
		Col:       cols/2 + 1,
		Character: '>',
	})
	snakeInit = append(snakeInit, Coordinate{
		Row:       rows / 2,
		Col:       cols/2 + 2,
		Character: '>',
	})

	SingleGameState = &GameState{
		Content: make([][]rune, rows),
		Snake:   snakeInit,
	}

	for i := range SingleGameState.Content {
		SingleGameState.Content[i] = make([]rune, cols) // Initialize each row with a slice of runes
	}

	// Build the square box using a nested loop
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				SingleGameState.Content[i][j] += '#' // Boundary of the box
			} else {
				SingleGameState.Content[i][j] += ' '
				// content += " " // Inside the box
			}
		}
		// SingleGameState.Content[i][j] += '#'
		// content += "\n" // Newline at the end of each row
	}

	SingleGameState.updateContent()
}

func (gs *GameState) updateContent() {
	// Clear the grid first
	// for i := range gs.Content {
	// 	for j := range gs.Content[i] {
	// 		gs.Content[i][j] = ' ' // Empty space
	// 	}
	// }

	// Update the grid with snake positions
	for _, pos := range gs.Snake {
		gs.Content[pos.Row][pos.Col] = pos.Character
	}

}

func (gs *GameState) GetContent() string {
	// Build the string representation of the grid
	var result string
	for _, row := range gs.Content {
		for _, cell := range row {
			result += string(cell)
		}
		result += "\n" // Add newline at the end of each row
	}
	return result
}
