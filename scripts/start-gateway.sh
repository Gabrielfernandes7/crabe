#!/bin/bash

set -e

GATEWAY_LOG="$HOME/.openclaw/gateway.log"
GATEWAY_PID="$HOME/.openclaw/gateway.pid"
OPENCLAW_DIR="$HOME/.openclaw"

mkdir -p "$HOME/.openclaw"

echo "🔌 Iniciando gateway do OpenClaw..."

# Evitar duplicação
if [ -f "$GATEWAY_PID" ]; then
  PID=$(cat "$GATEWAY_PID")
  if ps -p $PID > /dev/null 2>&1; then
    echo "✅ Gateway já está rodando (PID $PID)"
    return
  else
    echo "⚠️ PID antigo encontrado, limpando..."
    rm -f "$GATEWAY_PID"
  fi
fi

# Iniciar gateway em background
cd "$OPENCLAW_DIR"

nohup python3 -m openclaw.gateway \
  > "$GATEWAY_LOG" 2>&1 &

PID=$!
echo $PID > "$GATEWAY_PID"

sleep 2

if ps -p $PID > /dev/null; then
  echo "✅ Gateway iniciado (PID $PID)"
else
  echo "❌ Falha ao iniciar gateway"
  exit 1
fi