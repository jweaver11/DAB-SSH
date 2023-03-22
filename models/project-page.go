package models

import (
	"DAB-SSH/styling"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectPage struct {
	title                   string        // The title
	waterMark               string        // Watermark in top right corner of page
	navBar                  []string      // Nav bar below the title
	cursor                  int           // Used to track our cursor
	spinner                 spinner.Model // Spinner used as our cursor
	help                    help.Model    // The help bar at the bottom of the page
	keys                    WPkeyMap      // Key map for our help model
	termWidth, termHeight   int           // Size of the terminal
	modelWidth, modelHeight int           // Size of the model
	minWidth, minHeight     int           // Minimum size without model breaking
}

func CreateProjectPage() ProjectPage {

	// Sets the title
	title := "Digital Art Brokers"

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "-", "Descriptions"}

	return ProjectPage{
		title:       title,
		waterMark:   WM,
		navBar:      NB,
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

		// When q pressed, quit
		case "q":
			return p, tea.Quit

		// When ? pressed, switch between short help view and full help view
		case "?":
			p.help.ShowAll = !p.help.ShowAll
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
	// Adds the header
	s += styling.WPHeaderStyle.Render(p.title)

	// Padding for the watermark to fit in corner of page
	WMPadding := width - strings.Count(s, "")
	// Adds padding for watermark
	s += strings.Repeat(" ", WMPadding-2)
	// Addds the watermark
	s += styling.WaterMarkStyle.Render(p.waterMark) + "\n\n"

	// Adds the navbar and colors the selected page
	for i := range p.navBar {
		if i <= 2 {
			if i == 0 {
				s += styling.NavBarStyle.Foreground(lipgloss.Color("12")).Render("• " + p.navBar[i])
			} else if i == 1 {
				s += lipgloss.NewStyle().UnsetForeground().Faint(true).Render(p.navBar[i])
			} else if i == 2 {
				s += styling.NavBarStyle.UnsetForeground().Render(p.navBar[i])
			}
		} else {
			s += styling.NavBarStyle.UnsetForeground().UnsetFaint().Render(p.navBar[i])
		}
	}

	s += "\n"

	s += "Projects go here"

	// Returns model with final styling
	return styling.BorderStyle.Width(width).Height(height).Render(s)
}
