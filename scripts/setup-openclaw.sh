#!/bin/bash

set -euo pipefail

# Base path (caso rode via symlink)
SCRIPT_PATH="$(readlink -f "$0")"
BASE_DIR="$(dirname "$SCRIPT_PATH")/.."

# Importar cores
source "$BASE_DIR/cli/colors.sh"

# CONFIG
INSTALL_DIR="$HOME/.openclaw"
REPO_URL="https://github.com/openclaw/openclaw.git"

CONFIG_DIR="$HOME/.openclaw/config"
WORKSPACE_DIR="$HOME/.openclaw/workspace"

log_highlight "🦞 Configurando OpenClaw..."

# 1. Verificar dependências

check_command() {
  if ! command -v "$1" &>/dev/null; then
    log_error "Dependência não encontrada: $1"
    return 1
  fi
}

echo "Verificando dependências..."

MISSING=0

check_command git || MISSING=1
check_command docker || MISSING=1
check_command curl || MISSING=1

if [ "$MISSING" -eq 1 ]; then
  log_warn "Instale as dependências antes de continuar."
  exit 1
fi

log_info "Dependências OK"

# 2. Detectar Docker Compose

if docker compose version &>/dev/null; then
  DOCKER_COMPOSE="docker compose"
elif command -v docker-compose &>/dev/null; then
  DOCKER_COMPOSE="docker-compose"
else
  log_error "Docker Compose não encontrado"
  exit 1
fi

log_info "Docker Compose detectado"

# 3. Clonar ou atualizar

if [ -d "$INSTALL_DIR" ]; then
  log_warn "⚠️ OpenClaw já existe em $INSTALL_DIR"

  cd "$INSTALL_DIR"

  if [ -d ".git" ]; then
    log_highlight "Atualizando repositório..."

    git pull || {
      log_error "Falha ao atualizar repositório"
      exit 1
    }

    log_info "OpenClaw atualizado"
  else
    log_error "Diretório não é um repositório git válido"
    log_warn "Sugestão: rm -rf $INSTALL_DIR"
    exit 1
  fi
else
  log_highlight "Baixando OpenClaw..."

  git clone "$REPO_URL" "$INSTALL_DIR" || {
    log_error "Falha ao clonar repositório"
    exit 1
  }

  log_info "OpenClaw baixado com sucesso"
fi

cd "$INSTALL_DIR"

# 4. Validar estrutura

if [ ! -f "docker-compose.yml" ]; then
  log_error "docker-compose.yml não encontrado"
  exit 1
fi

log_info "Estrutura OK"

# 5. Criar diretórios obrigatórios (CORREÇÃO PRINCIPAL)

log_highlight "📁 Preparando diretórios..."

mkdir -p "$CONFIG_DIR"
mkdir -p "$WORKSPACE_DIR"

log_info "Diretórios criados"

# 6. Criar/ajustar .env (CORREÇÃO PRINCIPAL)

log_highlight "⚙️ Configurando ambiente..."

if [ -f ".env.example" ] && [ ! -f ".env" ]; then
  cp .env.example .env
  log_info ".env criado a partir do .env.example"
fi

# Garantir variáveis obrigatórias
touch .env

# Remove entradas antigas (evita duplicação)
sed -i '/OPENCLAW_CONFIG_DIR/d' .env
sed -i '/OPENCLAW_WORKSPACE_DIR/d' .env

# Adiciona corretamente
cat <<EOF >> .env
OPENCLAW_CONFIG_DIR=$CONFIG_DIR
OPENCLAW_WORKSPACE_DIR=$WORKSPACE_DIR
EOF

log_info "Variáveis OPENCLAW configuradas"

# 7. Perguntar para iniciar

log_highlight "🚀 Deseja iniciar o OpenClaw agora? (s/n)"
read -r START_NOW

if [[ "$START_NOW" =~ ^[Ss]$ ]]; then
  log_highlight "🐳 Subindo containers..."

  if $DOCKER_COMPOSE up -d --build; then
    log_info "OpenClaw iniciado"
  else
    log_error "Falha ao subir containers"
    log_warn "📜 Logs para debug:"
    $DOCKER_COMPOSE logs
    exit 1
  fi
else
  log_warn "⏭️ Inicialização ignorada"
fi

# FINAL

echo ""
log_highlight "🎉 OpenClaw pronto!"
log_info "📂 Local: $INSTALL_DIR"
log_info "👉 Para iniciar manualmente:"
echo "cd $INSTALL_DIR && $DOCKER_COMPOSE up -d"