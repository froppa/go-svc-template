#!/usr/bin/env bash
set -euo pipefail

# -----------------------------------------------------------------------------
# Usage:
#   ./scripts/build.sh [<cmd-dir>]
# Examples:
#   ./scripts/build.sh            # builds serve-http by default
#   ./scripts/build.sh serve-grpc # builds serve-grpc
# -----------------------------------------------------------------------------

# Repository root
cd $(git rev-parse --show-toplevel)

# Default to “serve-http” if no argument supplied
cmd_dir=${1:-serve-http}

# Map hyphens to underscores for binary name
#binary_name=${cmd_dir//-/_}
binary_name=main

echo "→ Cleaning"
bazel clean --expunge

echo
echo "→ Regenerating BUILD files"
bazel run //:gazelle

echo
echo "→ Building //cmd/$cmd_dir:$binary_name"
bazel build "//cmd/${cmd_dir}:$binary_name"

echo "✅ $cmd_dir has been built"
