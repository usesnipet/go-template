#!/usr/bin/env sh
set -eu

if [ "${SKIP_BUILD:-}" = "1" ] || [ "${SKIP_WEB_BUILD:-}" = "1" ]; then
  echo "Build skipped (SKIP_BUILD or SKIP_WEB_BUILD=1)."
  exit 0
fi

root="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
cd "$root"

echo "Running production build…"
pnpm run build

echo "Build passed."
