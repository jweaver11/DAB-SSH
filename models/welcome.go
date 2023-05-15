/* Controls the welcome page model, which is the page that
users first see when they connect to the SSH server. */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Our title page as a struct outlining the elements of our title page
type WelcomePage struct {
	title                            string           // The title
	waterMark                        string           // Watermark in top right corner of page
	help                             help.Model       // The help bar at the bottom of the page
	keys                             helpers.WPkeyMap // Key map for our help model
	modelWidth                       int              // Size of the model
	bigModelHeight, smallModelHeight int
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

/*
  ______ .______       _______     ___   .___________. _______     .___  ___.   ______    _______   _______  __
 /      ||   _  \     |   ____|   /   \  |           ||   ____|    |   \/   |  /  __  \  |       \ |   ____||  |
|  ,----'|  |_)  |    |  |__     /  ^  \ `---|  |----`|  |__       |  \  /  | |  |  |  | |  .--.  ||  |__   |  |
|  |     |      /     |   __|   /  /_\  \    |  |     |   __|      |  |\/|  | |  |  |  | |  |  |  ||   __|  |  |
|  `----.|  |\  \----.|  |____ /  _____  \   |  |     |  |____     |  |  |  | |  `--'  | |  '--'  ||  |____ |  `----.
 \______|| _| `._____||_______/__/     \__\  |__|     |_______|    |__|  |__|  \______/  |_______/ |_______||_______|

*/

// Creates and gives our model values
func CreateWelcomePage() WelcomePage {

	// Sets the title
	title := "Digital Art Brokers Official SSH Server"

	// Sets the watermark
	WM := " DAB "

	// Returns our created model
	return WelcomePage{
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
func (w WelcomePage) Init() tea.Cmd {
	// Returns no command
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
func (w WelcomePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		w.help.Width = msg.Width - styling.HelpBar.GetPaddingLeft()

		// Sets terminal width and height
		TerminalWidth = msg.Width
		TerminalHeight = msg.Height

	// Handles all keyboard presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// When q pressed, quit
		case "q", "esc", "ctrl+c":
			return w, tea.Quit

		// If key pressed and not one above, move to next page
		default:
			return CreateProjectPage(), tea.ClearScreen
		}
	}

	// Return our model and command
	return w, cmd
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
func (w WelcomePage) View() string {

	// Final string to be rendered through our border at the end
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if TerminalWidth <= w.modelWidth {
		width = w.modelWidth
	} else {
		width = TerminalWidth
	}

	// Logic for setting terminal height to not break model
	if TerminalHeight <= w.smallModelHeight {
		height = w.smallModelHeight
	} else {
		height = TerminalHeight
	}

	// Pirate string for easy return later
	var logo string

	// Sets the pirate ship to big or small one based on terminal size
	if TerminalHeight < w.bigModelHeight {
		logo = lilDABLogo
	} else {
		logo = BigDABLogo
	}

	// Adds the help bar at the bottom
	fullHelpView := w.help.View(w.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Adds the header
	s += styling.WPHeader.Render(w.title)

	// Adds watermark with padding to fit top right of page
	WMPadding := width - strings.Count(s, "")
	s += strings.Repeat(" ", WMPadding)
	s += styling.WaterMark.Render(w.waterMark) + "\n\n"

	// Adds the pirate picture
	s += styling.Logo.Render(logo) + "\n\n"

	// Puts help model at bottom of terminal with correct styling
	helpHeight := TerminalHeight - strings.Count(s, "\n") - 3
	if helpHeight < 0 {
		helpHeight = 0
	}
	s += strings.Repeat("\n", helpHeight)
	s += styling.HelpBar.Render(fullHelpView)

	// Returns model with final styling
	return styling.Border.Width(width).Height(height).Render(s)
}
