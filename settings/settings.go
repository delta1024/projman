package settings

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
)

type Settings struct {
	DefaultDirectory string
	ClearScreen      bool
	Shell            string
}

func Defaults() Settings {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return Settings{
		DefaultDirectory: homeDir,
		ClearScreen:      true,
		Shell:            "bash",
	}
}

type SettingsLoadedMsg struct {
	Settings Settings
}
type NoSettingsDefinedMsg struct {}
type SettingsLoadErrMsg struct {
	 Err error
}
func LoadSettings() tea.Msg {
	dataFile, err := xdg.DataFile("projman/settings.txt")
	if err != nil {
		return SettingsLoadErrMsg{Err: err}
	}
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		return NoSettingsDefinedMsg{}
	}
	var file *os.File
	file, err = os.Open(dataFile)
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return SettingsLoadErrMsg{Err: err}
	}
	scanner := bufio.NewScanner(file)
	var s Settings
	scanner.Scan()
	s.DefaultDirectory = scanner.Text()
	scanner.Scan()
	if strings.Compare("x", scanner.Text()) == 0 {
		s.ClearScreen = true
	} else {
		s.ClearScreen = false
	}
	scanner.Scan()
	s.Shell = scanner.Text()
	return SettingsLoadedMsg{Settings: s}
}
type SettingsSaveErrMsg struct {Err error}
type SettingsSavedMsg struct {}
func SaveSettings(s Settings) tea.Cmd {
	return func() tea.Msg {
		dataFile, err := xdg.DataFile("projman/settings.txt")

		if err != nil {
			return SettingsSaveErrMsg{Err: err}
		}
		var file *os.File
		file, err = os.Create(dataFile)
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
		if err != nil {
			return SettingsSaveErrMsg{Err: err}
		}
		file.WriteString(s.DefaultDirectory + "\n")
		if s.ClearScreen {
			file.WriteString("x\n")
		} else {
			file.WriteString(" \n")
		}
		file.WriteString(s.Shell + "\n")
		file.Sync()
		return SettingsSavedMsg{}
	}
}
