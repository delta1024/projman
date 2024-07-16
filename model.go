package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	keys keyMap
	help help.Model
}

func newModel() model {
	return model{
		keys: defaultKeys(),
		help: help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m model) View() string {
	msg := "\n\nPress q to exit\n\n"
	helpView := m.help.View(m.keys)

	return "\n" + msg + helpView
}
