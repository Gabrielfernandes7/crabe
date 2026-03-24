#!/bin/bash

function crabe_doctor() {
  echo "🩺 Diagnóstico do sistema"

  echo ""
  echo "🔹 Docker:"
  if command -v docker &> /dev/null; then
    echo "✅ Instalado"
  else
    echo "❌ Não encontrado"
  fi

  echo ""
  echo "🔹 Docker Compose:"
  if docker compose version &> /dev/null; then
    echo "✅ OK"
  else
    echo "❌ Problema no compose"
  fi

  echo ""
  echo "🔹 Ollama container:"
  if docker ps | grep -q "ollama"; then
    echo "✅ Rodando"
  else
    echo "❌ Não está rodando"
  fi

  echo ""
  echo "🔹 Porta 11434 (Ollama):"
  if lsof -i:11434 &> /dev/null; then
    echo "✅ Ativa"
  else
    echo "❌ Inativa"
  fi
}