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

---

## 📥 Installation

Choose the method that works best for your operating system.

### Linux / macOS

**Automatic Install** (Recommended):
Run this one-line command to download the latest binary and install it to `/usr/local/bin`:

```bash
curl -sL https://raw.githubusercontent.com/Shivamingale3/ProcPipe/main/install.sh | bash
```

**Go Install**:
If you have Go installed, you can build directly from source:

```bash
go install github.com/Shivamingale3/ProcPipe@latest
```

### Windows

**PowerShell** (Recommended):
Run this command in PowerShell to download and install:

```powershell
iwr -useb https://raw.githubusercontent.com/Shivamingale3/ProcPipe/main/install.ps1 | iex
```

**Manual Install**:

1. Go to [Releases](https://github.com/Shivamingale3/ProcPipe/releases).
2. Download the `procpipe-windows-amd64.exe` binary.
3. Rename it to `procpipe.exe` and place it in a folder in your `PATH`.
4. Run `procpipe install` to verify.

### Uninstall

To remove ProcPipe from your system:

```bash
# Linux/macOS
procpipe uninstall
# OR
./uninstall.sh

# Windows
procpipe uninstall
```

---

## 🤖 Telegram Bot Setup

To receive notifications, you need to create a Telegram bot and get your Chat ID.

### 1. Create a Bot

1. Open Telegram and search for **[@BotFather](https://t.me/BotFather)**.
2. Send the command `/newbot`.
3. Follow the instructions to name your bot (e.g., `MyProcPipeBot`).
4. **Copy the HTTP API Token** provided (it looks like `123456:ABC-DEF1234gh...`).

### 2. Get Your Chat ID

1. Search for **[@userinfobot](https://t.me/userinfobot)** on Telegram.
2. Send any message (like `/start`).
3. It will reply with your `Id` (e.g., `987654321`). **Copy this number**.

### 3. Configure ProcPipe

Run the interactive configuration wizard:

```bash
procpipe config
```

- Paste your **Bot Token**.
- Paste your **Chat ID**.
- The tool will send a test message to verify the connection.

> **Note:** Configuration is saved to `~/.config/procpipe/config.yaml` (or equivalent on Windows). You can view the file path with `procpipe config path`.

---

## 🚀 Usage

The basic syntax is `procpipe run -- <command>`. The `--` separator is required to distinguish ProcPipe flags from your command's flags.

### Basic Examples

**Run a simple command:**

```bash
procpipe run -- sudo apt update
```

**Run a build process:**

```bash
procpipe run -- make build
```

**Run a command with arguments:**

```bash
procpipe run -- python3 script.py --verbose --input file.txt
```

**Run chained commands:**
To run complex pipelines or chained commands, wrap them in quotes:

```bash
procpipe run -- "npm install && npm run build"
```

### Options

**Dry Run (Test without Telegram):**
Use `--dry-run` or just run without configuring Telegram to see output in the terminal only.

```bash
procpipe run --dry-run -- echo "hello world"
```

**Specify a different config file:**

```bash
procpipe run --config ./alternate-config.yaml -- ./deploy.sh
```

---

## ⚙️ Configuration Reference

You can manage configuration via the CLI or by editing the config file directly.

**CLI Commands:**

```bash
procpipe config            # Start interactive wizard
procpipe config show       # Show current valid configuration
procpipe config path       # Print path to config file
procpipe config test       # Send a test message to defined Telegram chat
```

**Config File Structure:**

```yaml
telegram:
  bot_token: "123456:ABC-DEF..."
  chat_id: 987654321
monitor:
  log_tail_lines: 20
  input_patterns: []
```

---

## 🔍 How Detection Works

| Feature                | Mechanism                               | CPU Usage                        |
| :--------------------- | :-------------------------------------- | :------------------------------- |
| **Process completion** | PTY read returns EOF -> check exit code | Zero (blocks on file descriptor) |
| **Input prompts**      | Regex matching on each output chunk     | Negligible                       |
| **Telegram replies**   | Long polling with 60s server timeout    | Zero (HTTP blocks)               |

**Input Patterns**:
ProcPipe automatically detects common prompts like:

- `[Y/n]`, `[yes/no]`
- `password:`, `passphrase:`
- `Press Enter to continue`

---

## 🛠 Project Structure

```
├── cmd/                 # Cobra commands (run, config, install)
├── config/              # Configuration loading & validation logic
├── process/             # Low-level PTY process spawning
├── monitor/             # Output monitoring & pattern matching
├── notify/              # Notification system interfaces
├── telegram/            # Telegram API client & long-polling
├── orchestrator/        # Main logic connecting process <-> telegram
├── logger/              # Terminal logging utilities
└── version/             # Versioning info
```

## 📢 Changelog

### v1.2.1 (2026-02-15)

- **Uninstall Support**: Added `procpipe uninstall` and scripts.
- **Code Quality**: Achieved **A+ (100%)** on Go Report Card.
- **CI/CD**: Automated release pipeline via GitHub Actions.
- **Cross-Platform**: Verified on Linux, macOS, and Windows.

## 🤝 Contributing

Contributions are welcome!

1. Fork the Project
2. Create your Feature Branch
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 💖 Support

If ProcPipe saves you time, please consider:

- ⭐ Staring the repo on [GitHub](https://github.com/Shivamingale3/ProcPipe)
- 🐛 Reporting issues
