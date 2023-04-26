package models

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type AboutPage struct {
	waterMark string     // Watermark in top right corner of page
	navBar    []string   // Nav bar below the title
	content   string     // The text to describe DAB
	help      help.Model // The help bar at the bottom of the page
}

func CreateAboutPage() AboutPage {

	return AboutPage{}
}

// Initializes our struct as a bubble tea model
func (a AboutPage) Init() tea.Cmd {
	return nil
}

/*
__    __  .______    _______       ___   .___________. _______
|  |  |  | |   _  \  |       \     /   \  |           ||   ____|
|  |  |  | |  |_)  | |  .--.  |   /  ^  \ `---|  |----`|  |__
|  |  |  | |   ___/  |  |  |  |  /  /_\  \    |  |     |   __|
|  `--'  | |  |      |  '--'  | /  _____  \   |  |     |  |____
 \______/  | _|      |_______/ /__/     \__\  |__|     |_______|
*/
// Updates our model everytime a key event happens, mainly window resizes and key presses
func (a AboutPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {
		}
	}

	return a, cmd
}

/*
____    ____  __   ___________    __    ____
\   \  /   / |  | |   ____\   \  /  \  /   /
 \   \/   /  |  | |  |__   \   \/    \/   /
  \      /   |  | |   __|   \            /
   \    /    |  | |  |____   \    /\    /
    \__/     |__| |_______|   \__/  \__/
*/
// Renders our model formatted to be viewed, then returns as a string
func (a AboutPage) View() string {
	return "deez bols bozo"
}
