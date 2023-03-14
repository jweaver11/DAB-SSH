package styling

import "github.com/charmbracelet/lipgloss"

var TitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12")).
	Background(lipgloss.Color("3"))

var FinalStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	PaddingTop(1)
