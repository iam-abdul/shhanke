package state

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Coordinate struct {
	Row       int
	Col       int
	Character rune
}

const (
	RefreshRate     time.Duration = 500 * time.Millisecond
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

	LastDirection string

	// movement
	UpVotes    uint32
	DownVotes  uint32
	LeftVotes  uint32
	RightVotes uint32
}

type RefreshTickMsg time.Time

var SingleGameState *GameState

const rows, cols = 40, 120

func init() {

	snakeInit := make([]Coordinate, 1)
	snakeInit[0] = Coordinate{
		Row:       20,
		Col:       20,
		Character: '▇',
	}

	SingleGameState = &GameState{
		Content:       make([][]rune, rows),
		Snake:         snakeInit,
		LastDirection: "Up",
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

	// Update the grid with snake positions
	for _, pos := range gs.Snake {
		gs.Content[pos.Row][pos.Col] = '▇'
	}

	// reset the votes
	atomic.StoreUint32(&gs.UpVotes, 0)
	atomic.StoreUint32(&gs.DownVotes, 0)
	atomic.StoreUint32(&gs.LeftVotes, 0)
	atomic.StoreUint32(&gs.RightVotes, 0)
}

// Function to find the largest vote and its direction
func (gs *GameState) getLargestVote() (uint32, string) {
	// Map directions to their respective votes
	votes := map[string]uint32{
		"Up":    gs.UpVotes,
		"Down":  gs.DownVotes,
		"Left":  gs.LeftVotes,
		"Right": gs.RightVotes,
	}

	// Iterate to find the largest vote
	var largestVote uint32
	// default direction
	var direction string = gs.LastDirection
	for dir, vote := range votes {
		if vote > largestVote {
			largestVote = vote
			direction = dir
		}
	}

	gs.LastDirection = direction

	return largestVote, direction
}

func (gs *GameState) MoveUp() {
	atomic.AddUint32(&gs.UpVotes, 1)
}

func (gs *GameState) MoveDown() {
	atomic.AddUint32(&gs.DownVotes, 1)
}

func (gs *GameState) MoveRight() {
	atomic.AddUint32(&gs.RightVotes, 1)
}

func (gs *GameState) MoveLeft() {
	atomic.AddUint32(&gs.LeftVotes, 1)
}

func (gs *GameState) MoveSnake() {
	// this looks at direction votes and updates the snake co ordinates
	_, direction := gs.getLargestVote()
	var newHead Coordinate

	currentHead := gs.Snake[0]
	if direction == "Right" {
		newHead = Coordinate{
			Row: currentHead.Row,
			Col: currentHead.Col + 1,
		}
	} else if direction == "Left" {
		newHead = Coordinate{
			Row: currentHead.Row,
			Col: currentHead.Col - 1,
		}
	} else if direction == "Up" {
		newHead = Coordinate{
			Row: currentHead.Row - 1,
			Col: currentHead.Col,
		}
	} else if direction == "Down" {
		newHead = Coordinate{
			Row: currentHead.Row + 1,
			Col: currentHead.Col,
		}
	}

	// Add the new head to the front of the snake
	gs.Snake = append([]Coordinate{newHead}, gs.Snake...)

	// Remove the last block of the snake to maintain its size
	// clear the tail character
	lengthOfSnake := len(gs.Snake)
	tailOfSnake := gs.Snake[lengthOfSnake-1]
	gs.Content[tailOfSnake.Row][tailOfSnake.Col] = ' '
	gs.Snake = gs.Snake[:lengthOfSnake-1]

}

func (gs *GameState) GetContent() string {
	fmt.Println("head is at ", gs.Snake[0])
	gs.MoveSnake()
	gs.updateContent()
	// Build the string representation of the grid
	var result string
	for _, row := range gs.Content {
		for _, cell := range row {
			result += string(cell)
		}
		result += "\n" // Add newline at the end of each row
	}
	fmt.Println("the result is ", result)
	return result
}
