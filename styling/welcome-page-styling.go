package styling

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")). //12
	Background(lipgloss.Color("15"))

var PirateStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("50")).
	PaddingTop(2).
	PaddingBottom(2)

var HelpBarStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var BorderStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	PaddingTop(1).
	Background(lipgloss.Color("#0c233b"))
