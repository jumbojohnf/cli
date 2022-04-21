#!/bin/bash

set -euo pipefail

function test_all
{
  echo "🧪 Testing all packages"
  current_dir=$(pwd)
  repo_root=$(git rev-parse --show-toplevel)
  cd "$repo_root" && go test ./...
  cd "$current_dir"
}
