package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/delta1024/projman/lists"
)

type model struct {
	keys keyMap
	help help.Model
	list lists.Model
	err error
}

func newModel() model {
	return model{
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
	case writeFileDone:
		return m, tea.Quit
	case writeFileError:
		m.err = msg.msg
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if m.list.Choice != "" {
		return m, writePathToFile(m.list.Choice)
	}
	return m, cmd
}
func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	return "\n" + m.list.View()
}
