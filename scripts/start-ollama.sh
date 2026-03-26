#!/bin/bash

set -euo pipefail

# CORES
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

log() { echo -e "$1"; }

# PARSE DE ARGUMENTOS
MODELS=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --model)
      if [[ -n "${2:-}" ]]; then
        MODELS+=("$2")
        shift 2
      else
        log "${RED}Erro: --model requer um valor${NC}"
        exit 1
      fi
      ;;
    *)
      log "${RED}Argumento desconhecido: $1${NC}"
      exit 1
      ;;
  esac
done

# INIT
log "${CYAN}🦞 Iniciando Ollama + Open-WebUI...${NC}"

# DOCKER CHECK
if ! docker ps >/dev/null 2>&1; then
    log "${RED}❌ Docker não está acessível${NC}"
    log "${YELLOW}Execute:${NC} sudo usermod -aG docker \$USER && newgrp docker"
    exit 1
fi
log "${GREEN}✅ Docker OK${NC}"

# VALIDAR DOCKER-COMPOSE
if [ ! -f docker-compose.yml ]; then
    log "${RED}❌ docker-compose.yml não encontrado${NC}"
    exit 1
fi

# SUBIR CONTAINER SE NECESSÁRIO
if docker ps --format "{{.Names}}" | grep -q "^ollama-cpu$"; then
    log "${GREEN}✅ Ollama já está rodando${NC}"
else
    log "${YELLOW}⚠️ Subindo containers...${NC}"
    docker compose up -d --remove-orphans
fi

# ESPERA INTELIGENTE
log "${YELLOW}⏳ Aguardando Ollama iniciar...${NC}"
until curl -s http://127.0.0.1:11434/api/tags >/dev/null; do
    sleep 2
done
log "${GREEN}✅ Ollama pronto na porta 11434${NC}"

# MODO INTERATIVO SE NÃO PASSAR --model
if [ ${#MODELS[@]} -eq 0 ]; then
    echo
    log "${PURPLE}Selecione modelos de código (multi-select):${NC}"
    log "1) ${GREEN}qwen2.5-coder:7b${NC}"
    log "2) ${GREEN}deepseek-coder:6.7b${NC}"
    log "3) ${GREEN}codellama:7b${NC}"
    log "4) ${YELLOW}llama3.2:3b${NC}"
    log "5) ${YELLOW}llama3.2:1b${NC}"
    log "6) ${CYAN}Pular${NC}"

    echo
    read -p "Digite os números separados por espaço: " -ra choices

    for choice in "${choices[@]}"; do
      case $choice in
        1) MODELS+=("qwen2.5-coder:7b") ;;
        2) MODELS+=("deepseek-coder:6.7b") ;;
        3) MODELS+=("codellama:7b") ;;
        4) MODELS+=("llama3.2:3b") ;;
        5) MODELS+=("llama3.2:1b") ;;
        6) ;;
        *) log "${RED}Opção inválida: $choice${NC}" ;;
      esac
    done
fi

# DOWNLOAD MODELOS
if [ ${#MODELS[@]} -gt 0 ]; then
    for model in "${MODELS[@]}"; do
        if docker exec ollama-cpu ollama list | grep -q "$model"; then
            log "${GREEN}Modelo já existe: $model${NC}"
        else
            log "${YELLOW}⬇ Baixando $model...${NC}"
            docker exec ollama-cpu ollama pull "$model"
        fi
    done
else
    log "${YELLOW}Nenhum modelo selecionado${NC}"
fi

# LISTAR MODELOS
echo
log "${BLUE}📋 Modelos disponíveis:${NC}"
docker exec ollama-cpu ollama list

# SALVAR CONFIG NO CRABE
CRABE_DIR="$HOME/.crabe"
mkdir -p "$CRABE_DIR"

if [ ${#MODELS[@]} -gt 0 ]; then
    JSON_MODELS=$(printf '"%s",' "${MODELS[@]}" | sed 's/,$//')
    echo "{ \"models\": [$JSON_MODELS] }" > "$CRABE_DIR/config.json"

    log "${GREEN} Modelos salvos no Crabe${NC}"
fi

# FINAL
echo
log "${GREEN}✅ Ambiente pronto!${NC}"
log "Open-WebUI: http://localhost:3000"

if [ ${#MODELS[@]} -gt 0 ]; then
    log "Modelos ativos:"
    for m in "${MODELS[@]}"; do
        log "  - ${PURPLE}$m${NC}"
    done
fi

echo
log "${CYAN}Próximo passo: usar com Crabe${NC}"