#!/bin/bash

set -euo pipefail

LOCAL_DST="/usr/local/bin/funcgql"

function deploy_local
{
  source_root=$(git rev-parse --show-toplevel)

  rm -rf "$LOCAL_DST"
  cd "$source_root" && go build .
  mv "$source_root/cli" "$LOCAL_DST"
  chmod +x "$LOCAL_DST"
}

function deploy
{
  flag=""
  if [ $# -gt 0 ]; then
    flag=$1
    shift
  fi

  deploy_local
  echo "✅ Deployed local "$LOCAL_DST""
}
