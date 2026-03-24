#!/bin/bash

set -euo pipefail

# ==============================
# CONFIG
# ==============================
INSTALL_DIR="$HOME/.openclaw"
REPO_URL="https://github.com/clawbot/openclaw.git"

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}🦞 Configurando OpenClaw...${NC}"

# ==============================
# 1. Verificar dependências
# ==============================

function check_command() {
  if ! command -v "$1" &> /dev/null; then
    echo -e "${RED}❌ Dependência não encontrada: $1${NC}"
    return 1
  fi
}

echo -e "${CYAN}🔍 Verificando dependências...${NC}"

MISSING=0

check_command git || MISSING=1
check_command docker || MISSING=1
check_command curl || MISSING=1

if [ "$MISSING" -eq 1 ]; then
  echo -e "${YELLOW}⚠️ Instale as dependências antes de continuar.${NC}"
  exit 1
fi

echo -e "${GREEN}✅ Dependências OK${NC}"

# ==============================
# 2. Verificar instalação existente
# ==============================

if [ -d "$INSTALL_DIR" ]; then
  echo -e "${YELLOW}⚠️ OpenClaw já existe em $INSTALL_DIR${NC}"

  cd "$INSTALL_DIR"

  # Verifica se é um repositório git válido
  if [ -d ".git" ]; then
    echo -e "${CYAN}🔄 Atualizando repositório...${NC}"

    if git pull; then
      echo -e "${GREEN}✅ OpenClaw atualizado${NC}"
    else
      echo -e "${RED}❌ Falha ao atualizar (problema de rede?)${NC}"
      exit 1
    fi
  else
    echo -e "${RED}❌ Diretório existe mas não é um repositório git válido${NC}"
    echo -e "${YELLOW}Sugestão: remover manualmente:${NC} rm -rf $INSTALL_DIR"
    exit 1
  fi

else
  # ==============================
  # 3. Clonar repositório
  # ==============================

  echo -e "${CYAN}📥 Baixando OpenClaw...${NC}"

  if git clone "$REPO_URL" "$INSTALL_DIR"; then
    echo -e "${GREEN}✅ OpenClaw baixado com sucesso${NC}"
  else
    echo -e "${RED}❌ Falha ao clonar repositório${NC}"
    echo -e "${YELLOW}Verifique conexão com internet ou acesso ao GitHub${NC}"
    exit 1
  fi
fi

# ==============================
# 4. Verificar arquivos críticos
# ==============================

cd "$INSTALL_DIR"

if [ ! -f "docker-compose.yml" ]; then
  echo -e "${RED}❌ docker-compose.yml não encontrado no OpenClaw${NC}"
  exit 1
fi

echo -e "${GREEN}✅ Estrutura do OpenClaw OK${NC}"

# ==============================
# 5. Subir serviços (opcional)
# ==============================

echo -e "${CYAN}🚀 Deseja iniciar o OpenClaw agora? (s/n)${NC}"
read -r START_NOW

if [[ "$START_NOW" =~ ^[Ss]$ ]]; then
  echo -e "${CYAN}🐳 Subindo OpenClaw...${NC}"

  if docker compose up -d; then
    echo -e "${GREEN}✅ OpenClaw iniciado${NC}"
  else
    echo -e "${RED}❌ Falha ao subir containers${NC}"
    exit 1
  fi
else
  echo -e "${YELLOW}⏭️ Inicialização ignorada${NC}"
fi

# ==============================
# FINAL
# ==============================

echo ""
echo -e "${GREEN}🎉 OpenClaw pronto!${NC}"
echo -e "📂 Local: $INSTALL_DIR"
echo -e "👉 Para iniciar manualmente:"
echo -e "   cd $INSTALL_DIR && docker compose up -d"