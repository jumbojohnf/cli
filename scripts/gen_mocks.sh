#!/bin/bash

set -euo pipefail

function generate_mocks
{
  go run "$SCRIPT_DIR/mockgen/main.go" "$MAKE_DIR/mockgen.yaml"
}
