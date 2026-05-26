#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

echo "==> Compiling WASM test binary..."
GOOS=js GOARCH=wasm go test -c -o "$TMPDIR/test.wasm" "$ROOT/wgpu/"

echo "==> Copying wasm_exec.js..."
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$TMPDIR/wasm_exec.js"

echo "==> Running tests in headless browser..."
node "$SCRIPT_DIR/wasm_test_runner.mjs" "$TMPDIR/test.wasm" "$TMPDIR/wasm_exec.js" "${@-}"
