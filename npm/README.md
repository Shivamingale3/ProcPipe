# ProcPipe 🔭

**Terminal process watcher with Telegram notifications.**

Spawn any long-running command — builds, updates, deployments — and walk away. ProcPipe watches it with near-zero resource usage and messages you on Telegram when it completes (with logs) or when it needs input.

## Install

```bash
npm install -g procpipe
```

## Quick Start

```bash
# Configure Telegram notifications
procpipe config

# Watch a command
procpipe run -- sudo apt update

# Watch a build
procpipe run -- make build
```

## Features

- 🚀 **PTY-based** — spawns commands in a real pseudo-terminal
- 📡 **Telegram integration** — sends rich notifications with logs
- ⚠️ **Input detection** — detects prompts like `[Y/n]`, `password:`, forwards your reply
- 🪶 **Zero polling** — uses blocking I/O, sleeps until something happens
- 📦 **Single binary** — static executable, no dependencies, cross-platform
- 🔧 **Interactive Setup** — built-in config wizard

## Supported Platforms

| OS      | Architecture               |
| ------- | -------------------------- |
| Linux   | x64, arm64                 |
| macOS   | x64, arm64 (Apple Silicon) |
| Windows | x64                        |

## Documentation

Full documentation, Telegram setup guide, and more at the [GitHub repository](https://github.com/Shivamingale3/ProcPipe).

## License

MIT
