package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type model struct {
	textInput textinput.Model
	spinner   spinner.Model
	loading   bool
	sent      bool
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter a message..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		textInput: ti,
		spinner:   sp,
		loading:   false,
		sent:      false,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.textInput.Value() != "" {
				m.loading = true
				m.sent = false
				// Here you would typically send the message to the server.
				// For this example, we'll just simulate a network request.
				return m, tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
					return "message sent"
				})
			}
		}

	case string:
		m.loading = false
		m.sent = true
		m.textInput.Reset()
		return m, nil

	case error:
		m.err = msg
		return m, nil
	}

	var cmds []tea.Cmd
	if m.loading {
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nError: %v\n\n", m.err)
	}

	var b strings.Builder

	b.WriteString(docStyle.Render(m.titleView()))
	b.WriteString("\n")

	if m.loading {
		b.WriteString(fmt.Sprintf("%s Sending message...", m.spinner.View()))
	} else {
		b.WriteString(m.textInput.View())
		if m.sent {
			b.WriteString(helpStyle.Render("\nMessage sent!"))
		}
	}

	b.WriteString(helpStyle.Render("\n\nPress Enter to send, Esc to quit."))

	return b.String()
}

func (m model) titleView() string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1).
		Render("Go Modern Scaffold CLI")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
