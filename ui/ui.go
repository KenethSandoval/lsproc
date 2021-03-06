package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

var (
	color    = termenv.EnvColorProfile().Color
	keyword  = termenv.Style{}.Foreground(color("204")).Background(color("235")).Styled
	help     = termenv.Style{}.Foreground(color("241")).Styled
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func NewProgram() *tea.Program {
	procs := []list.Item{
		item{title: "Angular", desc: "4200 port"},
		item{title: "PHP", desc: "8000 port"},
	}
	m := model{list: list.New(procs, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Processes TCP"
	return tea.NewProgram(m, tea.WithAltScreen())
}
