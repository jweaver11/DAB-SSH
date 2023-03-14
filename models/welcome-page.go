package models

import (
	"DAB-SSH/styling"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title                 string     // The title
	navBar                []string   // The navigation bar below the title
	help                  help.Model // The help bar at the bottom of the page
	keys                  keyMap     // Key map for our help model
	termWidth, termHeight int        //terminal width and height
}

// Creates our title page and returns it to be used later
func CreateTitlePage() TitlePage {

	// Sets the title
	title := "Digital Art Brokers"

	// Set terminal width and height
	TW, TH, _ := term.GetSize(0)

	// Sets the navbar values
	navBar := []string{"DAB", "Projects"}

	// Returns our created model
	return TitlePage{
		title:      title,
		navBar:     navBar,
		help:       help.New(),
		termWidth:  TW,
		termHeight: TH,
	}
}

// Initializes our struct as a bubble tea model
func (t TitlePage) Init() tea.Cmd {
	return nil
}

// Updates our model everytime a key event happens, mainly window resizes and key presses
func (t TitlePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		t.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		t.termWidth, t.termHeight, _ = term.GetSize(0)
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

	fullHelpView := t.help.View(t.keys)

	s += styling.HelpBarStyle.Render(fullHelpView)

	fmt.Println(pirate)

	return styling.BorderStyle.Render(s)
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
 ##                           ##
   ##                       ##
	 ##                   ##
	   ##               ##
		 ###############`
