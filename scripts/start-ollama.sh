#!/bin/bash
# =============================================
# start-ollama.sh - Gerencia Ollama + escolha de modelo de código
# =============================================

set -euo pipefail

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

echo -e "${CYAN}🦞 Iniciando Ollama + Open-WebUI...${NC}"

# Verifica permissão Docker
if ! docker ps >/dev/null 2>&1; then
    echo -e "${RED}❌ Erro de permissão no Docker!${NC}"
    echo -e "${YELLOW}Rode uma vez:${NC} sudo usermod -aG docker \$USER && newgrp docker"
    exit 1
fi
echo -e "${GREEN}✅ Docker OK${NC}"

# Verifica/inicia containers
if docker ps --filter "name=ollama-cpu" --format "{{.Names}}" | grep -q "ollama-cpu"; then
    echo -e "${GREEN}✅ Ollama já está rodando.${NC}"
else
    echo -e "${YELLOW}⚠️  Iniciando Ollama com docker compose...${NC}"
    if [ ! -f docker-compose.yml ]; then
        echo -e "${RED}❌ docker-compose.yml não encontrado!${NC}"
        exit 1
    fi
    docker compose up -d --remove-orphans
    echo -e "${YELLOW}⏳ Aguardando 25 segundos...${NC}"
    sleep 25
fi

# Testa conexão
if curl -s http://127.0.0.1:11434/api/tags >/dev/null; then
    echo -e "${GREEN}✅ Ollama respondendo na porta 11434${NC}"
else
    echo -e "${RED}❌ Ollama não respondeu. Verifique: docker logs ollama-cpu${NC}"
    exit 1
fi

# Menu de escolha de modelo SLM para código
echo -e "\n${PURPLE}Escolha o modelo de código (SLM) para usar com o Crabe:${NC}"
echo -e "1) ${GREEN}qwen2.5-coder:7b${NC}   → Recomendado (rápido e bom para projetos médios)"
echo -e "2) ${GREEN}qwen2.5-coder:14b${NC}  → Mais inteligente (usa mais RAM)"
echo -e "3) ${GREEN}deepseek-coder-v2:16b${NC} → Excelente em raciocínio"
echo -e "4) ${YELLOW}glm-4.7-flash${NC}     → Já está baixado (geral, não o melhor para código)"
echo -e "5) ${YELLOW}Não baixar agora${NC} (usar o que já tenho)"
echo -e "${CYAN}Digite o número da opção:${NC} "
read -r choice

case $choice in
    1)
        echo -e "${YELLOW}Baixando qwen2.5-coder:7b...${NC}"
        docker exec ollama-cpu ollama pull qwen2.5-coder:7b
        MODEL="qwen2.5-coder:7b"
        ;;
    2)
        echo -e "${YELLOW}Baixando qwen2.5-coder:14b...${NC}"
        docker exec ollama-cpu ollama pull qwen2.5-coder:14b
        MODEL="qwen2.5-coder:14b"
        ;;
    3)
        echo -e "${YELLOW}Baixando deepseek-coder-v2:16b...${NC}"
        docker exec ollama-cpu ollama pull deepseek-coder-v2:16b
        MODEL="deepseek-coder-v2:16b"
        ;;
    4)
        MODEL="glm-4.7-flash"
        ;;
    5)
        echo -e "${YELLOW}Pulando download.${NC}"
        MODEL=""
        ;;
    *)
        echo -e "${RED}Opção inválida. Usando qwen2.5-coder:7b${NC}"
        docker exec ollama-cpu ollama pull qwen2.5-coder:7b
        MODEL="qwen2.5-coder:7b"
        ;;
esac

# Lista modelos atuais
echo -e "\n${BLUE}📋 Modelos disponíveis:${NC}"
docker exec ollama-cpu ollama list

echo -e "\n${GREEN}✅ Ollama pronto!${NC}"
echo -e "   Open-WebUI: http://localhost:3000"
if [ -n "$MODEL" ]; then
    echo -e "   Modelo selecionado: ${PURPLE}$MODEL${NC}"
fi
echo -e "\n${CYAN}Próximo passo: vamos configurar o Crabe (OpenClaw) com o modelo escolhido.${NC}"

# Salvar modelo no config do Crabe
CRABE_DIR="$HOME/.crabe"
mkdir -p "$CRABE_DIR"

if [ -n "$MODEL" ]; then
    echo "{ \"model\": \"$MODEL\" }" > "$CRABE_DIR/config.json"
    echo -e "${GREEN}💾 Modelo salvo no Crabe: $MODEL${NC}"
fi