/* Short about page that provides a description about Digital
Art Brokers and some of the work they do. */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	waterMark               string           // Watermark in top left corner of page
	navBar                  []string         // Nav bar below the title
	content                 string           // The text to describe DAB
	help                    help.Model       // The help bar at the bottom of the page
	keys                    helpers.PPkeyMap // Key map for our help model
	termWidth, termHeight   int              // Size of the terminal
	modelWidth, modelHeight int              // Size of the model (not including help model)
}

var Content string = ``

/*
  ______ .______       _______     ___   .___________. _______     .___  ___.   ______    _______   _______  __
 /      ||   _  \     |   ____|   /   \  |           ||   ____|    |   \/   |  /  __  \  |       \ |   ____||  |
|  ,----'|  |_)  |    |  |__     /  ^  \ `---|  |----`|  |__       |  \  /  | |  |  |  | |  .--.  ||  |__   |  |
|  |     |      /     |   __|   /  /_\  \    |  |     |   __|      |  |\/|  | |  |  |  | |  |  |  ||   __|  |  |
|  `----.|  |\  \----.|  |____ /  _____  \   |  |     |  |____     |  |  |  | |  `--'  | |  '--'  ||  |____ |  `----.
 \______|| _| `._____||_______/__/     \__\  |__|     |_______|    |__|  |__|  \______/  |_______/ |_______||_______|

*/

// Creates and gives our model values
func CreateAboutPage() AboutPage {

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "About"}

	// Sets our content from the bottom of page
	content := Content

	// Sets the help model and styling
	help := help.New()
	help.Styles.ShortKey = styling.APHelpBarStyle
	help.Styles.FullKey = styling.APHelpBarStyle

	// Returns our created model
	return AboutPage{
		waterMark:   WM,
		navBar:      NB,
		content:     content,
		help:        help,
		keys:        helpers.APkeys, // Sets our keymap to the about page keys
		termHeight:  28,             // Init terminal height to not break model
		modelWidth:  66,             // Change to actual model width
		modelHeight: 24,             // Change to actual model height
	}
}

// Initializes our struct as a bubble tea model
func (a AboutPage) Init() tea.Cmd {
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
func (a AboutPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		a.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		a.termWidth = msg.Width
		a.termHeight = msg.Height

	// All key presses
	case tea.KeyMsg:

		// Converts the key press into a string
		switch msg.String() {

		// Switches back to project page
		case "tab":
			return CreateProjectPage(), tea.ClearScreen

		// Switches between full help view
		case "?":
			if a.termHeight-a.modelHeight >= 4 {
				a.help.ShowAll = !a.help.ShowAll
			}
		}

	}

	return a, cmd
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
func (a AboutPage) View() string {
	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if a.termWidth <= a.modelWidth {
		width = a.modelWidth
	} else {
		width = a.termWidth
	}

	// Logic for setting terminal height to not break model
	if a.termHeight <= a.modelHeight {
		height = a.modelHeight
	} else {
		height = a.termHeight
	}

	// Adds the help bar at the bottom
	fullHelpView := a.help.View(a.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Adds the watermark
	s += styling.WaterMarkStyle.Render(a.waterMark) + "\n\n"

	// Adds the navbar and highlights the selected page
	for i := range a.navBar {
		if i == 1 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("#7D56F4")).Render(a.navBar[i]) + "            "
		} else {
			s += styling.NavBarStyle.UnsetForeground().UnsetFaint().Render(a.navBar[i]) + "            "
		}
	}

	// Spacing
	s += "\n\n\n"

	// Counts empty lines to put help model at bottom of terminal
	emptyLines := a.termHeight - strings.Count(s, "\n") - 3
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
