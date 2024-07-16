package main

import (
	// "strings"

	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"

	// "github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/delta1024/projman/list"
)

type model struct {
	keys     keyMap
	help     help.Model
	items    list.Model
	projects []string
}

func newModel() model {
	m := model{
		keys:     defaultKeyMap,
		help:     help.New(),
		items:    list.New([]string{}),
		projects: []string{"ghost in the shell", "yu yu hakusho", "trigun"},
	}
	m.items = list.New(m.projects)
	return m
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
	items, cmd := m.items.Update(msg)
	list, ok := items.(list.Model)
	if !ok {
		return m, nil
	}
	m.items = list
	return m, cmd
}

func (m model) View() string {
	s := "Pick a project: \n\n"
	s += m.items.View()

	helpView := m.help.View(m.keys)
	height := strings.Count(s, "\n") - strings.Count(helpView, "\n")

	return "\n" + s + strings.Repeat("\n", height) + helpView
}
