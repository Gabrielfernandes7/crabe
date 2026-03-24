#!/bin/bash

echo "Configurando Crabe..."

CRABE_DIR="$HOME/.crabe"
mkdir -p "$CRABE_DIR"

if [ ! -f "$CRABE_DIR/config.json" ]; then
  echo '{ "model": "llama3.2:3b" }' > "$CRABE_DIR/config.json"
fi

chmod +x cli/crabe.sh

# Instalar comando global
mkdir -p ~/.local/bin
cp cli/crabe.sh ~/.local/bin/crabe

chmod +x ~/.local/bin/crabe

echo "✅ Crabe instalado com sucesso"
echo "👉 Reinicie o terminal se necessário"