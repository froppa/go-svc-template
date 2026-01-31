#!/usr/bin/env bash
set -euo pipefail

# -----------------------------------------------------------------------------
# add-dep.sh — add a Go module dependency via Bazel+rules_go+Gazelle
#
# Usage:
#   ./scripts/add-dep.sh github.com/foo/bar@v1.2.3
# -----------------------------------------------------------------------------

cd $(git rev-parse --show-toplevel)

if [ $# -ne 1 ]; then
  echo "Usage: $0 <module>@<version>" >&2
  exit 1
fi

module="$1"

echo "→ bazel run @rules_go//go:go -- get $module"
bazel run @rules_go//go:go -- get "$module"

echo
echo "→ bazel run //:gazelle update-repos"
bazel run //:gazelle

echo "✅ Added $module"
