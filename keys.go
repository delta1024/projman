package main

import (
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Quit       key.Binding
	Restart    key.Binding
	Fpicker    filepicker.KeyMap
	FilePicker bool
}

func (k keyMap) ShortHelp() []key.Binding {
	if k.FilePicker {
		return []key.Binding{k.Quit, k.Fpicker.Up, k.Fpicker.Down, k.Fpicker.Open, k.Fpicker.Back}
	} else {
		return []key.Binding{k.Quit, k.Restart}
	}
}
func (k keyMap) FullHelp() [][]key.Binding {
	if k.FilePicker {
		return [][]key.Binding{
			{k.Quit, k.Fpicker.Select, k.Fpicker.Up, k.Fpicker.Down, k.Fpicker.Open, k.Fpicker.Back},
			{k.Fpicker.PageUp, k.Fpicker.PageDown, k.Fpicker.GoToTop, k.Fpicker.GoToLast},
		}

	} else {
		return [][]key.Binding{
			{k.Quit, k.Restart},
		}
	}
}

func defaultKeys() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("q", "quit"),
		),
		Restart: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "choose again"),
		),
		Fpicker:    filepicker.DefaultKeyMap(),
		FilePicker: true,
	}
}
