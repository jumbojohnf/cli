#!/bin/bash

set -euo pipefail

function deploy_local
{
  rm -rf "/usr/local/bin/funcgql"
  cd "$SCRIPT_DIR" && go build .
  mv "$SCRIPT_DIR/cli" "/usr/local/bin/funcgql"
  chmod +x "/usr/local/bin/funcgql"
}

function deploy
{
  if [ $# -lt 1 ]; then
    echo "🛑 Missing deploy target"
    exit 1
  fi

  cmd=$1
  shift

  if [ $cmd == "local" ]; then
    deploy_local $@
    echo "✅ Deployed local funcgql"
  else
    echo "🤨 Unknown deploy target $cmd"
  fi
}
