#!/usr/bin/env bash

# Exit on error.
set -e

# Navigate to the root of the project.
cd "$(dirname "$0")/.."

# Run the tests.
go test "./..." -v -coverprofile="coverage.out" -covermode=count -parallel="$(nproc)"

# Display the coverage statistics and generate an HTML report.
go tool cover -func coverage.out
# go tool cover -html=coverage.out
