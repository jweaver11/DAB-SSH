package models

import tea "github.com/charmbracelet/bubbletea"

type DescriptionPage struct {
	waterMark               string   // Watermark in top right corner of page
	navBar                  []string // Nav bar below the title
	description             string   // Actual description of the project
	termWidth, termHeight   int      // Size of the terminal
	modelWidth, modelHeight int      // Size of the model (not including help model)
}

func CreateDescriptionPage(projectName int) DescriptionPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "About"}

	// Sets the description passed through
	description := Descriptions[projectName]

	return DescriptionPage{
		waterMark:   WM,
		navBar:      NB,
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
	return d.description
}

var Descriptions = [4]string{
	"deez",
	"coconuts",
	"rule",
	"yo"}
