#!/bin/bash

set -euo pipefail

# Source all scripts.
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
for script in "$SCRIPT_DIR/scripts/*.sh"; do source $script; done

function run
{
  if [ $# -lt 1 ]; then
    echo "🛑 Missing command"
    exit 1
  fi

  cmd=$1
  shift

  if [ $cmd == "deploy" ]; then
    deploy $@
  else
    echo "🛑 Unknown command $cmd"
  fi
}

run $@
