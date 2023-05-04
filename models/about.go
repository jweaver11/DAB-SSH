/*  About page that provides a description about Digital
Art Brokers and some of the work they do. */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	waterMark string           // Watermark in top left corner of page
	navBar    []string         // Nav bar below the title
	content   string           // The text to describe DAB
	viewport  viewport.Model   // Viewport for scrolling
	help      help.Model       // The help bar at the bottom of the page
	keys      helpers.PPkeyMap // Key map for our help model
	minWidth  int              // Minimum Width so model won't break
}

// The about page DAB body
var AboutContent, _ = os.ReadFile("content/aboutpage/about.md")

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

	// Sets the navbar and watermark
	NB := []string{"  Projects", "• About"}
	WM := " DAB "

	// Create Viewport and sets content
	viewport := viewport.New(TerminalWidth, TerminalHeight-10)
	viewport.SetContent(string(AboutContent))

	// Sets the help model and styling
	help := help.New()
	help.Styles.ShortKey = styling.APHelpBar
	help.Styles.FullKey = styling.APHelpBar

	// Returns our created model
	return AboutPage{
		waterMark: WM,
		navBar:    NB,
		viewport:  viewport,
		help:      help,
		keys:      helpers.APkeys, // Sets our keymap to the about page keys
		minWidth:  45,             // Change to actual model width
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
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		a.help.Width = msg.Width - styling.HelpBar.GetPaddingLeft()

		// Sets terminal width and height
		TerminalWidth = msg.Width
		TerminalHeight = msg.Height

		// Viewport Size
		a.viewport.Width = msg.Width - styling.Border.GetPaddingLeft()
		a.viewport.Height = msg.Height - 10

	// All key presses
	case tea.KeyMsg:

		// Converts the key press into a string
		switch msg.String() {

		// Quits the program
		case "q":
			return a, tea.Quit

		// Switches back to project page
		case "tab", "esc":
			return CreateProjectPage(), tea.ClearScreen

		// Switches between full help view
		case "?":
			a.help.ShowAll = !a.help.ShowAll

		}

	}

	// Handle keyboard and mouse events in the viewport
	a.viewport, cmd = a.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
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

	// Logic for setting terminal size to not break model
	if TerminalWidth <= a.minWidth {
		width = a.minWidth
	} else {
		width = TerminalWidth
	}
	height = TerminalHeight

	// Adds the help bar at the bottom
	fullHelpView := a.help.View(a.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Adds the navbar and highlights the selected page
	for i := range a.navBar {
		if i == 1 {
			s += styling.NavBar.Foreground(lipgloss.Color("12")).Render(a.navBar[i])
		} else {
			s += styling.NavBar.UnsetForeground().UnsetFaint().Render(a.navBar[i]) + "            "
		}
	}

	// Adds watermark with padding to fit top right of page
	WMPadding := width - strings.Count(s, "")
	s += strings.Repeat(" ", WMPadding)
	s += styling.WaterMark.Render(a.waterMark) + "\n\n"
	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft()))
	s += "\n\n"

	// Adds viewport
	s += styling.APViewport.Render(a.viewport.View()) + "\n\n"
	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft()))

	// Counts empty lines to put help model at bottom of terminal
	emptyLines := TerminalHeight - strings.Count(s, "\n") - 3
	if emptyLines < 0 {
		emptyLines = 0
	}

	// Add empty lines if there are any to bottom of terminal
	s += strings.Repeat("\n", emptyLines)

	// Render help bar in correct styling
	s += styling.HelpBar.Render(fullHelpView)

	// Returns model with final styling
	return styling.Border.Width(width).Height(height).Render(s)
}
