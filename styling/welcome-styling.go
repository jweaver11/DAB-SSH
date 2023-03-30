/* Styling for the welcome page only */

package styling

import "github.com/charmbracelet/lipgloss"

var WPHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var LogoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("50")).
	PaddingLeft(4)
