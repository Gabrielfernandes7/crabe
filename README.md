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

**Crabe** é uma **CLI moderna em Go** que simplifica drasticamente o uso de agentes de IA locais.

### O problema que resolvemos
Configurar um ambiente de IA local com **Ollama** para trabalhar no contexto do seu projeto atual costuma ser complicado: você precisa configurar volumes, portas, permissões e ferramentas manualmente. Isso gera muita fricção toda vez que você troca de projeto.

### A solução: Crabe
Entre em **qualquer pasta** do seu computador e execute um único comando:

```bash
crabe init
```

Pronto. O Crabe cuida de toda a orquestração e você ganha um **agente de IA inteligente rodando 100% localmente**, entendendo perfeitamente o contexto do projeto atual.

O **Crabe** adiciona a camada de **experiência do desenvolvedor (DX)** que estava faltando: instalação simples, diagnóstico automático, inicialização por projeto e interface bonita no terminal.

---

### Demo

<!--![Crabe Demo](https://i.imgur.com/XXXXXXX.gif)-->

![Crabe Demo](/docs/crabe-tui.png)  
*(Substitua pelo GIF real mostrando `crabe doctor`, `crabe init` e uma conversa com o agente)*

---

## Requisitos

- **Ollama** instalado
- Pelo menos 8 GB de RAM livre (16 GB+ recomendado para modelos maiores)
- Linux / macOS (suporte a Windows em desenvolvimento)

---

## Como instalar (uma única vez)

### Recomendado (com Makefile)

```bash
git clone https://github.com/Gabrielfernandes7/crabe.git
cd crabe
make install
```

Isso compila o binário em Go e coloca o comando `crabe` no seu PATH.

### Alternativa manual

```bash
cd crabe
go build -o crabe ./cmd/crabe
sudo mv crabe /usr/local/bin/
chmod +x /usr/local/bin/crabe
```

---

## Como usar (Fluxo diário)

```bash
# 1. Entre na pasta do seu projeto
cd ~/projetos/meu-app-nextjs

# 2. Inicialize o agente
crabe init
```

Depois disso, o agente já estará pronto. Você pode conversar com ele em linguagem natural dentro daquela pasta (ele terá acesso aos arquivos, git, terminal, etc.).

Para forçar uma reinicialização:
```bash
crabe init --force
```

---

## Comandos principais

| Comando                  | Descrição                                              |
|--------------------------|--------------------------------------------------------|
| `crabe init`             | Inicializa o agente no projeto atual                   |
| `crabe init --force`     | Força a reinicialização (útil para mudanças)           |
| `crabe doctor`           | Diagnóstico completo do ambiente (Docker, Ollama, etc.)|
| `crabe status`           | Mostra serviços rodando e modelo em uso                |
| `crabe version`          | Mostra a versão instalada                              |
| `crabe --help`           | Lista todos os comandos e opções                       |

---

## Modelos recomendados (via Ollama)

- **`qwen2.5-coder:7b`** → **Melhor equilíbrio** performance × consumo (recomendado para a maioria)
- **`qwen2.5-coder:14b`** → Mais inteligente e capaz (exige mais RAM)

O Crabe faz o download automático do modelo escolhido durante o `init`.

---

## Desenvolvimento

```bash
make build      # Compila o binário
make install    # Compila e instala
make doctor     # Executa crabe doctor
make init       # Executa crabe init no diretório atual
make clean      # Remove binários gerados
```

---

## Tecnologias

- **Go** + **Cobra** (CLI robusta)
- **Lipgloss** (interface moderna e colorida no terminal)
- **Ollama** (tudo rodando localmente)

---

## Dicas e Troubleshooting

- Sempre rode `make install` após alterar o código.
- Comando não encontrado? Execute `make remove-old && make install`.
- Rode `crabe doctor` sempre que tiver algum erro.

---

## Para excutar o comando `crabe`

```shell
make build

make install
```

**Pronto para começar?**

```bash
cd seu-projeto
crabe init
```

Agora você tem um agente de IA local trabalhando **exatamente** onde você precisa.

---

## Bibliotecas principais

- [spf13/cobra](https://github.com/spf13/cobra)
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)

---


**Contribuições são bem-vindas!**  
Abra uma issue ou envie um Pull Request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).