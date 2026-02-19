# Changelog

All notable changes to ProcPipe will be documented in this file.

## [v1.3.0] - 2026-02-19

### Added

- **Dual input racing**: Terminal stdin and Telegram replies are now raced simultaneously — whichever arrives first wins.
- **Dry-run input support**: Input prompts now work in `--dry-run` mode via terminal-only input (no Telegram).
- **Interactive terminal detection**: Log messages adapt based on whether stdin is a real terminal or piped.

### Changed

- Refactored `orchestrator/actions.go` — extracted stdin reading logic into `orchestrator/stdin.go` for cleaner separation.
- "Waiting for input" log now prints **before** input goroutines start, fixing a race condition where fast input could appear before the log.

### Dependencies

- Added `golang.org/x/term` v0.40.0 for terminal detection.
- Updated `golang.org/x/sys` from v0.25.0 to v0.41.0.

## [v1.2.2] - Previous

- Module rename to `github.com/Shivamingale3/ProcPipe`.

## [v1.2.1] - Previous

- Bug fixes and improvements.

## [v1.2.0] - Previous

- Initial Telegram notification support.

## [v1.1.0] - Previous

- Configuration system and input prompt detection.

## [v1.0.0] - Initial Release

- Core process monitoring with PTY support.
