package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

	return s
}

type model struct {
	index       int
	width       int
	height      int
	styles      *Styles
	questions   []Question
	answerField textinput.Model
}

type Question struct {
	question string
	answer   string
}

func NewQuestion(question string) Question {
	return Question{question: question}
}

func New(questions []Question) *model {
	styles := DefaultStyles()
	answerField := textinput.New()
	answerField.Placeholder = "Your answer here"
	answerField.Width = 75
	answerField.Focus()
	return &model{questions: questions, answerField: answerField, styles: styles}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			current := &m.questions[m.index]
			current.answer = m.answerField.Value()
			cangofurthur := m.Next()

			if m.index < len(m.questions) {

				if !cangofurthur {
					return m, nil
				}

				m.answerField.SetValue("")
				return m, nil
			}
		}
	}

	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	if m.index == len(m.questions) {
		var result string
		for i, q := range m.questions {
			result += fmt.Sprintf("%d. %s\n   Answer: %s\n\n", i+1, q.question, q.answer)
		}
		result += "Press Ctrl+C to exit"

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			m.styles.InputField.Render(result),
		)
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.questions[m.index].question,
			m.styles.InputField.Render(m.answerField.View()),
		),
	)
}

func (m *model) Next() bool {
	if m.index < len(m.questions)-1 {
		m.index++
		return true
	} else {
		m.index++
		return false
	}
}

func main() {
	questions := []Question{
		NewQuestion("What is your name?"),
		NewQuestion("What is your github username"),
		NewQuestion("why you want to work with us?"),
	}

	m := New(questions)

	f, err := tea.LogToFile("debug.log", "debug")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
