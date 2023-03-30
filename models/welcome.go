/* Controls the welcome page model, which is the page that
users first see when they connect to the SSH server. The model contains the
title of Digital Art Brokers, a navigation bar between the models, and a help bar
at the bottom of the page*/

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title                            string           // The title
	waterMark                        string           // Watermark in top right corner of page
	help                             help.Model       // The help bar at the bottom of the page
	keys                             helpers.WPkeyMap // Key map for our help model
	termWidth, termHeight            int              // Size of the terminal
	modelWidth                       int              // Size of the model
	bigModelHeight, smallModelHeight int
}

// Creates our title page gives it values
func CreateTitlePage() TitlePage {

	// Sets the title
	title := "Digital Art Brokers Official SSH Server"

	// Sets the watermark
	WM := " DAB "

	// Returns our created model
	return TitlePage{
		title:            title,
		waterMark:        WM,
		help:             help.New(),     // Creates a new help model
		keys:             helpers.WPkeys, // Sets our keymap to the welcome page keymap
		modelWidth:       52,
		bigModelHeight:   29,
		smallModelHeight: 14,
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
		case "q", "esc", "ctrl+c":
			return t, tea.Quit

		// If key pressed and not one above, move to next page
		default:
			return CreateProjectPage(), cmd
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
	if t.termHeight <= t.smallModelHeight {
		height = t.smallModelHeight
	} else {
		height = t.termHeight
	}

	// Pirate string for easy return later
	var logo string

	// Sets the pirate ship to big or small one based on terminal size
	if t.termHeight < t.bigModelHeight {
		logo = lilDABLogo
	} else {
		logo = BigDABLogo
	}

	// Adds the help bar at the bottom
	fullHelpView := t.help.View(t.keys)

	// RENDERING OUR MODEL
	// Adds the header
	s += styling.WPHeaderStyle.Render(t.title)

	// Padding for the watermark to fit in corner of page
	WMPadding := width - strings.Count(s, "")

	// Adds padding for watermark
	s += strings.Repeat(" ", WMPadding-2)

	// Addds the watermark
	s += styling.WaterMarkStyle.Render(t.waterMark) + "\n\n"

	// Adds the pirate picture
	s += styling.PirateStyle.Render(logo) + "\n\n"

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
var BigDABLogo string = `
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
var lilDABLogo string = `
⠀⠀⢀⣀⠄⠔⠐⠠⣀⡀⣀⢀⢀⡀⡀⠀
⠀⢇⠆⢣⠐⡀⠃⡇⠒⡐⢘⠐⢰⢑⠀
⠀⡃⡇⢊⠡⠢⠅⠅⠍⠌⠌⡑⠨⡢⠀
⠀⡣⡑⡐⡀⠀⠀⠀⠀⠀⢀⢊⠨⡢⠀
⠀⡣⡑⡐⡀⠀⢀⣀⡀⠀⢀⠢⢘⠔⠀
⠀⡣⡑⡐⡀⠀⢹⣿⡧⠀⠐⡨⢐⠕⠀
⠀⡕⡅⠢⠀⠀⠀⠀⠀⠀⢀⠢⢘⢌⠀
⠀⢎⠢⡡⠁⠀⠀⠀⠀⠀⠀⡌⢔⠢⠀
⠀⢇⡃⠢⠨⠡⠩⠨⢡⠩⡐⢔⠰⡑⠀
⠀⣇⣚⣒⣓⣒⣓⣒⠆⠀⠃⢨⣪⡬⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠑⠒⠠⠊⠀⠀⠀`
