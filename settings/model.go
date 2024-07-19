package settings

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/delta1024/projman/lists"

	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	view mode = iota - 1
	changeDefDir
	changeShell
)
type item string
func (i item) FilterView() string  {
	return ""
}
var allowedShells = []string{
	"bash",
	"zsh",
	"fish",
}

type Model struct {
	mode     mode
	Settings Settings
	options  []string
	pos      int
	keys     KeyMap
	help     help.Model
	fp       filepicker.Model
	sp lists.Model
	Done     bool
}


func (m Model) Init() tea.Cmd {
	return LoadSettings
}
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.mode == view {
			switch {
			case key.Matches(msg, m.keys.Up):
				if m.pos > 0 {
					m.pos -= 1
				}
				return m, nil
			case key.Matches(msg, m.keys.Down):
				if m.pos < len(m.options)-1 {
					m.pos += 1
				}
				return m, nil
			case key.Matches(msg, m.keys.Select):
			switch option(m.pos) {
				case defaultDir:
					m.fp = newFp()
					m.mode = changeDefDir
					return m, m.fp.Init()
				case clearScreen:
					m.Settings.ClearScreen = !m.Settings.ClearScreen
					return m, nil
				case shell:
					m.mode = changeShell
					m.sp = newList()
					return m, nil
				}
			case key.Matches(msg, m.keys.Done):
				m.Done = true
				return m, SaveSettings(m.Settings)
			}

		}
	}
	switch m.mode {
	case changeDefDir:
		var cmd tea.Cmd
		m.fp, cmd = m.fp.Update(msg)
		if selected, path := m.fp.DidSelectFile(msg); selected {
			m.Settings.DefaultDirectory = path
			m.mode = view
			return m, nil
		}
		return m, cmd
	case changeShell:
	var cmd tea.Cmd
		m.sp, cmd = m.sp.Update(msg)
		if m.sp.Choice != "" {
			m.Settings.Shell = m.sp.Choice
			m.mode = view
			return m, nil
		}
		return m, cmd
	}
	return m, nil
}

type option int

const (
	defaultDir  option = 0
	clearScreen option = 1
	shell       option = 2
)

func (m Model) View() string {
	switch m.mode {
	case changeDefDir:
		return m.fp.View()
	case changeShell:
		return m.sp.View()
	}
	s := "\n\nSettings:\n\n"
	for i := range m.options {
		pos := " "
		if i == m.pos {
			pos = ">"
		}

		switch option(i) {
		case defaultDir:
			s += fmt.Sprintf("%s default directory:  %s\n", pos, m.Settings.DefaultDirectory)
		case clearScreen:
			clrScreen := "[ ]"
			if m.Settings.ClearScreen {
				clrScreen = "[x]"
			}
			s += fmt.Sprintf("%s clear screen:       %s\n", pos, clrScreen)
		case shell:
			s += fmt.Sprintf("%s shell:              %s\n", pos, m.Settings.Shell)
		}
	}
	helpView := m.help.View(m.keys)
	return s + "\n" + helpView
}
func newFp() filepicker.Model {
	fp := filepicker.New()
	fp.DirAllowed = true
	fp.FileAllowed = false
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fp.CurrentDirectory = dir
	fp.Height = 7
	return fp
}
func newList() lists.Model {
	l := lists.New(allowedShells, nil)
	l.List.Title = "Choose Shell"
	return l
}
func New(settings Settings) Model {

	return Model{
		Settings: settings,
		options: []string{
			"default directory",
			"clear screen",
			"shell",
		},
		keys: DefaultKeyMap(),
		help: help.New(),
		mode: view,
		fp: newFp(),
		sp: newList(),
	}
}
