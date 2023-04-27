/* Short about page that provides a description about Digital
Art Brokers and some of the work they do. */

package models

import (
	"DAB-SSH/styling"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	waterMark               string     // Watermark in top left corner of page
	navBar                  []string   // Nav bar below the title
	content                 string     // The text to describe DAB
	help                    help.Model // The help bar at the bottom of the page
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model (not including help model)
}

func CreateAboutPage() AboutPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "About"}

	content := "deez bolls my guy"

	// Returns our created model
	return AboutPage{
		waterMark: WM,
		navBar:    NB,
		content:   content,
	}
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

		case "tab":
			return CreateProjectPage(), tea.ClearScreen
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
	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if a.termWidth <= a.modelWidth {
		width = a.modelWidth
	} else {
		width = a.termWidth
	}

	// Logic for setting terminal height to not break model
	if a.termHeight <= a.modelHeight {
		height = a.modelHeight
	} else {
		height = a.termHeight
	}

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// temp
	fmt.Println(width + height)

	// Adds the watermark
	s += styling.WaterMarkStyle.Render(a.waterMark) + "\n\n"

	// Adds the navbar and highlights the selected page
	for i := range a.navBar {
		if i == 1 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("#7D56F4")).Render(a.navBar[i]) + "		"
		} else {
			s += styling.NavBarStyle.UnsetForeground().UnsetFaint().Render(a.navBar[i]) + "		"
		}
	}

	return s
}
