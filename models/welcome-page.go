package models

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Our title page as a struct outlining the elements of our title page
type TitlePage struct {
	title  string     // The title
	navBar string     // The navigation bar below the title
	pirate string     // The pirate picture
	help   help.Model // The help bar at the bottom of the page
}

// Creates our title page and returns it to be used later
func CreateTitlePage() TitlePage {
	return TitlePage{
		title:  "cheese",
		navBar: "chedda",
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

	s += t.title + "\n\n"

	s += t.navBar + "\n\n"

	s += t.pirate + "\n\n"

	return s
}
