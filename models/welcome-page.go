package models

// Make description comment at top of each file
// Use other method to get terminal size, this one too smol

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title                   string     // The title
	navBar                  []string   // The navigation bar below the title
	help                    help.Model // The help bar at the bottom of the page
	keys                    keyMap     // Key map for our help model
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model
}

// Creates our title page and returns it to be used later
func CreateTitlePage() TitlePage {

	// Sets the title
	title := "Digital Art Brokers"

	// Sets the navbar values
	navBar := []string{"DAB", "Projects"}

	// Returns our created model
	return TitlePage{
		title:  title,
		navBar: navBar,
		help:   help.New(),
		keys:   keys,
	}
}

// Initializes our struct as a bubble tea model
func (t TitlePage) Init() tea.Cmd {
	// Returns no command
	return nil
}

// Updates our model everytime a key event happens, mainly window resizes and key presses
func (t TitlePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		t.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		t.termWidth = msg.Width
		t.termHeight = msg.Height

	// Handles all keyboard presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// When q pressed, quit
		case "q":
			return t, tea.Quit

		// When ? pressed, switch between short help view and full help view
		case "?":
			t.help.ShowAll = !t.help.ShowAll
		}
	}

	return t, cmd
}

// Renders our model formatted to be viewed, then returns as a string
func (t TitlePage) View() string {

	// Final string to be rendered through our border at the end
	var s string

	// Adds the header
	s += "  " + styling.HeaderStyle.Render(t.title) + "\n\n"

	// Adds the navbar and colors the selected page
	for i := range t.navBar {
		if i == 0 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("12")).Render(t.navBar[i])
		} else {
			s += styling.NavBarStyle.UnsetForeground().Render(t.navBar[i])
		}
	}

	// Adds the pirate picture
	s += "\n\n" + pirate + "\n\n"

	// Adds the help bar at the bottom
	fullHelpView := t.help.View(t.keys)

	s += styling.HelpBarStyle.Render(fullHelpView) + "\n"

	// Returns our string model according to the terminal size with some padding
	return styling.BorderStyle.Width(t.termWidth).Height(t.termHeight).Render(s)
}

// Pirate image
var pirate string = `
				##
				#####
				########
				###########
				########
				#####
				##
				#
				#
				#
				#
				#
##################################
 ###############################
   ###########################
	 #######################
	   ###################
		 ###############`
