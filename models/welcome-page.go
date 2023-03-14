package models

import (
	"DAB-SSH/styling"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title  string     // The title
	navBar []string   // The navigation bar below the title
	pirate string     // The pirate picture
	help   help.Model // The help bar at the bottom of the page
}

// Creates our title page and returns it to be used later
func CreateTitlePage() TitlePage {
	title := "Digital Art Brokers"

	navBar := []string{"DAB", "Projects"}
	return TitlePage{
		title:  title,
		navBar: navBar,
		pirate: "mozzarella",
		help:   help.Model{},
	}
}

// Initializes our struct as a bubble tea model
func (t TitlePage) Init() tea.Cmd {
	return nil
}

// Updates our model everytime a key event happens, mainly window resizes and key presses
func (t TitlePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	return t, cmd
}

// Renders our model formatted to be viewed, then returns as a string
func (t TitlePage) View() string {
	var s string

	s += "  " + styling.TitleStyle.Render(t.title) + "\n\n"

	for i := range t.navBar {
		if i == 0 {
			s += styling.NavBarStyle.Foreground(lipgloss.Color("12")).Render(t.navBar[i])
		} else {
			s += styling.NavBarStyle.UnsetForeground().Render(t.navBar[i])
		}
	}

	s += t.pirate + "\n\n"

	return s
}
