# ProcPipe 🔭

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

### Automated Install (Linux/macOS)

```bash
# If you have the install script locally:
./install.sh
```

### Manual Install

1. Download the binary for your OS
2. Add to your system PATH
3. Run `procpipe version` to verify

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
