package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)
func writePathToFile(path string) tea.Cmd {
	return func () tea.Msg {
		file, err := os.CreateTemp("", "ProjManCd")
		defer file.Close()
		if err != nil {
			return writeFileError{msg: err}
		}
		_, err = file.WriteString(path)
		if err != nil {
			return writeFileError{msg: err}
		}
		return writeFileDone{}
	}
}

type writeFileDone struct {}
type writeFileError struct{
	msg error
}
func (e writeFileError) Error() string  {
	return e.msg.Error()
}

