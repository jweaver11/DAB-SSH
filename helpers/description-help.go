/* The help model for the About Page
Declares the keymap struct and all the keys used for
just the project page */

package helpers

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type DPkeyMap struct {
	Back   key.Binding
	Scroll key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (d DPkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{d.Scroll, d.Back}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (d DPkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.Scroll}, // First collumn
		{d.Back},   // Second collumn
	}
}

// Sets keys as our object using our keyMap struct from above
var DPkeys = DPkeyMap{
	Back: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Scroll: key.NewBinding(
		key.WithKeys("↑↓/ws"),
		key.WithHelp("↑↓/ws", "scroll"),
	),
}
