package terminal

import (
	"fmt"
	"strings"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/agent"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type messageRole string

const (
	roleUser  messageRole = "user"
	roleAgent messageRole = "agent"
	roleError messageRole = "error"
)

type chatMessage struct {
	role messageRole
	body string
}

type model struct {
	agent     *agent.Agent
	modelName string
	history   []chatMessage
	input     string
	err       error
	loading   bool
	spinner   int
}

func initialModel(a *agent.Agent, modelName string) model {
	return model{
		agent:     a,
		modelName: modelName,
		history:   []chatMessage{},
		input:     "",
		loading:   false,
		spinner:   0,
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
			m.spinner = (m.spinner + 1) % len(spinnerFrames)
		}
		return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	case agentRespMsg:
		m.loading = false
		m.history = append(m.history, chatMessage{role: roleAgent, body: msg.resp})
		return m, nil
	case errMsg:
		m.loading = false
		m.err = msg.err
		m.history = append(m.history, chatMessage{role: roleError, body: msg.err.Error()})
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.loading {
				return m, nil
			}
			if m.input != "" {
				if strings.HasPrefix(m.input, "/") {
					command := m.input
					m.history = append(m.history, chatMessage{role: roleUser, body: command})
					m.input = ""
					
					switch command {
					case "/help":
						m.history = append(m.history, chatMessage{role: roleAgent, body: "Comandos disponíveis:\n/help - Mostra ajuda\n/skills - Lista suas habilidades\n/init - Inicializa o projeto"})
					case "/skills":
						m.history = append(m.history, chatMessage{role: roleAgent, body: "Habilidades:\n- Leitura de arquivos\n- Escrita de arquivos\n- Listagem de diretórios"})
					default:
						m.history = append(m.history, chatMessage{role: roleError, body: "Comando não reconhecido: " + command})
					}
					return m, nil
				}
				userIn := m.input
				m.history = append(m.history, chatMessage{role: roleUser, body: userIn})
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
	s.WriteString(statusBarStyle.Render(" Local agent ", " "+m.modelName+" ", " /help /init /skills "))
	s.WriteString("\n\n")

	start := 0
	if len(m.history) > 8 {
		start = len(m.history) - 8
	}

	for i := start; i < len(m.history); i++ {
		s.WriteString(renderMessage(m.history[i]))
		s.WriteString("\n")
	}

	if m.loading {
		s.WriteString("\n")
		s.WriteString(thinkingStyle.Render(spinnerFrames[m.spinner%len(spinnerFrames)] + " pensando"))
		s.WriteString("\n")
	}

	if len(m.history) == 0 && !m.loading {
		s.WriteString(emptyStateStyle.Render("Pergunte algo ou digite / para comandos."))
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString(promptStyle.Render("❯"))
	s.WriteString(" ")
	s.WriteString(inputStyle.Render(m.input))
	s.WriteString(cursorStyle.Render(" "))

	return s.String()
}

var (
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1A1A1A")).
			Background(lipgloss.Color("#FFCC00")).
			Padding(0, 1)

	userLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D7FF"))

	agentLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC00"))

	errorLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF4D4D"))

	messageBodyStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E6E6E6")).
				PaddingLeft(2)

	errorBodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF9B9B")).
			PaddingLeft(2)

	thinkingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFCC00")).
			PaddingLeft(1)

	emptyStateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8A8A8A")).
			Italic(true).
			PaddingLeft(1)

	promptStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC00"))

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FFCC00"))
)

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴"}

func renderMessage(msg chatMessage) string {
	switch msg.role {
	case roleUser:
		return userLabelStyle.Render("Você") + "\n" + messageBodyStyle.Render(msg.body)
	case roleAgent:
		return agentLabelStyle.Render("Crabe") + "\n" + messageBodyStyle.Render(msg.body)
	case roleError:
		return errorLabelStyle.Render("Erro") + "\n" + errorBodyStyle.Render(msg.body)
	default:
		return messageBodyStyle.Render(msg.body)
	}
}

func Run(modelName string) {
	a := agent.NewAgent("http://localhost:11434", modelName)
	p := tea.NewProgram(initialModel(a, modelName))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao iniciar terminal: %v", err)
	}
}
