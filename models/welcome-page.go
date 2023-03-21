/* This file controls the welcome page model, which is the page that
users first see when they connect to the SSH server. The model contains the
title of Digital Art Brokers, a navigation bar between the models, and a help bar
at the bottom of the page*/

// Tasks - Add DAB logo to end of nav bar
// get smaller DAB logo for title page

package models

import (
	"DAB-SSH/styling"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title                   string     // The title
	navBar                  []string   // The navigation bar below the title
	help                    help.Model // The help bar at the bottom of the page
	keys                    WPkeyMap   // Key map for our help model
	termWidth, termHeight   int        // Size of the terminal
	modelWidth, modelHeight int        // Size of the model
	minHeight               int        // Minimum size without model breaking
}

// Creates our title page gives it values
func CreateTitlePage() TitlePage {

	// Sets the title
	title := "Digital Art Brokers"

	// Sets the navbar values
	navBar := []string{"DAB", "Projects"}

	// Returns our created model
	return TitlePage{
		title:       title,
		navBar:      navBar,
		help:        help.New(),
		keys:        keys,
		modelWidth:  40,
		modelHeight: 30,
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

		// Sets the minimum height so model won't break
		t.minHeight = 22

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

	// Return our model and command
	return t, cmd
}

// Renders our model formatted to be viewed, then returns as a string
func (t TitlePage) View() string {

	// Final string to be rendered through our border at the end
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if t.termWidth <= t.modelWidth {
		width = t.modelWidth
	} else {
		width = t.termWidth
	}

	// Logic for setting terminal height to not break model
	if t.termHeight <= t.minHeight {
		height = t.minHeight
	} else {
		height = t.termHeight
	}

	// Pirate string for easy return later
	var pirate string

	// Sets the pirate ship to big or small one based on terminal size
	if t.termHeight < t.modelHeight { // 30 works for me
		pirate = lilPirate
	} else {
		pirate = DABLogo
	}

	// Adds the help bar at the bottom
	fullHelpView := t.help.View(t.keys)

	// Helps with sizing while testing
	fmt.Printf("Terminal height: %d\n", t.termHeight)
	fmt.Printf("Terminal width: %d\n", t.termWidth)
	fmt.Printf("Model height: %d\n", t.modelHeight)
	fmt.Printf("Model width: %d\n", t.modelWidth)

	// RENDERING OUR MODEL
	// Adds the header
	s += styling.WPHeaderStyle.Render(t.title) + "\n\n"

	// Adds the navbar and colors the selected page
	for i := range t.navBar {
		if i == 0 {
			s += styling.NavBarStyle.Faint(false).Render("• "+t.navBar[i]) + "	"
		} else {
			s += styling.NavBarStyle.UnsetForeground().Faint(true).Render("  "+t.navBar[i]) + "		\n"
		}
	}

	// Adds the pirate picture
	s += styling.PirateStyle.Render(pirate) + "\n\n"

	// Counts empty lines to put help model at bottom of terminal
	helpHeight := t.termHeight - strings.Count(s, "\n") - 3
	if helpHeight < 0 {
		helpHeight = 0
	}

	// Add empty lines if there are any to bottom of terminal
	s += strings.Repeat("\n", helpHeight)

	// Render help bar in correct styling
	s += styling.HelpBarStyle.Render(fullHelpView)

	// Returns model with final styling
	return styling.BorderStyle.Width(width).Height(height).Render(s)
}

// DAB logo using braille art thx to Braille ASCII Art generator
var DABLogo string = `
⠀⠀⠀⠀⠀⠀⠀⢀⢀⣠⠤⡄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⣀⣀⣔⣨⠉⠂⠀⠂⠁⠀⠱⣒⣠⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⡀⠀⠀
⠀⠀⢺⠱⣅⢊⢸⡀⢠⠀⢐⠀⡇⢀⡇⢕⢈⠢⡑⣈⠢⡑⣈⠢⡑⡨⠎⡇⠀⠀
⠀⠀⣹⢐⢼⠈⠠⢱⠐⠅⢸⠀⢨⠀⡇⢁⠁⡁⠡⠀⠅⡈⠄⠡⠈⡇⠅⡇⠀⠀
⠀⠀⢼⠐⢼⠀⢸⡁⠉⠣⠼⣀⡘⠒⠉⠉⠉⠉⠉⠉⠉⠈⠩⡆⠂⡥⠡⡇⠀⠀
⠀⠀⢺⠨⢺⠀⢸⠄⠐⡖⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⢺⠀⢘⠆⠀⡇⠌⡇⠀⠀
⠀⠀⣹⠨⢺⠈⢸⠀⠄⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢨⡃⠀⡇⠅⡇⠀⠀
⠀⠀⢼⢈⢺⠀⢸⠂⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠰⡅⠀⡇⠅⡇⠀⠀
⠀⠀⣺⢐⢹⠀⢸⠂⠁⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢘⠆⠁⡥⠑⡇⠀⠀
⠀⠀⢼⢐⢹⢀⢸⠂⢀⡇⠀⠀⠀⣄⣄⣤⣠⡀⠀⠀⢸⠀⢨⡃⠀⡇⠅⡇⠀⠀
⠀⠀⣹⠠⣹⠀⢸⠠⠀⡇⠀⠀⠐⣿⣿⣿⣿⡅⠀⠀⢸⠀⠰⡅⠂⡕⠡⡇⠀⠀
⠀⠀⢺⠨⣸⠀⢸⠂⠀⡇⠀⠀⠈⡿⡿⡿⡿⠆⠀⠀⢸⠀⢘⠆⠀⡇⠅⡇⠀⠀
⠀⠀⢽⢐⢼⠀⢸⠂⠁⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢨⡃⠄⡣⠡⡇⠀⠀
⠀⠀⣺⠐⢼⠐⢘⡂⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠰⡅⠀⡇⠅⡇⠀⠀
⠀⠀⢼⠨⢺⠀⢸⠠⠐⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢘⠆⡀⢇⠅⡇⠀⠀
⠀⠀⣹⢈⢺⢀⢸⠂⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢨⡃⠀⡇⠌⡇⠀⠀
⠀⠀⢺⢐⢹⠀⢸⠀⠁⡉⠌⠊⠈⠊⠈⠊⠘⠈⠡⠉⢊⣠⢐⠅⠄⡣⠡⡇⠀⠀
⠀⠀⢽⢐⢹⠀⡘⢑⠒⢂⠒⠒⠒⡑⠒⡑⠒⢪⢲⠉⡣⠀⠕⢳⡀⢇⠅⡇⠀⠀
⠀⠀⢺⠠⣹⢤⢤⢤⡢⣄⣆⢥⣰⣠⡢⡤⣬⡂⢸⠀⡊⠀⠄⢸⣤⡣⡈⡇⠀⠀
⠀⠀⣹⣮⣅⣆⣥⣢⣬⣰⣠⣅⣆⣔⣤⣱⡸⠀⠘⠀⠂⠀⠃⢸⣢⣸⣢⡇⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠒⠥⣠⡄⡠⢀⡀⡠⠎⠒⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠈⠈⠈⠀⠀⠀⠀⠀⠀⠀⠀`

// Small Pirate image (formatting messed up on github)
var lilPirate string = `

				#
				###
				#####
				###
				#
				#
				#
		#################
		  #############
			#########`
