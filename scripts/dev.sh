#!/usr/bin/env bash
set -euo pipefail

if ! command -v air >/dev/null 2>&1; then
  echo "air is not installed. Install with: go install github.com/air-verse/air@latest"
  exit 1
fi

air
