# Crabe 🦀 - Project Instructions

This document provides context and guidelines for developing and interacting with the Crabe project.

## Project Overview
Crabe is a modern CLI tool written in Go that simplifies the orchestration of local AI agents. It automates the setup of **Ollama** and **Docker** environments, providing a seamless Developer Experience (DX) for running LLMs within specific project contexts.

### Main Technologies
- **Go 1.26+**: Core programming language.
- **Cobra**: CLI framework for command management.
- **Lipgloss**: UI library for modern terminal styling.
- **Docker & Docker Compose**: For containerized AI service orchestration.
- **Ollama**: Local LLM runner.

### Architecture
- `cmd/crabe/`: The main entry point and command registrations.
- `internal/`: Core business logic partitioned by command:
    - `doctor/`: Diagnostics for Docker, Ollama, and system ports.
    - `initcmd/`: Logic for project-level agent initialization.
    - `inspect/`: System resource analysis (CPU, GPU, RAM).
    - `setup/`: Infrastructure setup (Docker Compose, Ollama configuration).
    - `ui/`: Centralized styling and terminal output helpers.

## Building and Running
The project uses a `Makefile` for common development tasks.

| Command | Description |
|---------|-------------|
| `make build` | Compiles the binary into the root directory. |
| `make install` | Compiles and installs the binary to `~/.local/bin`. |
| `make clean` | Removes generated binaries. |
| `make doctor` | Builds and runs the `crabe doctor` command. |
| `make init` | Builds and runs the `crabe init` command. |

## Development Conventions
- **Command Structure**: New commands should be added as packages under `internal/` and registered in `cmd/crabe/main.go` using the `New<Command>Cmd()` pattern.
- **UI & Styling**: Always use the `internal/ui` package for terminal output to ensure visual consistency. Avoid direct `fmt.Println` for user-facing messages.
- **Error Handling**: Use `ui.Error()` or `ui.Fatal()` for user-facing error reporting.
- **Platform Support**: Primarily targets Linux and macOS. Windows support is in development.
- **Environment Checks**: Use `internal/doctor` logic to verify dependencies before performing operations that require Docker or Ollama.

## Key Files
- `Makefile`: Central task runner.
- `go.mod`: Project dependencies and Go version.
- `cmd/crabe/main.go`: CLI root definition.
- `docker/docker-compose.yml`: Defines the local AI infrastructure.

## Roadmap & TODOs
- [ ] Complete Windows support.
- [ ] Implement `crabe status` to show running services.
- [ ] Enhance model selection during `crabe init`.
