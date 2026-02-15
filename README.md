# ProcPipe 🔭

![ProcPipe Banner](version/banner.png)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/Shivamingale3/ProcPipe)](https://goreportcard.com/report/github.com/Shivamingale3/ProcPipe)
[![GitHub Release](https://img.shields.io/github/v/release/Shivamingale3/ProcPipe)](https://github.com/Shivamingale3/ProcPipe/releases)

**Terminal process watcher with Telegram notifications.**

Spawn any long-running command — builds, updates, deployments — and walk away. ProcPipe watches it with near-zero resource usage and messages you on Telegram when it completes (with logs) or when it needs input.

## Features

- 🚀 **PTY-based** — spawns commands in a real pseudo-terminal
- 📡 **Telegram integration** — sends rich notifications with logs
- ⚠️ **Input detection** — detects prompts like `[Y/n]`, `password:`, forwards your reply
- 🪶 **Zero polling** — uses blocking I/O, sleeps until something happens
- 📦 **Single binary** — static executable, no dependencies, cross-platform
- 🔧 **Interactive Setup** — built-in config wizard

## Installation

## Installation

### Linux / macOS

Run this one-line command to download and install automatically:

```bash
curl -sL https://raw.githubusercontent.com/Shivamingale3/ProcPipe/main/install.sh | bash
```

### Windows (PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/Shivamingale3/ProcPipe/main/install.ps1 | iex
```

### Manual Install

1. Go to [Releases](https://github.com/Shivamingale3/ProcPipe/releases)
2. Download the binary for your OS
3. Run: `./procpipe install`

### Uninstall

```bash
procpipe uninstall
# OR
./uninstall.sh
```

## Quick Start

### 1. Setup

Run the interactive wizard to set up your Telegram bot:

```bash
procpipe config
```

This will guide you through creating a bot and verifying the connection.

### 2. Run Commands

Prefix any command with `procpipe`:

```bash
# Basic usage (defaults to 'run')
procpipe -- sudo apt update

# Explicit run command
procpipe run -- make build

# Test locally (no Telegram needed)
procpipe run --dry-run -- echo "hello world"
```

## CLI Reference

```bash
procpipe run -- <cmd>      # Run command (default)
procpipe config            # Interactive config wizard
procpipe config show       # Show current config
procpipe config test       # Test Telegram connection
procpipe config path       # Print config file location
procpipe install           # Install to system PATH
procpipe uninstall         # Remove from system PATH
procpipe version           # Show version info
```

## How Detection Works

| What                   | How                                    | CPU Cost            |
| ---------------------- | -------------------------------------- | ------------------- |
| **Process completion** | PTY read returns EOF → check exit code | Zero (blocks on fd) |
| **Input prompts**      | Regex matching on each output chunk    | Negligible          |
| **Telegram replies**   | Long polling with 60s server timeout   | Zero (HTTP blocks)  |

**No timers, no polling loops.** The app sleeps between events.

### Built-in Input Patterns

Automatically detected prompts:

- `[Y/n]`, `[yes/no]`, `(y/n)`
- `password:`, `passphrase:`
- `Enter ...:`, `Continue?`
- `Press Enter`, `Type X to confirm`
- `[sudo] password`, `Do you want to`, `Are you sure`

## Cross-Platform Builds

```bash
make build-all    # Linux, macOS, Windows (amd64 + arm64)
make build        # Current platform only
```

## Project Structure

```
├── cmd/                 # Cobra subcommands (run, config, install, etc.)
├── config/              # YAML config + loader
├── process/             # PTY process spawner
├── monitor/             # Output reader + pattern matcher
├── notify/              # Notifier interface
├── telegram/            # Telegram client + poller
├── orchestrator/        # Main event loop
├── logger/              # Terminal logger
└── version/             # Build info
```

## 🤝 Contributing

Contributions are welcome! Whether it's reporting a bug, suggesting a feature, or writing code, I'd love to see it.

1.  **Fork** the repository
2.  Create your feature branch (`git checkout -b feature/amazing-feature`)
3.  Commit your changes (`git commit -m 'Add some amazing feature'`)
4.  Push to the branch (`git push origin feature/amazing-feature`)
5.  Open a **Pull Request**

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 💖 Support

If ProcPipe saves you time, please consider supporting the project:

- ⭐ **Star** on [GitHub](https://github.com/Shivamingale3/ProcPipe)
- 🐛 **Report** issues / **Fork** and contribute
