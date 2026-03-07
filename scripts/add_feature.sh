#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: ./scripts/add_feature.sh <feature_name>"
  exit 1
fi

go run ./tools/add_feature "$1"
