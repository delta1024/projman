package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
	}
}

var defaultKeyMap = keyMap{
	Help: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "show help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "exit app"),
	),
}
