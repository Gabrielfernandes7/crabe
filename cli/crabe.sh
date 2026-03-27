#!/bin/bash

set -euo pipefail

# Resolve caminho real mesmo via symlink
SCRIPT_PATH="$(realpath "$0" 2>/dev/null || greadlink -f "$0")"
BASE_DIR="$(cd "$(dirname "$SCRIPT_PATH")/.." && pwd)"

CRABE_DIR="$HOME/.crabe"
PROJECT_DIR="$(pwd)"

# Paths
CORE="$BASE_DIR/core/context-resolver.sh"
START="$BASE_DIR/scripts/start.sh"
DOCTOR="$BASE_DIR/scripts/doctor.sh"
SETUP_OPENCLAW="$BASE_DIR/scripts/setup-openclaw.sh"
START_OLLAMA="$BASE_DIR/scripts/start-ollama.sh"
UNINSTALL="$BASE_DIR/scripts/uninstall.sh"
COLORS="$BASE_DIR/cli/colors.sh"

CONFIG_FILE="$CRABE_DIR/config.json"

# Importar cores com fallback
if [ -f "$COLORS" ]; then
  source "$COLORS"
else
  log_info() { echo "[INFO] $1"; }
  log_warn() { echo "[WARN] $1"; }
  log_error() { echo "[ERROR] $1"; }
  log_highlight() { echo "==> $1"; }
fi

# -------- VALIDAÇÕES --------
check_dependencies() {
  command -v jq >/dev/null || { log_error "jq não está instalado"; exit 1; }

  [ -f "$CORE" ] || { log_error "core não encontrado: $CORE"; exit 1; }
  [ -f "$START" ] || { log_error "start não encontrado: $START"; exit 1; }
  [ -f "$DOCTOR" ] || { log_error "doctor não encontrado: $DOCTOR"; exit 1; }
}

# -------- CONFIG --------
load_config() {
  local LOCAL_CONFIG="$PROJECT_DIR/model.crabe.json"
  local GLOBAL_CONFIG="$CONFIG_FILE"

  MODEL=""
  SOURCE=""

  if [ -f "$LOCAL_CONFIG" ]; then
    MODEL="$(jq -r '.model // empty' "$LOCAL_CONFIG")"
    SOURCE="local"
  fi

  if [ -z "$MODEL" ] && [ -f "$GLOBAL_CONFIG" ]; then
    MODEL="$(jq -r '.model // empty' "$GLOBAL_CONFIG")"
    SOURCE="global"
  fi

  if [ -z "$MODEL" ]; then
    MODEL="llama3.2:3b"
    SOURCE="default"
  fi
}

save_model() {
  mkdir -p "$CRABE_DIR"
  jq -n --arg model "$1" '{model: $model}' > "$CONFIG_FILE"
}

# -------- IMPORTS --------
check_dependencies
source "$CORE"
source "$START"
source "$DOCTOR"

# -------- COMMAND DISPATCH --------
COMMAND="${1:-}"

[ -n "$COMMAND" ] || {
  log_warn "Uso: crabe {init|status|doctor|model|version|install|uninstall}"
  exit 1
}

load_config

case "$COMMAND" in
  init)
    log_highlight "Crabe iniciando..."

    crabe_start
    crabe_set_context "$PROJECT_DIR"

    echo
    log_info "Modelo: $MODEL ($SOURCE)"
    log_info "Projeto: $PROJECT_DIR"
    log_info "Crabe pronto"
    ;;

  status)
    crabe_status
    ;;

  doctor)
    crabe_doctor
    ;;

  model)
    if [ -z "${2:-}" ]; then
      log_info "Modelo atual: $MODEL ($SOURCE)"
    else
      save_model "$2"
      log_info "Modelo alterado para: $2"
    fi
    ;;

    install)
      shift
      case "${1:-}" in
        openclaw)
          log_highlight "Instalando OpenClaw..."
          bash "$SETUP_OPENCLAW"
          ;;

        ollama|--model)
          if [ "${1}" = "--model" ]; then
            MODEL_NAME="${2:-}"
          else
            MODEL_NAME="${2:-}"
          fi

          if [ -z "$MODEL_NAME" ]; then
            log_error "Faltou o nome do modelo."
            log_info "Uso: crabe install --model <nome>"
            exit 1
          fi

          bash "$BASE_DIR/scripts/install-model.sh" "$MODEL_NAME"
          ;;

        *)
          log_warn "Uso: crabe install {openclaw | --model <nome_do_modelo>}"
          log_info "Exemplo: crabe install --model qwen2.5-coder:7b"
          exit 1
          ;;
      esac
    ;;

  uninstall)
    bash "$UNINSTALL" "${2:-}"
    ;;

  version)
    log_highlight "Crabe version 26.1"
    echo "Base dir: $BASE_DIR"
    ;;

  *)
    log_warn "Uso: crabe {init|status|doctor|model|version|install|uninstall}"
    exit 1
    ;;
esac