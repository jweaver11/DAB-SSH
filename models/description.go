package models

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

//const useHighPerformanceRenderer = false

type DescriptionPage struct {
	projectName             string     // Watermark in top left corner of page
	summary                 string     // Short summary of project at top of page
	description             string     // Actual description of the project
	help                    help.Model // Help bar at bottom of page
	modelWidth, modelHeight int        // Size of the model (not including help model)
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
func CreateDescriptionPage(projectAddress int, projectName string, summary string) DescriptionPage {

	// Sets the description passed through
	description := Descriptions[projectAddress]

	// Return our created model
	return DescriptionPage{
		projectName: projectName,
		summary:     summary,
		description: description,
	}
}

// Initializes our struct as a bubble tea model
func (d DescriptionPage) Init() tea.Cmd {
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
func (d DescriptionPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var cmd tea.Cmd

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:
		TerminalWidth = msg.Width
		TerminalHeight = msg.Height

	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Back to project page
		case "esc":
			return CreateProjectPage(), cmd

		// Quit the program
		case "q":
			return d, tea.Quit
		}
	}

	return d, cmd
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
func (d DescriptionPage) View() string {

	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if TerminalWidth <= d.modelWidth {
		width = d.modelWidth
	} else {
		width = TerminalWidth
	}

	// Logic for setting terminal height to not break model
	if TerminalHeight <= d.modelHeight {
		height = d.modelHeight
	} else {
		height = TerminalHeight
	}

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Addds the watermark
	s += d.projectName + "\n\n"

	// Add the summary
	s += d.summary

	// Adds spacing
	s += "\n\n\n"

	// Adds the description
	s += d.description

	return styling.BorderStyle.Width(width).Height(height).Render(s)
}

var Descriptions = []string{`Game Overview:
BotB is a battle-card strategy game that combines choice-driven narrative features and programmed
charitable impact. All in-game transactions are visible throughout the two-week experience,
ensuring complete transparency for the players. The game's smart contracts guarantee a higher
reward, unique to the NFT space: programmatically-assured charitable impact. Winning gamers will
have the option to share their "spoils" with DAV for further impact.
Charitable giving on the blockchain offers the promise of complete transparency. DAB has
guaranteed that promise by disallowing human intervention in BotB's contract-cemented code. The
game's strict parameters and automated controls circumvent "trust-based" agreements, which are
vulnerable to human error, greed, or changeability. BotB self-monitors by way of a system built on
total transparency.
BotB's core objective is to do great while doing great. It challenges the existing NFT gaming
community, the crypto-curious, and even the most skeptical crypto critics to get involved. Players
will be empowered, entertained, and educated while supporting a truly deserving cause.
	
	Gameplay:
In BotB, "Cards" or “player game pieces” are NFTs, and their value appreciates the more the game is
played. Buccaneers in BotB have three categories of stats: physical, mental, and spiritual, and belong
to one of three classes (human, cyborg, and robot), each starting out with the same stats for the 1st
iteration. Battles are initiated by any player at any time. The game also features a defense mode that
grants players an increased chance of winning if they are attacked.
Interactive elements and battle outcomes are determined by the game's programming and informed
by player choices. All actions carry a small cost, and the game continuously distributes these
amounts to its main three pools: DAV, the treasure chest (prize), and DAB (BotB's team).`,
	"coconuts",
	"rule",
	"yo"}
