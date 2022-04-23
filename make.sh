#!/bin/bash

set -euo pipefail

# Source all scripts.
MAKE_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
SCRIPT_DIR="$MAKE_DIR/scripts"
for script in "$SCRIPT_DIR/*.sh"; do source $script; done

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
