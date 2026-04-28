package setup

type SystemState struct {
	DockerAvailable bool
	DockerRunning   bool
	OllamaRunning   bool
	Models          []string
}
