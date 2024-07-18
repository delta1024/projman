package dirs

import (
	"os"
	"path"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

const height = 8
type Model struct {
	filepicker filepicker.Model
	Selected   string
}

func New() Model  {
	fp := filepicker.New()
	fp.FileAllowed = false
	fp.DirAllowed = true
	fp.Height = height
	fp.AutoHeight = false
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fp.CurrentDirectory = path.Join(dir, "code")
	return Model{
		filepicker: fp,
	}
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)
	if selected, file := m.filepicker.DidSelectFile(msg); selected {
		m.Selected = file
		return m, nil
	}
	return m, cmd
}
func (m Model) View() string {
	return "\n" + m.filepicker.View()
}
