package terminal

import (
	"fmt"
	"strings"

	"github.com/Gabrielfernandes7/crabe/internal/agent"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	agent   *agent.Agent
	history []string
	input   string
	err     error
}

func initialModel(a *agent.Agent) model {
	return model{
		agent:   a,
		history: []string{},
		input:   "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.input != "" {
				userIn := m.input
				m.history = append(m.history, "> "+userIn)
				m.input = ""

				// Executa o agente de forma síncrona por agora (melhorar com tea.Cmd depois)
				resp, err := m.agent.Process(userIn)
				if err != nil {
					m.err = err
					m.history = append(m.history, "Erro: "+err.Error())
				} else {
					m.history = append(m.history, "Agente: "+resp)
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

	// Mostrar apenas as últimas 10 mensagens
	start := 0
	if len(m.history) > 10 {
		start = len(m.history) - 10
	}

	for i := start; i < len(m.history); i++ {
		s.WriteString(m.history[i] + "\n")
	}

	s.WriteString("\n> " + m.input + "_")

	return s.String()
}

func Run() {
	// Por enquanto usando valores padrão
	a := agent.NewAgent("http://localhost:11434", "qwen2.5:3b") // Usando 3b para ser mais rápido nos testes
	p := tea.NewProgram(initialModel(a))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao iniciar terminal: %v", err)
	}
}
