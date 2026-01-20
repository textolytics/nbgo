#!/usr/bin/env bash
set -e

echo "Running tests..."

# Run Go tests
echo "Running Go tests..."
go test -v -cover ./...

# Run Robot Framework tests if present
if command -v robot &> /dev/null; then
    echo "Running Robot Framework tests..."
    if [ -d "test" ]; then
        robot --outputdir build/test-results test/
    fi
fi

echo "All tests completed successfully!"
