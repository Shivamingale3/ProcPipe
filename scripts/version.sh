#!/bin/bash

# Check if any tags exist reachable from HEAD
if DESC=$(git describe --tags 2>/dev/null); then
    VERSION="$DESC"
else
    # Fallback: v0.0.<commit_count>-<short_hash>
    COUNT=$(git rev-list --count HEAD 2>/dev/null || echo 0)
    HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
    VERSION="v0.0.${COUNT}-${HASH}"
fi

# Add -dirty if needed
if [[ -n $(git status --porcelain 2>/dev/null) ]]; then
    VERSION="${VERSION}-dirty"
fi

echo "$VERSION"
