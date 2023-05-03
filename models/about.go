/*  About page that provides a description about Digital
Art Brokers and some of the work they do. */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
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
	NB := []string{"Projects", "• About"}
	WM := " DAB "

	// Create Viewport
	viewport := viewport.New(TerminalWidth, TerminalHeight-10)
	viewport.SetContent(Content)

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
		case "tab":
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
	s += styling.APViewport.Render(a.viewport.View()) + "\n"

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

// The about DAB body
var Content string = `Mission:
Digital Art Brokers (DAB) is a blockchain and NFT gamification company dedicated to creating
unique gaming experiences while giving back through charitable partnerships.
Buccaneers of the Blockchain (BotB) is DAB's first project, created in partnership with the national
nonprofit organization DAV (Disabled American Veterans). BotB is a crypto-based game that allows
players to learn the basics of blockchain gaming and engage in a transparent support system for
disabled veterans – all while competing to earn real money.
BotB is the first piece of a larger project aiming to illustrate DAB’s gamified model of service that
re-imagines the donor/recipient relationship. DAB has plans to offer a host of experiences to offer
another space where community-building and Web3 experience(s) can expand.
DAB is committed to transparency, charitable impact, and providing unique gaming experiences.
DAB’s ultimate objective is to utilize Web3, NFTs, and blockchain-based technologies to provide
community value and initiate participation. Achievement of this will be satisfied by coding
service into its tech, re-imagining “traditional” models of charitable giving, and
re-conceptualizing the donor/recipient paradigm.
BotB's partnership with DAV is an example of DAB's mission to do great while doing great. Players
can enjoy a fun and competitive gaming experience while contributing to a worthy cause.

Partnership with DAV:
DAV empowers veterans to lead high-quality lives with respect and dignity. It is dedicated to a single
purpose: keeping our promise to America's veterans. DAV ensures that veterans and their families
can access the full range of benefits available to them, fights for the interests of America's injured
heroes on Capitol Hill, provides employment resources to veterans and their families, and educates
the public about the great sacrifices and needs of veterans transitioning back to civilian life.
DAB has committed 25% (from mint) & 40% (of in-game purchases) of all BotB revenues to DAV, and
any action performed within the game will result in further contribution. DAV CEO Mark Burgess is
thrilled to partner with a great start-up company like DAB.

DAB Head Developer Josh Ferguson has stated that bringing the DAV's purpose forward and into
the blockchain space would demonstrate the far-reaching potential of the blockchain to support
DAV in a wildly new and meaningful way – adding that he and the team anticipate great things to
come.`
