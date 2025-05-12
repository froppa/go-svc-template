#!/usr/bin/env bash
set -euo pipefail

# -----------------------------------------------------------------------------
# Usage:
#   ./scripts/update-deps.sh
# Regenerates go_repository rules from go.mod via Gazelle.
# -----------------------------------------------------------------------------

cd $(git rev-parse --show-toplevel)

echo "→ Cleaning (bazel clean --expunge)"
bazel clean --expunge

echo
echo "→ Tidying Bazel modules (bazel mod tidy)"
bazel mod tidy

echo
echo "→ Regenerating go_repository rules (bazel run //:gazelle)"
bazel run //:gazelle
