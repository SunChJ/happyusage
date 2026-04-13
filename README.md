# happyusage-cli

A tiny cross-platform CLI for checking local AI tool usage.

The project is designed for two audiences:

- **humans** — readable default output with lightweight visual cues
- **agents** — compact text via `--agent` and structured output via `--json`

## Install

### macOS / Linux

```bash
go install github.com/SunChJ/happyusage-cli/cmd/hu@latest
```

### Windows

```powershell
go install github.com/SunChJ/happyusage-cli/cmd/hu@latest
```

### Local build

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

List provider IDs currently available from the local API:

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

```bash
--base-url   custom local API base URL
--timeout    HTTP timeout
--agent      compact agent-friendly text
--json       JSON envelope
```

## Data source

The current version reads from a local usage HTTP API:

- `GET http://127.0.0.1:6736/v1/usage`
- `GET http://127.0.0.1:6736/v1/usage/:providerId`

## Example JSON

```json
{
  "ok": true,
  "source": "local_usage_http_api",
  "base_url": "http://127.0.0.1:6736",
  "checked_at": "2026-04-13T03:00:00Z",
  "provider": {
    "provider_id": "claude",
    "display_name": "Claude",
    "plan": "Pro",
    "progress": [
      {
        "label": "Session",
        "used": 25,
        "limit": 100,
        "remaining": 75,
        "percent_used": 25,
        "unit": "percent"
      }
    ],
    "texts": [
      {
        "label": "Today",
        "value": "$2.10 · 3M tokens"
      }
    ]
  }
}
```

## Development

```bash
go test ./...
go run ./cmd/hu
go run ./cmd/hu usage
go run ./cmd/hu usage claude --agent
go run ./cmd/hu usage claude --json
```

## Roadmap

- native provider fallback probes
- release binaries for macOS, Linux, and Windows
- Homebrew first, then Scoop / winget

## License

MIT
