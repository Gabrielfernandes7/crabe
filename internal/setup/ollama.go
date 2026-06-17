package setup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

const defaultModel = "qwen2.5:3b"
const ollamaBaseURL = "http://localhost:11434"

type modelList struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

func listModels() []string {
	resp, err := http.Get(ollamaBaseURL + "/api/tags")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var ml modelList
	if err := json.NewDecoder(resp.Body).Decode(&ml); err != nil {
		return nil
	}

	var models []string
	for _, m := range ml.Models {
		models = append(models, m.Name)
	}

	return models
}

func EnsureModel(models []string) (string, error) {
	if len(models) > 0 {
		return models[0], nil
	}

	ui.Section("Modelo")
	ui.Info("Baixando modelo padrão: %s", defaultModel)

	// Ollama pull API
	payload := map[string]string{"name": defaultModel}
	jsonData, _ := json.Marshal(payload)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", ollamaBaseURL+"/api/pull", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to pull model: %d", resp.StatusCode)
	}

	return defaultModel, nil
}