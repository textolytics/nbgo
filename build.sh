#!/usr/bin/env bash
set -e

# Build configuration
VERSION="1.0.0"
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

# Output directory
OUTPUT_DIR="build/output"
mkdir -p "$OUTPUT_DIR"

echo "Building NBGO v$VERSION (commit: $COMMIT)"
echo "Build time: $BUILD_TIME"

# Build for multiple platforms
TARGETS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

for target in "${TARGETS[@]}"; do
    IFS='/' read -r OS ARCH <<< "$target"
    echo "Building for $OS/$ARCH..."

    OUTPUT_FILE="$OUTPUT_DIR/nbgo-$VERSION-$OS-$ARCH"
    [ "$OS" = "windows" ] && OUTPUT_FILE="$OUTPUT_FILE.exe"

    GOOS=$OS GOARCH=$ARCH CGO_ENABLED=0 go build \
        -ldflags "-X main.Version=$VERSION -X main.Commit=$COMMIT -X main.BuildTime=$BUILD_TIME" \
        -o "$OUTPUT_FILE" \
        .

    echo "✓ Built $OUTPUT_FILE"
done

echo "Build complete! Outputs in $OUTPUT_DIR"
