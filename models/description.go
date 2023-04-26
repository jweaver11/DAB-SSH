package models

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

const useHighPerformanceRenderer = false

type DescriptionPage struct {
	projectName             string     // Watermark in top left corner of page
	summary                 string     // Short summary of project at top of page
	description             string     // Actual description of the project
	help                    help.Model // Help bar at bottom of page
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model (not including help model)
}

func CreateDescriptionPage(projectAddress int, projectName string, summary string) DescriptionPage {

	// Sets the description passed through
	description := Descriptions[projectAddress]

	// Return our created model
	return DescriptionPage{
		projectName: projectName,
		summary:     summary,
		description: description,
	}
}

// Initializes our struct as a bubble tea model
func (d DescriptionPage) Init() tea.Cmd {
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
func (d DescriptionPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Back to project page
		case "esc":
			return CreateProjectPage(), cmd

		// Quit the program
		case "q":
			return d, tea.Quit
		}
	}

	return d, cmd
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
func (d DescriptionPage) View() string {

	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if d.termWidth <= d.modelWidth {
		width = d.modelWidth
	} else {
		width = d.termWidth
	}

	// Logic for setting terminal height to not break model
	if d.termHeight <= d.modelHeight {
		height = d.modelHeight
	} else {
		height = d.termHeight
	}

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// Addds the watermark
	s += d.projectName + "\n\n"

	s += d.summary

	s += "\n\n\n"

	s += d.description

	return styling.BorderStyle.Width(width).Height(height).Render(s)
}

var Descriptions = []string{
	"deez",
	"coconuts",
	"rule",
	"yo"}
