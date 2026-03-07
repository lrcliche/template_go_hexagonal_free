#!/usr/bin/env bash
set -euo pipefail

if [ -n "$(gofmt -l .)" ]; then
  echo "[FAIL] gofmt found unformatted files"
  gofmt -l .
  exit 1
fi

echo "[PASS] gofmt check"

go vet ./...
echo "[PASS] go vet ./..."

go test ./...
echo "[PASS] go test ./..."
