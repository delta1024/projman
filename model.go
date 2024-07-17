package main

import (
	"os"
	"path"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type progMode int

const (
	Init progMode = iota - 1
	DisplayFile
	RePickFile
)

type model struct {
	keys         keyMap
	help         help.Model
	filepicker   filepicker.Model
	selectedFile string
	mode         progMode
}

func newFp() filepicker.Model {
	fp := filepicker.New()
	hDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fp.CurrentDirectory = path.Join(hDir, "code")
	fp.AutoHeight = false
	fp.Height = 5
	return fp
}
func newModel() model {
	return model{
		keys:       defaultKeys(),
		help:       help.New(),
		filepicker: newFp(),
		mode:       RePickFile,
	}
}

func (m model) Init() tea.Cmd {
	return m.filepicker.Init()
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Restart):
			if m.mode == DisplayFile {
				m.mode = RePickFile
				m.filepicker = newFp()
				m.keys.FilePicker = true
				return m, m.filepicker.Init()
			}
		}
	}
	switch m.mode {
	case RePickFile:
		var cmd tea.Cmd
		m.filepicker, cmd = m.filepicker.Update(msg)
		if selected, file := m.filepicker.DidSelectFile(msg); selected {
			m.selectedFile = file
			m.keys.FilePicker = false
			m.mode = DisplayFile
		}
		return m, cmd
	}
	return m, nil
}
func (m model) View() string {
	var msg string
	if m.mode == DisplayFile {
		msg = m.selectedFile + "\n"
	} else if m.mode == Init {
		return ""
	} else {
		msg = m.filepicker.View() + "\n"
	}
	helpView := m.help.View(m.keys)

	return "\n" + msg + helpView
}
