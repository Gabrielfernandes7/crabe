package agent

import (
	"context"
	"github.com/Gabrielfernandes7/crabe/internal/llm"
	"github.com/Gabrielfernandes7/crabe/internal/tools"
)

type Agent struct {
	client   *llm.OllamaClient
	registry *tools.Registry
	history  []llm.ChatMessage
}

func NewAgent(baseURL, model string) *Agent {
	r := tools.NewRegistry()
	r.Register(tools.ReadFileTool())
	r.Register(tools.WriteFileTool())
	r.Register(tools.ListFilesTool())

	a := &Agent{
		client:   llm.NewOllamaClient(baseURL, model),
		registry: r,
		history:  []llm.ChatMessage{},
	}

	// Identidade do Agente
	a.history = append(a.history, llm.ChatMessage{
		Role:    "system",
		Content: "Você é o Crabe, um Agente de IA local especializado em negócios, empreendedorismo e estratégia. Sua missão é ajudar usuários a validar ideias, pesquisar mercados, criar planos de negócios (SWOT, Canvas, 5W2H) e gerar artefatos de negócio úteis. Seja profissional, analítico e focado em resultados práticos.",
	})

	return a
}

func (a *Agent) Process(input string) (string, error) {
	a.history = append(a.history, llm.ChatMessage{Role: "user", Content: input})

	// Simple reasoning loop (without tool-use for now, just direct chat)
	// We'll add tool-use detection and execution next
	
	resp, err := a.client.Chat(context.Background(), llm.ChatRequest{Model: a.client.Model, Messages: a.history})
	if err != nil {
		return "", err
	}

	a.history = append(a.history, llm.ChatMessage{Role: "assistant", Content: resp.Content})
	return resp.Content, nil
}
