#!/bin/bash

set -e

CRABE_DIR="$HOME/.crabe"
PROJECT_DIR="$(pwd)"

# Carregar config
CONFIG_FILE="$CRABE_DIR/config.json"

function load_config() {
  LOCAL_CONFIG="$PROJECT_DIR/model.crabe.json"
  GLOBAL_CONFIG="$HOME/.crabe/config.json"

  # 1. Config local (projeto)
  if [ -f "$LOCAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$LOCAL_CONFIG")
    SOURCE="local"
  fi

  # 2. Config global
  if [ -z "${MODEL:-}" ] && [ -f "$GLOBAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$GLOBAL_CONFIG")
    SOURCE="global"
  fi

  # 3. Fallback
  if [ -z "${MODEL:-}" ]; then
    MODEL="llama3.2:3b"
    SOURCE="default"
  fi
}

# Importar módulos
source "$(dirname "$0")/../core/context-resolver.sh"
source "$(dirname "$0")/../scripts/start.sh"
source "$(dirname "$0")/../scripts/stop.sh"
source "$(dirname "$0")/../scripts/doctor.sh"

# Comandos
COMMAND=$1

load_config

case $COMMAND in
  init)
  echo "🦞 Crabe iniciando..."

  crabe_start

  crabe_set_context "$PROJECT_DIR"

  echo ""
  echo "🧠 Modelo: $MODEL ($SOURCE)"
  echo "📂 Projeto: $PROJECT_DIR"
  echo "🔌 Gateway: ativo"
  echo ""
  echo "✅ Crabe pronto"
  ;;

  status)
    crabe_status
    ;;

  stop)
    crabe_stop
    ;;

  doctor)
    crabe_doctor
    ;;

  model)
    NEW_MODEL=$2
    if [ -z "$NEW_MODEL" ]; then
      echo "Modelo atual: $MODEL"
    else
      mkdir -p "$CRABE_DIR"
      echo "{ \"model\": \"$NEW_MODEL\" }" > "$CONFIG_FILE"
      echo "✅ Modelo alterado para: $NEW_MODEL"
    fi
    ;;

  *)
    echo "Uso: crabe {init|status|stop|doctor|model}"
    ;;
esac