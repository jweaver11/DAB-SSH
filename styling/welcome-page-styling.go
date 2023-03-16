package styling

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12")).
	Background(lipgloss.Color("3")).
	PaddingTop(1).
	PaddingBottom(1)

var PirateStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("55")).
	PaddingTop(2).
	PaddingBottom(2)

var HelpBarStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var BorderStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Background(lipgloss.Color("#0c233b"))
