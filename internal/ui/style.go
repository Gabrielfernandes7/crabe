// internal/ui/style.go
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Cores amarelas inspiradas no Claude Code (quente, moderno e clean)
	primary = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00")) // amarelo Claude
	success = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9D")) // verde neon suave
	warning = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00")) // amarelo (mesmo do primary)
	error   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4D4D")) // vermelho
	info    = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")) // cinza claro

	// Estilos de caixas (igual Claude Code)
	titleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFCC00")).
			Padding(0, 2).
			MarginBottom(1)

	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#1A1A1A")).     // texto escuro para bom contraste
		Background(lipgloss.Color("#FFCC00")).     // fundo amarelo Claude
		Padding(0, 1)

	// Emoji + texto
	crab = "🦀"
)

// Init deve ser chamado no início de todo comando
func Init() {
	// Detecta automaticamente o melhor perfil de cores do terminal (truecolor quando possível)
	// ColorProfile() returns the current terminal's color profile
}

// Title — usado no topo de cada comando (igual Claude)
func Title(text string) {
	fmt.Println(titleBox.Render(fmt.Sprintf("%s  %s", crab, text)))
}

// Success — mensagem de sucesso com check
func Success(msg string, args ...any) {
	fmt.Println(success.Render(fmt.Sprintf(msg, args...)))
}

// Error — mensagem de erro
func Error(msg string, args ...any) {
	fmt.Println(error.Render(fmt.Sprintf(msg, args...)))
}

// Info — informação neutra
func Info(msg string, args ...any) {
	fmt.Println(info.Render(fmt.Sprintf(msg, args...)))
}

// Highlight — destaque (usado em status, modelo atual, etc.)
func Highlight(msg string, args ...any) {
	fmt.Println(primary.Render(fmt.Sprintf(msg, args...)))
}

// Warning
func Warning(msg string, args ...any) {
	fmt.Println(warning.Render(fmt.Sprintf(msg, args...)))
}

// Section — separa seções
func Section(title string) {
	fmt.Println()
	fmt.Println(header.Render(" "+title+" "))
	fmt.Println(strings.Repeat("─", 50))
}

// RenderHeader — cabeçalho estilizado para o terminal interativo
func RenderHeader() string {
	var s strings.Builder

	logo := ` ▗▄▄▖   Crabe CLI
 ▐▌     Business Agent
 ▝▀▄▄▖  Local AI`

	s.WriteString(primary.Render(logo))
	s.WriteString("\n")
	s.WriteString(strings.Repeat("─", 40))
	s.WriteString("\n")
	
	return s.String()
}