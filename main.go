package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	index  int
	width  int
	height int
}

func New() *model {
	return &model{index: 0}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return "timer"
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case string:
		if msg == "timer" {
			m.index++
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			if m.index == 3 || m.index == 0 {
				return m, nil
			}
			m.index++
		case "left":
			if m.index == 1 || m.index == 0 {
				return m, nil
			}
			m.index--
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	var content string
	content += `
     ██╗  █████╗  ██╗ ███╗  ██╗ ███████╗ ██╗  ██╗
     ██║ ██╔══██╗ ██║ ████╗ ██║ ██╔════╝ ╚██╗██╔╝
     ██║ ███████║ ██║ ██╔██╗██║ █████╗    ╚███╔╝ 
██╗  ██║ ██╔══██║ ██║ ██║╚████║ ██╔══╝    ██╔██╗ 
╚█████╔╝ ██║  ██║ ██║ ██║ ╚███║ ███████╗ ██╔╝╚██╗
 ╚════╝  ╚═╝  ╚═╝ ╚═╝ ╚═╝  ╚══╝ ╚══════╝ ╚═╝  ╚═╝
`

	navbox := lipgloss.NewStyle().
		Width(m.width - 20).Align(lipgloss.Center).PaddingTop(2)

	navlink1 := `home`
	navlink2 := `projects`
	navlink3 := `experience`

	switch m.index {
	case 1:
		content = navlink1
	case 2:
		content = navlink2
	case 3:
		content = navlink3
	}

	navstyle := lipgloss.NewStyle().
		Width(20).
		Padding(0, 1).
		Align(lipgloss.Center)

	navstyleActive := navstyle.Copy().
		Width(20).
		Padding(0, 1).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))

	nav := lipgloss.JoinHorizontal(lipgloss.Center,
		lipgloss.NewStyle().Render(
			func() string {
				if m.index == 1 {
					return navstyleActive.Render(navlink1)
				}
				return navstyle.Render(navlink1)
			}()),
		lipgloss.NewStyle().Render(func() string {
			if m.index == 2 {
				return navstyleActive.Render(navlink2)
			}
			return navstyle.Render(navlink2)
		}()),
		lipgloss.NewStyle().Render(func() string {
			if m.index == 3 {
				return navstyleActive.Render(navlink3)
			}
			return navstyle.Render(navlink3)
		}()),
	)

	navRendered := navbox.Render(nav)

	box := lipgloss.NewStyle().
		Width(m.width).
		Height(func() int {
			if m.index >= 1 {
				return m.height - 20
			}
			return m.height
		}()).
		MarginTop(2).
		Align(lipgloss.Center, lipgloss.Center)

	contentRendered := lipgloss.NewStyle().
		Width(m.width-50).
		Height(m.height-20).
		Align(lipgloss.Center, lipgloss.Center).
		Padding(2, 10, 2, 10).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Render(content)

	boxRendered := box.Render(contentRendered)

	if m.index == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, boxRendered)
	} else {
		return lipgloss.JoinVertical(lipgloss.Center, navRendered, boxRendered)
	}
}

func main() {

	m := New()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
