package settings

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Done   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Select, k.Done,
	}
}
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Done: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "go back"),
		),
	}
}
