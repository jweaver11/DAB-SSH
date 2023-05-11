/* Styling for elements that appear in multiple pages */

package styling

import "github.com/charmbracelet/lipgloss"

var WaterMark = lipgloss.NewStyle().
	Background(lipgloss.Color("25")).
	Bold(true)

var NavBar = lipgloss.NewStyle().
	Faint(false)

var HelpBar = lipgloss.NewStyle().
	PaddingLeft(2)

var Border = lipgloss.NewStyle().
	PaddingLeft(2).
	PaddingTop(1)

var LightBlue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#13EDFF"))

var Blue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var White = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff"))
