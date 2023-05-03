/* Styling for the welcome page only */

package styling

import "github.com/charmbracelet/lipgloss"

var WPHeader = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var Logo = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#13EDFF")). // 50
	PaddingLeft(4)
