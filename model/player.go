package model

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iam-abdul/go_snake.sh/state"
)

type TestMultiplayer struct {
	Character string

	// here i guess i can store player level info
	// like when he joined and stuff
}

type tickMsg time.Time

// func tick() tea.Msg {
// 	time.Sleep(time.Millisecond * 200)
// 	return tickMsg{}
// }

func tick() tea.Msg {

	interval := state.RefreshRate
	// Get the current time
	now := time.Now()

	// Calculate milliseconds elapsed in the current second
	ms := now.Nanosecond() / 1e6 // Convert nanoseconds to milliseconds

	// Convert the interval to milliseconds
	intervalMs := int(interval.Milliseconds())

	// Calculate how many milliseconds until the next interval boundary
	sleepDuration := (intervalMs - (ms % intervalMs)) % intervalMs

	// Sleep until the next interval
	time.Sleep(time.Millisecond * time.Duration(sleepDuration))

	// Return a tickMsg
	return tickMsg(time.Now())
}
func (T TestMultiplayer) Init() tea.Cmd {
	return tick
}

func (T TestMultiplayer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return T, tea.Quit

		case "up":
			state.SingleGameState.MoveUp()
			return T, nil

		case "down":
			state.SingleGameState.MoveDown()
			return T, nil

		case "left":
			state.SingleGameState.MoveLeft()
			return T, nil

		case "right":
			state.SingleGameState.MoveRight()
			return T, nil

		default:
			// gameState := state.GetGameState().Content
			// state.UpdateGameContent(gameState + msg.String())
			return T, nil
		}

	case tickMsg:
		// tickTime := time.Time(msg) // Convert tickMsg back to time.Time
		// fmt.Println("Received tick at:", tickTime.Format("15:04:05.000"))
		T.Character = state.SingleGameState.GetContent()
		return T, tick

	default:
		fmt.Println("inside the default of msg.type ", msg)
	}

	return T, nil
}

func (T TestMultiplayer) View() string {
	fmt.Println("inside the view of player", T.Character)
	return T.Character
}
