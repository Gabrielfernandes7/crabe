package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tool struct {
	Name        string
	Description string
	Parameters  interface{}
	Execute     func(args map[string]interface{}) (string, error)
}

type Registry struct {
	tools map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

func (r *Registry) Register(tool Tool) {
	r.tools[tool.Name] = tool
}

func (r *Registry) GetTools() []Tool {
	var list []Tool
	for _, t := range r.tools {
		list = append(list, t)
	}
	return list
}

func (r *Registry) Execute(name string, args map[string]interface{}) (string, error) {
	tool, ok := r.tools[name]
	if !ok {
		return "", fmt.Errorf("tool %s not found", name)
	}
	return tool.Execute(args)
}

// Tool implementations

func ReadFileTool() Tool {
	return Tool{
		Name:        "read_file",
		Description: "Reads the content of a file within the workspace",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{"type": "string"},
			},
			"required": []string{"path"},
		},
		Execute: func(args map[string]interface{}) (string, error) {
			path, ok := args["path"].(string)
			if !ok {
				return "", fmt.Errorf("missing path parameter")
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return "", err
			}
			return string(data), nil
		},
	}
}

func WriteFileTool() Tool {
	return Tool{
		Name:        "write_file",
		Description: "Writes content to a file within the workspace",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":    map[string]interface{}{"type": "string"},
				"content": map[string]interface{}{"type": "string"},
			},
			"required": []string{"path", "content"},
		},
		Execute: func(args map[string]interface{}) (string, error) {
			path, ok := args["path"].(string)
			if !ok {
				return "", fmt.Errorf("missing path parameter")
			}
			content, ok := args["content"].(string)
			if !ok {
				return "", fmt.Errorf("missing content parameter")
			}
			
			// Simple security check: don't allow writing outside workspace if not trusted
			// (This should be handled by a higher level workspace manager)
			
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("File %s written successfully", path), nil
		},
	}
}

func ListFilesTool() Tool {
	return Tool{
		Name:        "list_files",
		Description: "Lists files in a directory",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{"type": "string"},
			},
			"required": []string{"path"},
		},
		Execute: func(args map[string]interface{}) (string, error) {
			path, ok := args["path"].(string)
			if !ok {
				path = "."
			}
			entries, err := os.ReadDir(path)
			if err != nil {
				return "", err
			}
			var files []string
			for _, e := range entries {
				files = append(files, e.Name())
			}
			data, _ := json.Marshal(files)
			return string(data), nil
		},
	}
}
