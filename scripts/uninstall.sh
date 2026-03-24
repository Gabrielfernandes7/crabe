#!/bin/bash

set -e

echo "🧹 Desinstalando Crabe..."

# 1. Remover comando global
if [ -f "$HOME/.local/bin/crabe" ]; then
  rm "$HOME/.local/bin/crabe"
  echo "✅ CLI removido (~/.local/bin/crabe)"
else
  echo "ℹ️ CLI não encontrado"
fi

# 2. Remover config global
if [ -d "$HOME/.crabe" ]; then
  rm -rf "$HOME/.crabe"
  echo "✅ Configurações removidas (~/.crabe)"
fi

# 3. Parar containers (se existirem)
if command -v docker &> /dev/null; then
  echo "🐳 Parando containers..."
  docker compose down 2>/dev/null || true
fi

echo ""
echo "✅ Crabe desinstalado com sucesso"