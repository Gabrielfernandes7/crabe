#!/bin/bash
# scripts/install-model.sh
# Baixa um modelo SLM (Small Language Model) para o Ollama

set -euo pipefail

# CONFIGURAÇÃO DE PATHS
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BASE_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

CRABE_DIR="$HOME/.crabe"
DOCKER_DIR="$BASE_DIR/docker"

# Importar cores (com fallback)
if [ -f "$BASE_DIR/cli/colors.sh" ]; then
  source "$BASE_DIR/cli/colors.sh"
else
  log_info()    { echo "[INFO]    $1"; }
  log_warn()    { echo "[WARN]    $1"; }
  log_error()   { echo "[ERROR]   $1"; }
  log_highlight(){ echo "==> $1"; }
fi

# VALIDAÇÃO
MODEL="${1:-}"
if [ -z "$MODEL" ]; then
  log_error "Uso: crabe install --model <nome_do_modelo>"
  echo
  log_info "Exemplos:"
  log_info "   crabe install --model llama3.2:1b"
  log_info "   crabe install --model qwen2.5-coder:7b"
  log_info "   crabe install --model phi4:mini"
  exit 1
fi

log_highlight "Instalando modelo: $MODEL"

# 1. Garantir que o Ollama esteja rodando
if ! docker ps --filter "name=ollama-cpu" --format "{{.Names}}" | grep -q ollama-cpu; then
  log_info "Container Ollama não está rodando. Iniciando..."
  cd "$DOCKER_DIR" && docker compose up -d ollama
  sleep 4
fi

# 2. Esperar Ollama ficar disponível
log_info "Aguardando Ollama responder..."
for i in {1..25}; do
  if curl -s --max-time 3 http://localhost:11434/api/tags >/dev/null; then
    log_info "Ollama está pronto ✓"
    break
  fi
  sleep 2
done

# 3. Fazer o pull do modelo
log_info "Baixando modelo '$MODEL' (pode demorar vários minutos)..."
docker exec ollama-cpu ollama pull "$MODEL"

# 4. Verificar sucesso
if docker exec ollama-cpu ollama list | grep -q "$MODEL"; then
  log_highlight "✅ Modelo $MODEL instalado com sucesso!"

  # Atualiza configuração global
  mkdir -p "$CRABE_DIR"
  echo "{\"model\": \"$MODEL\"}" > "$CRABE_DIR/config.json"
  log_info "Modelo padrão atualizado em ~/.crabe/config.json"
else
  log_error "Falha ao instalar o modelo $MODEL"
  exit 1
fi

echo
log_info "Agora você pode usar:"
log_info "   crabe model          → ver modelo atual"
log_info "   crabe doctor         → diagnóstico completo"
log_info "   crabe init           → iniciar agente no projeto"
