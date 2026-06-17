package terminal

import (
	"fmt"
	"strings"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/agent"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type model struct {
	agent   *agent.Agent
	history []string
	input   string
	err     error
	loading bool
	spinner int
}

func initialModel(a *agent.Agent) model {
	return model{
		agent:   a,
		history: []string{},
		input:   "",
		loading: false,
		spinner: 0,
	}
}

type errMsg struct{ err error }
type agentRespMsg struct{ resp string }

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.loading {
			m.spinner = (m.spinner + 1) % 4
		}
		return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	case agentRespMsg:
		m.loading = false
		m.history = append(m.history, "Agente: "+msg.resp)
		return m, nil
	case errMsg:
		m.loading = false
		m.err = msg.err
		m.history = append(m.history, "Erro: "+msg.err.Error())
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.input != "" {
				userIn := m.input
				m.history = append(m.history, "> "+userIn)
				m.loading = true
				m.input = ""

				// Executa o agente de forma assíncrona
				return m, func() tea.Msg {
					resp, err := m.agent.Process(userIn)
					if err != nil {
						return errMsg{err}
					}
					return agentRespMsg{resp}
				}
			}
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	s.WriteString(ui.RenderHeader())

	start := 0
	if len(m.history) > 10 {
		start = len(m.history) - 10
	}

	for i := start; i < len(m.history); i++ {
		s.WriteString(m.history[i] + "\n")
	}

	if m.loading {
		frames := []struct {
			icon  string
			color string
	}{
			{"⠋", "\033[33m"},
			{"⠙", "\033[33m"},
			{"⠹", "\033[38;5;208m"},
			{"⠸", "\033[38;5;208m"},
			{"⠼", "\033[31m"},
			{"⠴", "\033[31m"},
		}

		frame := frames[m.spinner%len(frames)]

		s.WriteString(
			"\n" +
				frame.color +
				frame.icon +
				" Thinking..." +
				"\033[0m\n",
		)
	}

	s.WriteString("\n> " + m.input + "_")

	return s.String()
}

func Run() {
	// Por enquanto usando valores padrão
	a := agent.NewAgent("http://localhost:11434", "llama3.1:8b") // Usando 8b para ser mais rápido nos testes
	p := tea.NewProgram(initialModel(a))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao iniciar terminal: %v", err)
	}
}
