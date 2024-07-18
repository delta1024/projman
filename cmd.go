package main

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func openSubShell(path string) tea.Cmd {
	cmd := exec.Command("bash")
	cmd.Dir = path
	cl := exec.Command("clear")
	cl.Dir = path
	return tea.Sequence(
		tea.ExecProcess(cl, nil),
		tea.ExecProcess(cmd, nil), tea.ClearScreen)
}
