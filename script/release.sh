#!/usr/bin/env bash

set -euo pipefail

targets=(
  darwin-amd64
  darwin-arm64
  linux-amd64
  windows-amd64
)

# Create zipped binaries.
# chdir into dist so zip files sit at root of zip instead of in 'dist/pggen'.
pushd dist >/dev/null
for target in "${targets[@]}"; do
  binary="bdix-tester-${target}"
  if [[ "$binary" == *windows* ]]; then
    binary+='.exe'
  fi
  echo -n "zipping ${binary} ... "
  zip --quiet -9 "bdix-tester-${target}.zip" "${binary}"
  echo "done"
done
popd >/dev/null