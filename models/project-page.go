package models

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectPage struct {
	waterMark               string     // Watermark in top right corner of page
	navBar                  []string   // Nav bar below the title
	cursor                  int        // Used to track our cursor
	projects                []string   // An array of strings our our projects
	help                    help.Model // The help bar at the bottom of the page
	keys                    WPkeyMap   // Key map for our help model
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model
	minWidth, minHeight     int        // Minimum size without model breaking
}

func CreateProjectPage() ProjectPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "OtherThing", "AnotherThing", "MoreThing"}

	// Sets the cursor to 0
	cursor := 0

	projects := []string{"Project 1", "Project 2", "Project 3", "Project 4"}

	return ProjectPage{
		waterMark:   WM,
		navBar:      NB,
		cursor:      cursor,
		projects:    projects,
		help:        help.New(),
		keys:        PPkeys, // Sets our keymap to the project page keys
		modelWidth:  32,     // Change to actual model width
		modelHeight: 29,     // Change to actual model height
		minWidth:    32,
		minHeight:   14, // Might not work for this
	}
}

// Initializes our struct as a bubble tea model
func (p ProjectPage) Init() tea.Cmd {
	// Returns no command
	return nil
}

func (p ProjectPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		p.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		p.termWidth = msg.Width
		p.termHeight = msg.Height

	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Quits program
		case "q":
			return p, tea.Quit

		// Switches between full help view
		case "?":
			p.help.ShowAll = !p.help.ShowAll

		// Move cursor up
		case "up", "w":
			if p.cursor > 0 {
				p.cursor--
			}

		// Move cursor down
		case "down", "s":
			if p.cursor < len(p.projects)-1 {
				p.cursor++
			}
		}

	}

	return p, cmd
}

func (p ProjectPage) View() string {

	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if p.termWidth <= p.modelWidth {
		width = p.modelWidth
	} else {
		width = p.termWidth
	}

	// Logic for setting terminal height to not break model
	if p.termHeight <= p.minHeight {
		height = p.minHeight
	} else {
		height = p.termHeight
	}

	// RENDERING OUR MODEL *********************

	// Addds the watermark
	s += styling.WaterMarkStyle.Render(p.waterMark) + "\n\n"

	// Adds the navbar and colors the selected page
	for i := range p.navBar {

		if i == 0 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("12")).Render(p.navBar[i]) + "		"
		} else {
			s += styling.NavBarStyle.UnsetForeground().UnsetFaint().Render(p.navBar[i]) + "		"
		}

	}

	s += "\n\n"

	for i := range p.projects {

		cursor := "  "
		styling.SelectedProjectStyle.UnsetForeground()

		if p.cursor == i {
			cursor = "• "
			styling.SelectedProjectStyle.Foreground(lipgloss.Color("#7D56F4"))
		}

		s += styling.SelectedProjectStyle.Render(cursor+p.projects[i]) + "\n\n"
	}

	// Returns model with final styling
	return styling.BorderStyle.Width(width).Height(height).Render(s)
}
