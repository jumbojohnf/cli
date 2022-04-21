#!/bin/bash

set -euo pipefail

LOCAL_DST="/usr/local/bin/funcgql"

function deploy_local
{
  rm -rf "$LOCAL_DST"
  cd "$MAKE_DIR" && go build .
  mv "$MAKE_DIR/cli" "$LOCAL_DST"
  chmod +x "$LOCAL_DST"
}

function deploy
{
  flag=""
  if [ $# -gt 0 ]; then
    flag=$1
    shift
  fi

  deploy_local $@
  echo "✅ Deployed local "$LOCAL_DST""

  if [[ $flag == "--local" ]]; then
    return
  fi
}
