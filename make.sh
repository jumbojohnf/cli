#!/bin/bash

set -euo pipefail

MAKE_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
SCRIPT_DIR="$MAKE_DIR/scripts"

# Source all scripts.
source "$SCRIPT_DIR/deploy.sh"
source "$SCRIPT_DIR/gen_mocks.sh"
source "$SCRIPT_DIR/test.sh"

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
  elif [ $cmd == "test" ]; then
    generate_mocks
    test_all
  elif [ $cmd == "genmocks" ]; then
    generate_mocks
  else
    echo "🛑 Unknown command $cmd"
  fi
}

run $@
