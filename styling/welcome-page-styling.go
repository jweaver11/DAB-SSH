/* Styling for the welcome page only */

package styling

import "github.com/charmbracelet/lipgloss"

var WPHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var PirateStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("50")).
	PaddingLeft(4)

var BorderStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	PaddingTop(1)
