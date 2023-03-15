package models

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit}, // first row
		{k.Help, k.Quit}, // second row
	}
}

// Sets keys as our object using our keyMap struct from above
var keys = keyMap{
	Help: key.NewBinding( // Show full help view
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding( // Quit program
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q", "quit"),
	),
}
