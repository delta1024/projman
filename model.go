package main

import (
	"strings"

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
		keys: defaultKeyMap,
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
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Pick a project: \n\n"

	helpView := m.help.View(m.keys)
	height := 8 - strings.Count(s, "\n") - strings.Count(helpView, "\n")
	return "\n" + s + strings.Repeat("\n", height) + helpView
}
