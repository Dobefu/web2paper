#!/usr/bin/env bash

# Exit on error.
set -e

# Navigate to the root of the project.
cd "$(dirname "$0")/.."

# Run the benchmarks.
go test -bench=. "./..." "$@" -benchmem -benchtime 100x -run notest -cpu 4 -count 1
