/* The help model for the Project Page
Declares the keymap struct and all the keys used for
just the project page */

package models

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type PPkeyMap struct {
	Enter    key.Binding
	Navigate key.Binding
	Quit     key.Binding
	Tab      key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (k PPkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Navigate, k.Enter, k.Tab, k.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (k PPkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Navigate, k.Enter, k.Tab, k.Quit},
	}
}

// Sets keys as our object using our keyMap struct from above
var PPkeys = PPkeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select project"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Navigate: key.NewBinding(
		key.WithKeys("↑↓/ws"),
		key.WithHelp("↑↓/ws", "navigate"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next page"),
	),
}
