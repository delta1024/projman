package main

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)


type spawnError struct {
	err error
}

func (e spawnError) Error() string {
	return e.err.Error()
}

func openSubShell(path string) tea.Cmd {
	cmd := exec.Command("bash")
	cmd.Dir = path
	cl := exec.Command("clear")
	cl.Dir = path
	return tea.Sequence(
		tea.ExecProcess(cl, nil),
		tea.ExecProcess(cmd, func(err error) tea.Msg {
			return spawnError{err: err}
		}), tea.ClearScreen)
}
