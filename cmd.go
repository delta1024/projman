package main

import (
	"bufio"
	"errors"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
)

func openSubShell(path, shell string, clear bool) tea.Cmd {
	cmd := exec.Command(shell)
	cmd.Dir = path
	cl := exec.Command("clear")
	cl.Dir = path
	if clear {
	return tea.Sequence(
		tea.ExecProcess(cl, nil),
		tea.ExecProcess(cmd, nil), tea.ClearScreen)
	} else {
		return tea.ExecProcess(cmd, nil)
	}
}

type noSavedDataMsg struct{}
type loadedDataMsg struct {
	data []string
}
type loadDataErrMsg struct {
	err error
}

func (e loadDataErrMsg) Error() string {
	return e.err.Error()
}

func loadData() tea.Msg {
	dataFile, err := xdg.DataFile("projman/projects.txt")
	if err != nil {
		return loadDataErrMsg{err: err}
	}
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		return noSavedDataMsg{}
	}
	var file *os.File
	file, err = os.Open(dataFile)
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return loadDataErrMsg{err: err}
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return loadDataErrMsg{err: err}
	}
	return loadedDataMsg{data: lines}
}

type savedDataFinishedMsg struct{}
type savedDataErrMsg struct {
	err error
}

func (e savedDataErrMsg) Error() string {
	return e.err.Error()
}
func saveData(data []string) tea.Cmd {
	return func() tea.Msg {
		dataFile, err := xdg.DataFile("projman/projects.txt")

		if err != nil {
			return savedDataErrMsg{err: err}
		}
		var file *os.File
		file, err = os.Create(dataFile)
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
		if err != nil {
			return savedDataErrMsg{err: err}
		}
		if len(data) == 0 {
			return savedDataFinishedMsg{}
		}
		for _, str := range data {
			file.WriteString(str + "\n")
		}
		file.Sync()
		return savedDataFinishedMsg{}
	}
}
