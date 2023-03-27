/* Styling for elements that appear in multiple pages */

package styling

import "github.com/charmbracelet/lipgloss"

var WaterMarkStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("25")).
	Bold(true)

var NavBarStyle = lipgloss.NewStyle().
	Faint(false)

var HelpBarStyle = lipgloss.NewStyle().
	PaddingLeft(2)
