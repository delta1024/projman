package lists

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
const listHeight = 14

const DefaultTitle =  "Select project"
var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
type Model struct {
	List list.Model
	Choice string
	Keys KeyMap
}
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	m.List.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
	switch  {
		case key.Matches(msg, m.Keys.Select):
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, nil
		
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
func (m Model) View() string {
	if m.Choice != "" {
		return ""
	}
	return "\n" + m.List.View()
}
func NewNoExt(paths []string, keys []key.Binding) Model {
	items := make([]list.Item, 0)
	for _, path := range paths {
		items = append(items, Item(path))
	}
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = DefaultTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return append(keys, DefaultKeyMap().Select)	
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
			return append(keys, DefaultKeyMap().Select)
	}

	 return Model{List: l, Keys: DefaultKeyMap()}
	
}
func New(paths []string, keys []key.Binding , extra key.Binding) Model  {
	items := make([]list.Item, 0)
	for _, path := range paths {
		items = append(items, Item(path))
	}
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = DefaultTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return append(keys, DefaultKeyMap().Select)	
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
			return append(keys, extra, DefaultKeyMap().Select)
	}

	 return Model{List: l, Keys: DefaultKeyMap()}
}
