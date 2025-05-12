#!/usr/bin/env bash
set -euo pipefail

# -----------------------------------------------------------------------------
# Usage:
#   ./scripts/run.sh [--direct] <cmd-dir> [-- <args>]
#
# Examples:
#   # default: bazel run (no real signal forwarding, but simpler)
#   ./scripts/run.sh serve-http -- --config=configs/config.yml
#
#   # direct binary build+exec (with SIGINT→SIGINT forwarding)
#   ./scripts/run.sh --direct serve-http -- --config=configs/config.yml
# -----------------------------------------------------------------------------

cd $(git rev-parse --show-toplevel)

direct_mode=0
if [[ "${1-}" == "--direct" ]]; then
  direct_mode=1
  shift
fi

if [ $# -lt 1 ]; then
  echo "Usage: $0 [--direct] <cmd-dir> [-- <args>]" >&2
  exit 1
fi

cmd_dir=$1; shift
binary_name=main
pkg="//cmd/${cmd_dir}"

if [ $direct_mode -eq 0 ]; then
  target="${pkg}:${binary_name}"
  echo "→ bazel run ${target} -- $*"
  bazel run "${target}" -- "$@"
  exit
fi

# direct mode: build & exec binary for true signal forwarding
target="${pkg}:${binary_name}"
echo "→ bazel build ${target}"
bazel build "${target}"

bin_dir="$(bazel info bazel-bin)/cmd/${cmd_dir}/${binary_name}_"
exe="${bin_dir}/${binary_name}"

if [ ! -x "$exe" ]; then
  echo "ERROR: binary not found at $exe" >&2
  echo "Contents of $bin_dir:" >&2
  ls -lha "$bin_dir" >&2
  exit 1
fi

echo
echo "→ exec $exe $*"
"$exe" "$@" &
child=$!

# forward CTRL+C/SIGTERM
trap 'echo; echo "→ Caught signal, forwarding to service…"; kill -SIGINT "$child"' INT TERM
wait "$child"
