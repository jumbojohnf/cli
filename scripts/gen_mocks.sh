#!/bin/bash

set -euo pipefail

function generate_mocks
{
  script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
  repo_root=$(git rev-parse --show-toplevel)
  go run "$script_dir/mockgen/main.go" "$repo_root/mockgen.yaml"
}
