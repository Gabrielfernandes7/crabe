# Repository Guidelines

## Project Structure & Module Organization
Crabe is a Go CLI module (`github.com/Gabrielfernandes7/crabe`). The executable entry point is `cmd/crabe/main.go`, where Cobra commands are registered. Command and feature logic lives under `internal/`, grouped by responsibility: `doctor/`, `initcmd/`, `inspect/`, `setup/`, `install/`, `uninstall/`, `workspace/`, `llm/`, `tools/`, and shared terminal UI helpers in `internal/ui/`. Project images and screenshots are in `docs/`. Root-level generated binaries are named `crabe` and should not be committed.

## Build, Test, and Development Commands
- `make build`: compiles the CLI to `./crabe` using `go build -o crabe ./cmd/crabe`.
- `make install`: removes old local copies, builds, and installs to `~/.local/bin/crabe`.
- `make doctor`: builds and runs `./crabe doctor` for local environment diagnostics.
- `make init` / `make init-force`: builds and runs initialization flows.
- `make clean`: removes the local binary and runs `go clean`.
- `go test ./...`: runs all Go tests. Use this before opening a PR, even when adding only small command changes.

## Coding Style & Naming Conventions
Use standard Go formatting: run `gofmt` on changed `.go` files. Keep package names short, lowercase, and aligned with the feature directory. Add new CLI commands as focused packages under `internal/<command>/` and expose constructors such as `NewDoctorCmd()` or `NewInitCmd()`, then register them in `cmd/crabe/main.go`. Prefer `internal/ui` helpers for user-facing terminal output instead of direct `fmt.Println`.

## Testing Guidelines
Place tests next to implementation files using Go’s `*_test.go` convention. Favor table-driven tests for command parsing, setup decisions, and environment inspection logic. When behavior touches the CLI surface, include coverage for success and failure paths. There is no dedicated coverage threshold yet; at minimum, new logic should have focused unit tests and pass `go test ./...`.

## Commit & Pull Request Guidelines
Recent history uses short imperative summaries, sometimes with conventional prefixes such as `refactor:`. Keep commits concise and scoped, for example `refactor: simplify setup preflight`. Pull requests should include a clear description, test results (`go test ./...`, relevant `make` commands), linked issues when applicable, and screenshots or terminal output when changing user-facing CLI/UI behavior.

## Security & Configuration Tips
Do not commit local machine artifacts, generated binaries, or user-specific Crabe config. Treat commands that inspect the host system or call Ollama as local-only operations and document any new external dependency in `README.md`.
