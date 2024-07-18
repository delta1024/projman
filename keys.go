package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Quit key.Binding
	Add  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Add}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Add},
	}
}

func defaultKeys() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("q", "quit"),
		),
		Add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add"),
		),
	}
}
