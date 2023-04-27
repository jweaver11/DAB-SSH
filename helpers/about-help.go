/* The help model for the About Page
Declares the keymap struct and all the keys used for
just the project page */

package helpers

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type APkeyMap struct {
	Tab      key.Binding
	Enter    key.Binding
	Help     key.Binding
	Quit     key.Binding
	Navigate key.Binding
	Copy     key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (a APkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{a.Tab, a.Enter, a.Help, a.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (a APkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{a.Tab, a.Enter},     // First collumn
		{a.Help, a.Quit},     // Second collumn
		{a.Navigate, a.Copy}, // Third collumn
	}
}

// Sets keys as our object using our keyMap struct from above
var APkeys = PPkeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
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
		key.WithHelp("tab", "page"),
	),
	Copy: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "copy"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
}
