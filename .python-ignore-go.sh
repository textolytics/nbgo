#!/usr/bin/env bash
if [[ "$1" == *.go ]]; then
  echo "ERROR: Refusing to run Go file with Python: $1"
  exit 1
fi
exec python "$@"
