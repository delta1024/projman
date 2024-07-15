package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type state struct {
	projects []string
	menu     []string
	cursor   int
	selected map[int]struct{}
}

func initState() state {
	return state{
		projects: []string{},
		menu: []string{
			"Add Project",
		},
		selected: make(map[int]struct{}),
	}
}

func (s state) Init() tea.Cmd {
	return nil
}

func (s state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}

		case "down", "j":
			if s.cursor < (len(s.projects)-1)+(len(s.menu)-1) {
				s.cursor++
			}
		case "enter", " ":
			_, ok := s.selected[s.cursor]
			if ok {
				delete(s.selected, s.cursor)
			} else {
				s.selected[s.cursor] = struct{}{}
			}
		}
	}
	return s, nil
}
func (s state) View() string {
	o := "Please select a project to work on:\n\n"

	for i, choice := range s.menu {
		cursor := " "
		if s.cursor == i {
			cursor = ">"
		}
		o += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	o += "\n---\n\n"
	pi := len(s.menu) - 1
	for i, choice := range s.projects {
		cursor := " "
		if s.cursor == i+pi {
			cursor = ">"
		}
		o += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	o += "\nPress q to quit.\n"
	return o
}
func main() {
	p := tea.NewProgram(initState())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
