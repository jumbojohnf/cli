#!/bin/bash

set -euo pipefail

source "$SCRIPT_DIR/gen_mocks.sh"

function test_all
{
  generate_mocks

  echo "🧪 Testing all packages"
  current_dir=$(pwd)
  repo_root=$(git rev-parse --show-toplevel)
  cd "$repo_root" && go test ./...
  cd "$current_dir"
}
