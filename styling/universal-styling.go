package styling

import "github.com/charmbracelet/lipgloss"

var WaterMarkStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("25")) //#14C126

var NavBarStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	Faint(false)

var HelpBarStyle = lipgloss.NewStyle().
	PaddingLeft(2)
