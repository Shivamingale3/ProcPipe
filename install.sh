#!/bin/bash
set -e

# Detect OS and Arch
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

BINARY="procpipe-${OS}-${ARCH}"
URL="https://github.com/yourusername/procpipe/releases/latest/download/$BINARY"
DEST="/usr/local/bin/procpipe"

echo "Downloading ProcPipe for $OS/$ARCH..."
# In real distribution, download from github releases
# curl -L -o procpipe $URL

# For now, assume building from source or using dist folder if present
if [ -f "dist/$BINARY" ]; then
    cp "dist/$BINARY" procpipe
elif [ -f "dist/procpipe" ]; then
    cp "dist/procpipe" procpipe
else
    echo "Building from source..."
    make build
    cp "dist/procpipe" procpipe
fi

echo "Installing to $DEST..."
if [ -w "$(dirname "$DEST")" ]; then
    mv procpipe "$DEST"
else
    sudo mv procpipe "$DEST"
fi
sudo chmod +x "$DEST"

echo "✅ Installed successfully!"
echo "Run 'procpipe config' to get started."
