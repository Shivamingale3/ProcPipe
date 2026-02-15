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
URL="https://github.com/Shivamingale3/ProcPipe/releases/latest/download/$BINARY"
DEST="/usr/local/bin/procpipe"

echo "Downloading ProcPipe for $OS/$ARCH..."
if curl --fail --silent --show-error -L -o procpipe "$URL"; then
    chmod +x procpipe
else
    echo "Download failed. Building from source..."
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
