# Crabe 🦀

<p align="center">
  <img src="./docs/clawbot-icon.png" width="140" height="140" alt="Crabe icon" style="border-radius: 15px">
</p>

<p align="center">
  <a href="https://github.com/Gabrielfernandes7/crabe/actions">
    <img src="https://img.shields.io/badge/Go-1.23+-00ADD8.svg?logo=go" alt="Go Version">
  </a>
  <a href="https://github.com/Gabrielfernandes7/crabe/releases">
    <img src="https://img.shields.io/github/v/release/Gabrielfernandes7/crabe?color=success" alt="Latest Release">
  </a>
  <img src="https://img.shields.io/badge/CLI%20em%20Go-brightgreen" alt="Built with Go">
  <img src="https://img.shields.io/badge/100%25%20Local-blue" alt="100% Local">
</p>

**Crabe** is a **modern CLI written in Go** that drastically simplifies the use of local AI agents.

### The Problem it Solves
Setting up a powerful AI agent (like **OpenClaw**) with **Ollama** and **Docker** to work within the context of your current project is usually complicated. You have to manually manage Docker Compose, download models, configure volumes, ports, permissions, and tools. This creates significant friction every time you switch projects.

### The Solution: Crabe
Navigate to **any folder** on your computer and run a single command:

```bash
crabe init
```

That's it. Crabe handles all the orchestration, providing you with a **smart AI agent running 100% locally**, with a full understanding of your current project's context.

**Crabe does not replace OpenClaw** — it uses OpenClaw as its core engine while adding the missing **Developer Experience (DX)** layer: simple installation, automatic diagnostics, per-project initialization, and a beautiful terminal interface.

---

### Demo

![Crabe Demo](https://i.imgur.com/XXXXXXX.gif)  
*(Replace with an actual GIF showing `crabe doctor`, `crabe init`, and a conversation with the agent)*

---

## Requirements

- **Docker** installed and running
- **Ollama** installed (Crabe automatically pulls recommended models)
- At least 8 GB of free RAM (16 GB+ recommended for larger models)
- Linux / macOS (Windows support currently in development)

---

## How to Install (One-time setup)

### Recommended (via Makefile)

```bash
git clone https://github.com/Gabrielfernandes7/crabe.git
cd crabe
make install
```

This compiles the Go binary and adds the `crabe` command to your PATH.

### Manual Alternative

```bash
cd crabe
go build -o crabe ./cmd/crabe
sudo mv crabe /usr/local/bin/
chmod +x /usr/local/bin/crabe
```

---

## Usage (Daily Workflow)

```bash
# 1. Enter your project folder
cd ~/projects/my-nextjs-app

# 2. Initialize the agent
crabe init
```

Once initialized, the agent is ready. You can interact with it using natural language within that folder (it will have access to files, git, terminal, etc.).

To force a re-initialization:
```bash
crabe init --force
```

---

## Core Commands

| Command                  | Description                                              |
|--------------------------|----------------------------------------------------------|
| `crabe init`             | Initializes the agent in the current project             |
| `crabe init --force`     | Forces re-initialization (useful for changes)            |
| `crabe doctor`           | Full environment diagnostic (Docker, Ollama, etc.)       |
| `crabe status`           | Shows running services and the model in use              |
| `crabe version`          | Displays the installed version                           |
| `crabe --help`           | Lists all commands and options                           |

---

## Recommended Models (via Ollama)

- **`qwen2.5-coder:7b`** → **Best balance** between performance and resource consumption (recommended for most).
- **`qwen2.5-coder:14b`** → Smarter and more capable (requires more RAM).

Crabe automatically handles the download of the chosen model during `init`.

---

## Development

```bash
make build      # Compiles the binary
make install    # Compiles and installs
make doctor     # Runs crabe doctor
make init       # Runs crabe init in the current directory
make clean      # Removes generated binaries
```

---

## Technologies

- **Go** + **Cobra** (Robust CLI framework)
- **Lipgloss** (Modern and colorful terminal UI)
- **OpenClaw** + **Ollama** + **Docker** (Everything running locally)

---

## Tips & Troubleshooting

- Always run `make install` after modifying the code.
- Command not found? Run `make remove-old && make install`.
- Docker issues? Ensure your user is in the `docker` group (`sudo usermod -aG docker $USER`).
- Run `crabe doctor` whenever you encounter an error.

---

**Ready to start?**

```bash
cd your-project
crabe init
```

Now you have a local AI agent working **exactly** where you need it.

---

## Main Libraries

- [spf13/cobra](https://github.com/spf13/cobra)
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)

---

**Contributions are welcome!** Feel free to open an issue or submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
