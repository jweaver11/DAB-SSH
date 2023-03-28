/* Controls the project page model, which is the page that
shows users the projects DAB is working on. The model contains the watermark
and navbar, a cursor to point at the projects, the projects and short description
of them, and a help model at bottom of page */

package models

import (
	"DAB-SSH/styling"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectPage struct {
	waterMark                     string     // Watermark in top right corner of page
	navBar                        []string   // Nav bar below the title
	cursor                        int        // Used to track our cursor
	projects, descriptions, links []string   // An array of strings of our projects and descriptions
	help                          help.Model // The help bar at the bottom of the page
	keys                          PPkeyMap   // Key map for our help model
	termWidth, termHeight         int        // Size of the terminal
	modelWidth, modelHeight       int        // Size of the model (not including help model)
}

func CreateProjectPage() ProjectPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "About"}

	// Sets the cursor to 0
	cursor := 0

	// Sets our projects
	projects := []string{"Buccaneers of the Blockchain",
		"SSH App",
		"Discord",
		"AI Development"}

	// Sets the short descriptions
	descriptions := []string{"NFT buccaneers that look pretty darn cool",
		"The SSH app you are currently using right now",
		"The official DAB discord server",
		"Artfifical Intelligence of the future"}

	links := []string{"https://buccaneers.io",
		"https://github.com/jweaver11/DAB-SSH",
		"https://discord.com/invite/dabinc",
		"AI.org bozo"}

	// Returns our newly created model
	return ProjectPage{
		waterMark:    WM,
		navBar:       NB,
		cursor:       cursor,
		projects:     projects,
		descriptions: descriptions,
		links:        links,
		help:         help.New(),
		keys:         PPkeys, // Sets our keymap to the project page keys
		modelWidth:   66,     // Change to actual model width
		modelHeight:  24,     // Change to actual model height
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

	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Quits program
		case "q":
			return p, tea.Quit

		// Switches between full help view
		case "?":
			if p.termHeight-p.modelHeight >= 4 {
				p.help.ShowAll = !p.help.ShowAll
			}

		// Copy link to clipboard
		case "c":

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

		// Move to next page
		case "tab":
			return p, cmd
		}

	}

	// Return the model and cmd
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
	if p.termHeight <= p.modelHeight {
		height = p.modelHeight
	} else {
		height = p.termHeight
	}

	// Adds the help bar at the bottom
	fullHelpView := p.help.View(p.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// Addds the watermark
	s += styling.WaterMarkStyle.Render(p.waterMark) + "\n\n"

	// Adds the navbar and highlights the selected page
	for i := range p.navBar {
		if i == 0 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("12")).Render(p.navBar[i]) + "		"
		} else {
			s += styling.NavBarStyle.UnsetForeground().UnsetFaint().Render(p.navBar[i]) + "		"
		}
	}

	s += "\n\n\n"

	// Adds our listed projects and short descriptions
	for i := range p.projects {

		// Reset formatting
		styling.SelectedProjectStyle.UnsetFaint().UnsetForeground()

		// Sets cursor to blank if not selected
		cursor := "  "

		// Sets our cursor to dot if selected
		if p.cursor == i {
			cursor = "• "
			styling.SelectedProjectStyle.Foreground(lipgloss.Color("#7D56F4"))
		}

		// Adds the project and description
		s += styling.SelectedProjectStyle.Render(cursor+p.projects[i]) + "\n"
		s += styling.SelectedProjectStyle.UnsetForeground().Faint(true).Render("   "+p.descriptions[i]) + "\n"
		s += styling.SelectedProjectStyle.UnsetFaint().Foreground(lipgloss.Color("25")).Render("    "+p.links[i]) + "\n\n\n"
	}

	// Counts empty lines to put help model at bottom of terminal
	emptyLines := p.termHeight - strings.Count(s, "\n") - 3
	if emptyLines < 0 {
		emptyLines = 0
	}

	// Add empty lines if there are any to bottom of terminal
	s += strings.Repeat("\n", emptyLines)

	// Render help bar in correct styling
	s += styling.HelpBarStyle.Render(fullHelpView)

	// Returns model with final styling
	return styling.BorderStyle.Width(width).Height(height).Render(s)
}
