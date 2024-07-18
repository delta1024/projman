package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/delta1024/projman/dirs"
	"github.com/delta1024/projman/lists"
)

type mode int

const (
	SelectMode mode = iota - 1
	AddMode
)

type model struct {
	mode mode
	keys keyMap
	help help.Model
	list lists.Model
	fp   dirs.Model
	err  error
}

func newModel() model {
	return model{
		mode: SelectMode,
		keys: defaultKeys(),
		help: help.New(),
		list: lists.New([]string{
			"/home/jake/code/go/projman",
			"/home/jake/code/zig/projman",
		}),
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
		case key.Matches(msg, m.keys.Add):
			m.mode = AddMode
			m.fp = dirs.New()
			return m, m.fp.Init()
		}
	case spawnError:
		return m, nil
	}
	switch m.mode {
	case SelectMode:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		if m.list.Choice != "" {
			path := m.list.Choice
			m.list.Choice = ""
			return m, openSubShell(path)
		}
		return m, cmd
	case AddMode:
		var cmd tea.Cmd
		m.fp, cmd = m.fp.Update(msg)
		if m.fp.Selected != "" {
			items := m.list.List.Items()
			var strs []string
			for _, item := range items {
				strs = append(strs, string(item.(lists.Item)))
			}
			m.mode = SelectMode
			m.list = lists.New(append(strs, m.fp.Selected))
			return m, nil
		}
		return m, cmd
	}
	return m, nil
}
func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	switch m.mode {
	case SelectMode:
		return "\n" + m.list.View()
	case AddMode:
		return "\n" + m.fp.View()
	default:
		return ""
	}
}
