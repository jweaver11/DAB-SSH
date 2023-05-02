/*  About page that provides a description about Digital
Art Brokers and some of the work they do. */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct {
	waterMark               string   // Watermark in top left corner of page
	navBar                  []string // Nav bar below the title
	content                 string   // The text to describe DAB
	ready                   bool
	viewport                viewport.Model   // Viewport for scrolling
	help                    help.Model       // The help bar at the bottom of the page
	keys                    helpers.PPkeyMap // Key map for our help model
	termWidth, termHeight   int              // Size of the terminal
	modelWidth, modelHeight int              // Size of the model (not including help model)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const useHighPerformanceRenderer = false

func (a AboutPage) headerView() string {
	title := styling.WaterMarkStyle.Render(a.waterMark)
	line := strings.Repeat("─", max(0, a.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

// Adds the page scrolling part at the bottom
func (a AboutPage) footerView() string {
	info := fmt.Sprintf("%3.f%%", a.viewport.ScrollPercent()*100)
	line := strings.Repeat("─", max(0, a.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
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

	// Sets the watermark
	WM := " DAB "

	// Sets the navbar values
	NB := []string{"Projects", "About"}

	viewport := viewport.New(66, 28)
	viewport.SetContent(Content)

	// Sets the help model and styling
	help := help.New()
	help.Styles.ShortKey = styling.APHelpBarStyle
	help.Styles.FullKey = styling.APHelpBarStyle

	// Returns our created model
	return AboutPage{
		waterMark:   WM,
		navBar:      NB,
		viewport:    viewport,
		help:        help,
		keys:        helpers.APkeys, // Sets our keymap to the about page keys
		termHeight:  28,             // Init terminal height to not break model 40
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

	//var cmd tea.Cmd

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
		a.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		a.termWidth = msg.Width
		a.termHeight = msg.Height

		// Model height - helpbar
		headerHeight := lipgloss.Height(a.headerView())
		footerHeight := lipgloss.Height(a.footerView())
		//verticalMarginHeight := a.modelHeight - 4
		verticalMarginHeight := headerHeight + footerHeight

		if !a.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			a.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			a.viewport.YPosition = headerHeight
			a.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			a.viewport.SetContent(Content)
			a.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			a.viewport.YPosition = headerHeight + 1

			// Render the viewport one line below the header.
			a.viewport.YPosition = 7 + 1
		} else {
			a.viewport.Width = msg.Width
			a.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(a.viewport))
		}

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
	//s += styling.WaterMarkStyle.Render(a.waterMark) + "\n\n"

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

	s += a.headerView() + a.viewport.View() + a.footerView()

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
