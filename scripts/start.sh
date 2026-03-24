#!/bin/bash

CRABE_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

function crabe_start() {
  echo "🚀 Inicializando ambiente Crabe..."

  # 1. Docker
  if ! command -v docker &> /dev/null; then
    echo "❌ Docker não instalado"
    exit 1
  fi

  # 2. Ollama
  echo "🧠 Iniciando Ollama..."
  (cd "$CRABE_ROOT" && bash scripts/start-ollama.sh)

  if ! curl -s http://127.0.0.1:11434/api/tags >/dev/null; then
    echo "❌ Ollama não respondeu"
    exit 1
  fi

  # 3. OpenClaw (instalar se necessário)
  if [ ! -d "$HOME/.openclaw" ]; then
    echo "📦 Instalando OpenClaw..."
    bash "$CRABE_ROOT/scripts/setup-openclaw.sh"
  else
    echo "✅ OpenClaw já instalado"
  fi

  # 4. Subir OpenClaw
  echo "🐳 Subindo OpenClaw..."
  (cd "$HOME/.openclaw" && docker compose up -d)

  # 5. Gateway
  echo "🔌 Iniciando gateway..."
  bash "$CRABE_ROOT/scripts/start-gateway.sh"

  echo "✅ Ambiente completo pronto"
}