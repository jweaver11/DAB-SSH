package styling

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12")).
	Background(lipgloss.Color("3"))

var PirateStyle = lipgloss.NewStyle()

var HelpBarStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var BorderStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	PaddingTop(1).
	PaddingRight(2).
	PaddingBottom(1)
