package models

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type DescriptionPage struct {
	waterMark               string     // Watermark in top right corner of page
	summary                 string     // Short summary of project at top of page
	description             string     // Actual description of the project
	help                    help.Model // Help bar at bottom of page
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model (not including help model)
}

func CreateDescriptionPage(projectName int, summary string) DescriptionPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the description passed through
	description := Descriptions[projectName]

	return DescriptionPage{
		waterMark:   WM,
		summary:     summary,
		description: description,
	}
}

func (d DescriptionPage) Init() tea.Cmd {
	return nil
}

func (d DescriptionPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	return d, cmd
}

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
	s += styling.WaterMarkStyle.Render(d.waterMark) + "\n\n"

	s += d.summary

	s += "\n\n\n"

	s += d.description

	return styling.BorderStyle.Width(width).Height(height).Render(s)
}

var Descriptions = [4]string{
	"deez",
	"coconuts",
	"rule",
	"yo"}
