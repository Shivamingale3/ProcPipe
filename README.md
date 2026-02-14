# ProcPipe 🔭

**Terminal process watcher with Telegram notifications.**

Spawn any long-running command — builds, updates, deployments — and walk away. ProcPipe watches it with near-zero resource usage and messages you on Telegram when it completes (with logs) or when it needs input.

## Features

- 🚀 **PTY-based** — spawns commands in a real pseudo-terminal
- 📡 **Telegram integration** — sends rich notifications with logs
- ⚠️ **Input detection** — detects prompts like `[Y/n]`, `password:`, forwards your reply
- 🪶 **Zero polling** — uses blocking I/O, sleeps until something happens
- 📦 **Single binary** — static executable, no dependencies, cross-platform

## Quick Start

```bash
# Build
make build

# Test locally (no Telegram needed)
./dist/procpipe --dry-run -- echo "hello world"

# With Telegram
./dist/procpipe --token "YOUR_BOT_TOKEN" --chat 123456789 -- apt update && apt upgrade
```

## Setup Telegram Credentials

### Step 1: Create a Bot

1. Open Telegram and search for **@BotFather**
2. Send `/newbot`
3. Choose a name (e.g., "ProcPipe Bot")
4. Choose a username (e.g., "my_procpipe_bot")
5. BotFather will give you a **Bot Token** like: `7123456789:AAF1k2j3h4g5f6d7s8a9`
6. Save this token

### Step 2: Get Your Chat ID

1. Search for **@userinfobot** on Telegram
2. Send it any message
3. It will reply with your **Chat ID** (a number like `987654321`)
4. Save this number

### Step 3: Start Your Bot

1. Search for your bot by its username on Telegram
2. Press **Start** — this is required before the bot can send you messages

### Step 4: Configure ProcPipe

Create `~/.procpipe.yaml`:

```yaml
telegram:
  bot_token: "7123456789:AAF1k2j3h4g5f6d7s8a9"
  chat_id: 987654321

monitor:
  log_tail_lines: 50
```

Or use CLI flags:

```bash
./dist/procpipe --token "TOKEN" --chat 123456 -- your-command
```

## Usage

```bash
# Basic usage
procpipe [flags] -- <command> [args...]

# Flags
--token     Telegram bot token (overrides config)
--chat      Telegram chat ID (overrides config)
--config    Path to config file (default: ~/.procpipe.yaml)
--dry-run   Print notifications to stdout instead of Telegram

# Examples
procpipe -- make build
procpipe -- npm run build
procpipe -- sudo apt update && sudo apt upgrade
procpipe -- docker build -t myapp .
procpipe --dry-run -- sleep 5
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

Add custom patterns in config:

```yaml
monitor:
  input_patterns:
    - "my custom prompt:"
    - "(?i)confirm.*:"
```

## Cross-Platform Builds

```bash
make build-all    # Linux, macOS, Windows (amd64 + arm64)
make build        # Current platform only
```

## Project Structure

```
├── main.go              # Entry point
├── config/              # YAML config + CLI flags + validation
├── process/             # PTY process spawner + lifecycle
├── monitor/             # Output reader + ring buffer + pattern matcher
├── notify/              # Notifier interface + dry-run implementation
├── telegram/            # Telegram Bot API client + formatters
├── orchestrator/        # Main event loop + handlers
├── logger/              # Colored terminal logger
├── Makefile             # Build targets
└── procpipe.example.yaml # Example config
```
