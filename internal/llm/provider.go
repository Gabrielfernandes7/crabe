package llm

import (
	"context"
)

// ChatRequest defines the structure for a chat request.
type ChatRequest struct {
	Model    string
	Messages []ChatMessage // Using ChatMessage from ollama.go for now, or define a new one here.
}

// ChatMessage defines a generic chat message.
type ChatMessage struct {
	Role    string
	Content string
}

// ChatResponse defines the structure for a chat response.
type ChatResponse struct {
	Content string
}

// LLMProvider defines the interface for interacting with an LLM.
type LLMProvider interface {
	Name() string
	Health(ctx context.Context) error
	Chat(ctx context.Context, request ChatRequest) (ChatResponse, error)
}
