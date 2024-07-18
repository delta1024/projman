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
	items []string
	fp   dirs.Model
	dataLoaded bool
	err  error
}

func newModel() model {
	return model{
		mode: SelectMode,
		keys: defaultKeys(),
		help: help.New(),
		items: make([]string, 0),
		dataLoaded: false,
		list: lists.New([]string{ }),
	}
}

func (m model) Init() tea.Cmd {
	return loadData
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, saveData(m.items)
		case key.Matches(msg, m.keys.Add):
			m.mode = AddMode
			m.fp = dirs.New()
			return m, m.fp.Init()
		}
	case loadDataErrMsg:
	panic(msg.err)
	case loadedDataMsg:
		m.items = msg.data
		m.list = lists.New(m.items)
		return m, nil
	case savedDataFinishedMsg:
		return m, tea.Quit
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
			m.mode = SelectMode
			if !m.dataLoaded {

			m.items = append(m.items, m.fp.Selected)
			} else {
				m.dataLoaded = true
			m.items = make([]string, 1)
				m.items[0] = m.fp.Selected
		}
			m.list = lists.New(m.items)
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
