#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

printf "\n==> Hexagonal architecture guard\n"

fail() {
  echo "[FAIL] $1"
  exit 1
}

pass() {
  echo "[PASS] $1"
}

if rg -n 'infrastructure/' domain --glob '*.go' >/dev/null 2>&1; then
  fail "domain must not depend on infrastructure"
else
  pass "domain does not import infrastructure"
fi

if rg -nE 'pgxpool|sql\.DB|gorm|bun\.' presentation --glob '*.go' >/dev/null 2>&1; then
  fail "presentation must not access database implementations"
else
  pass "presentation does not access database implementations"
fi

if rg -n 'infrastructure/' application --glob '*.go' >/dev/null 2>&1; then
  fail "application must not depend on infrastructure"
else
  pass "application does not import infrastructure"
fi

if ! rg -n 'domain/ports' application/services --glob '*.go' >/dev/null 2>&1; then
  fail "application services must depend on repository ports"
else
  pass "application services depend on repository ports"
fi

if ! rg -n 'var _ .*ProductRepository' infrastructure/repositories --glob '*.go' >/dev/null 2>&1; then
  fail "infrastructure repositories should assert interface compliance"
else
  pass "repository implementation asserts port compliance"
fi

if rg -n 'services\.New[A-Za-z]+\(' --glob '*.go' --glob '!**/*_test.go' --glob '!tests/**' --glob '!presentation/container/**' . >/dev/null 2>&1; then
  fail "only DI container may wire runtime services (tests are exempt)"
else
  pass "runtime service wiring is restricted to DI container (tests exempt)"
fi

if ! command -v go >/dev/null 2>&1; then
  fail "go toolchain not found"
fi

go test ./... >/dev/null
pass "go test ./..."

echo "Architecture checks completed successfully."
