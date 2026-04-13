# happyusage-cli

A tiny cross-platform CLI for checking local AI tool usage.

The project is designed for two audiences:

- **humans** — readable default output with lightweight visual cues
- **agents** — compact text via `--agent` and structured output via `--json`

## Install

### Quick install (macOS / Linux)

```bash
curl -fsSL https://raw.githubusercontent.com/SunChJ/happyusage-cli/main/scripts/install.sh | sh
```

Options:

```bash
# Install a specific version
VERSION=v0.1.0 curl -fsSL https://raw.githubusercontent.com/SunChJ/happyusage-cli/main/scripts/install.sh | sh

# Custom install directory
BIN_DIR=/usr/local/bin curl -fsSL https://raw.githubusercontent.com/SunChJ/happyusage-cli/main/scripts/install.sh | sh
```

### Go

```bash
go install github.com/SunChJ/happyusage-cli/cmd/hu@latest
```

### Build from source

```bash
go build -o bin/hu ./cmd/hu
```

## Quick start

Run the command with no arguments to see the built-in help:

```bash
hu
```

## Commands

### Help

```bash
hu help
hu help usage
```

### Version

```bash
hu version
```

### Usage

Show all configured providers in the default human-friendly view:

```bash
hu usage
```

List provider IDs currently available from the built-in provider probes:

```bash
hu usage list
```

Show a single provider:

```bash
hu usage claude
hu usage codex
```

Show compact agent-friendly text:

```bash
hu usage claude --agent
```

Show JSON output:

```bash
hu usage claude --json
hu usage --json
hu usage list --json
```

## Flags

These flags currently apply to `hu usage`:

```bash
--agent   compact agent-friendly text
--json    JSON envelope
```

## Data collection model

`hu` does not depend on a running desktop app or any external local usage API.

The current implementation ships with built-in native provider collection scripts,
primarily tuned for macOS first. The bundled probe set currently covers:

- Claude
- Codex
- Cursor
- Copilot
- Gemini
- Windsurf

## Example JSON

```json
{
  "ok": true,
  "source": "native_provider_scripts",
  "checked_at": "2026-04-13T03:00:00Z",
  "provider": {
    "provider": "claude",
    "ok": true,
    "plan": "Pro",
    "quotas": [
      {
        "name": "session",
        "period": "5h",
        "used_pct": 25,
        "left_pct": 75
      }
    ]
  }
}
```

## Platform status

- **macOS** — first-class target for the current native probes
- **Linux / Windows** — installation is supported, but provider collection is not yet at feature parity

## Development

```bash
go test ./...
go run ./cmd/hu
go run ./cmd/hu usage
go run ./cmd/hu usage claude --agent
go run ./cmd/hu usage claude --json
```

## Roadmap

- replace shell-heavy probes with native Go implementations where it makes sense
- expand Linux and Windows collection support
- release binaries for macOS, Linux, and Windows
- Homebrew first, then Scoop / winget

## License

MIT
