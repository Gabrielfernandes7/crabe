#!/bin/bash

function crabe_set_context() {
  PROJECT_PATH=$1

  export CRABE_PROJECT_PATH="$PROJECT_PATH"

  echo "📂 Contexto definido:"
  echo "   $CRABE_PROJECT_PATH"

  # Persistência local por projeto
  mkdir -p "$PROJECT_PATH/.crabe"
  echo "$PROJECT_PATH" > "$PROJECT_PATH/.crabe/context"
}