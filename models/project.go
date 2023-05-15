/* Controls the project page model, which is the page that
shows users the projects DAB is working on.  */

package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model for the page
type ProjectPage struct {
	navBar                   []string         // Nav bar below the title
	waterMark                string           // Watermark in top right corner of page
	cursor                   int              // Used to track our cursor
	projects, summary, links []string         // An array of strings of our projects and descriptions
	help                     help.Model       // The help bar at the bottom of the page
	keys                     helpers.PPkeyMap // Key map for our help model
	minWidth                 int              // Minimum width so model won't break
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
func CreateProjectPage() ProjectPage {

	// Sets the navbar and watermark
	NB := []string{"• Projects", "  About"}
	WM := " DAB "

	// Sets the cursor to 0
	cursor := 0

	// Sets our projects
	projects := []string{"Buccaneers of the Blockchain",
		"SSH App",
		"Discord",
		"AI Development"}

	// Sets the short descriptions
	summary := []string{"NFT buccaneers that look pretty darn cool",
		"The SSH app you are currently using right now",
		"The official DAB discord server",
		"Artfifical Intelligence of the future"}

	// Sets the links
	links := []string{"https://buccaneers.io",
		"https://github.com/jweaver11/DAB-SSH",
		"https://discord.com/invite/dabinc",
		"AI.org bozo"}

	// Sets the help model and styling
	help := help.New()
	help.Styles.ShortKey = styling.PPHelpBar
	help.Styles.FullKey = styling.PPHelpBar

	// Returns our newly created model
	return ProjectPage{
		waterMark: WM,
		navBar:    NB,
		cursor:    cursor,
		projects:  projects,
		summary:   summary,
		links:     links,
		help:      help,
		keys:      helpers.PPkeys, // Sets our keymap to the project page keys
		minWidth:  55,
	}
}

// Initializes our struct as a bubble tea model
func (p ProjectPage) Init() tea.Cmd {
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
func (p ProjectPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		p.help.Width = msg.Width - styling.HelpBar.GetPaddingLeft()

		// Sets terminal width and height
		TerminalWidth = msg.Width
		TerminalHeight = msg.Height

	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Quits program
		case "q", "ctrl+c":
			return p, tea.Quit

		case "esc":
			return CreateWelcomePage(), tea.ClearScreen

		// Switches between full help view
		case "?":
			p.help.ShowAll = !p.help.ShowAll

		// Copy link to clipboard
		case "c":
			link := p.links[p.cursor]
			err := clipboard.WriteAll(link)
			if err != nil {
				panic(err)
			}
			str := p.links[p.cursor]
			p.links[p.cursor] = "copied!"

			go func() {
				time.AfterFunc(time.Second, func() {
					p.links[p.cursor] = str // change the string after one second
				})
			}()

		// Move cursor up
		case "up", "w", "left":
			if p.cursor > 0 {
				p.cursor--
			}

		// Move cursor down
		case "down", "s", "right":
			if p.cursor < len(p.projects)-1 {
				p.cursor++
			}

		// Move to next page
		case "tab":
			return CreateAboutPage(), tea.ClearScreen

		// Returns description page of selected project
		case "enter":
			return CreateDescriptionPage(p.cursor, p.projects[p.cursor], p.summary[p.cursor]), tea.ClearScreen
		}

	}

	// Return the model and cmd
	return p, cmd
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
func (p ProjectPage) View() string {

	// Our string to build our model on
	var s string

	// Logic for setting terminal size to not break model
	var width, height int
	if TerminalWidth <= p.minWidth {
		width = p.minWidth
	} else {
		width = TerminalWidth
	}

	// Adds the help bar at the bottom
	fullHelpView := p.help.View(p.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Adds the navbar and highlights the selected page
	for i := range p.navBar {
		if i == 0 {
			s += styling.NavBar.Foreground(lipgloss.Color("12")).Render(p.navBar[i]) + "            "
		} else {
			s += styling.NavBar.UnsetForeground().UnsetFaint().Render(p.navBar[i])
		}
	}
	// Adds watermark with padding to fit top right of page
	WMPadding := width - strings.Count(s, "")
	s += strings.Repeat(" ", WMPadding)
	s += styling.WaterMark.Render(p.waterMark) + "\n\n"
	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft()))
	s += "\n\n"

	// Adds our listed projects and short descriptions
	for i := range p.projects {

		// Reset formatting
		styling.SelectedProject.UnsetFaint().UnsetForeground()

		// Sets to a line if select of blank if not
		cursor := "  "
		if p.cursor == i {
			cursor = "┃ "
			styling.SelectedProject.Foreground(lipgloss.Color("12"))
		} else {
			styling.SelectedProject.Faint(true).Foreground(lipgloss.Color("12"))
		}

		// Adds the cursor, project name, summary, and links
		s += styling.SelectedProject.Render(cursor+p.projects[i]) + "\n"
		if cursor == "┃ " {
			s += styling.Blue.Render(cursor) + styling.SelectedProject.UnsetFaint().Foreground(lipgloss.Color("#ffffff")).Render("  "+p.summary[i]) + "\n"
			s += styling.Blue.Render(cursor) + styling.SelectedProject.UnsetFaint().Foreground(lipgloss.Color("25")).Render("   "+p.links[i]) + "\n\n\n"
		} else {
			s += styling.Blue.Render(cursor) + styling.SelectedProject.UnsetForeground().Foreground(lipgloss.Color("#ffffff")).Faint(true).Render("  "+p.summary[i]) + "\n"
			s += styling.Blue.Render(cursor) + styling.SelectedProject.Faint(true).Foreground(lipgloss.Color("25")).Render("   "+p.links[i]) + "\n\n\n"
		}

	}

	// Puts help model at bottom of terminal with correct styling
	emptyLines := TerminalHeight - strings.Count(s, "\n") - 5
	if emptyLines < 0 {
		emptyLines = 0
	}
	s += strings.Repeat("\n", emptyLines)
	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft())) + "\n\n"
	s += styling.HelpBar.Render(fullHelpView)

	// Returns model with final styling
	return styling.Border.Width(width).Height(height).Render(s)
}
