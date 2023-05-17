/* The help model for the About Page */

package helpers

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type APkeyMap struct {
	Tab    key.Binding
	Quit   key.Binding
	Scroll key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (a APkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{a.Tab, a.Scroll, a.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (a APkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{a.Tab}, // First collumn
		{a.Scroll},
		{a.Quit},
	}
}

// Sets keys as our object using our keyMap struct from above
var APkeys = APkeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Scroll: key.NewBinding(
		key.WithKeys("↑↓/ws"),
		key.WithHelp("↑↓/ws", "scroll"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "page"),
	),
}
