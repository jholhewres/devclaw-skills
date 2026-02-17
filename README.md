# DevClaw Skills Catalog

Official skill collection for [DevClaw Copilot](https://github.com/jholhewres/devclaw).

Skills extend the agent's capabilities by providing contextual instructions on how to use
CLI tools, APIs, and workflows. Each skill is a `SKILL.md` file with YAML frontmatter
(ClawdHub/OpenClaw format).

## Available Skills

### Builtin (no API key required)

| Skill | Description |
|-------|-------------|
| **weather** | Weather forecasts via wttr.in / Open-Meteo |
| **calculator** | Math calculations and unit conversions |
| **timer** | Timers, alarms, and Pomodoro sessions |
| **reminders** | Time-based reminders with scheduling |
| **notes** | Quick notes, lists, and ideas (local markdown) |
| **translate** | Translate text between any languages |

### Data & Web

| Skill | Description | Requirements |
|-------|-------------|-------------|
| **web-search** | Web search via Brave API or DuckDuckGo | `BRAVE_API_KEY` (optional) |
| **web-fetch** | Fetch and extract readable content from URLs | none |
| **summarize** | Summarize URLs, videos, articles | `yt-dlp` (optional) |

### Development

| Skill | Description | Requirements |
|-------|-------------|-------------|
| **github** | Issues, PRs, releases, CI via `gh` CLI | `gh` CLI |

### Productivity & Integrations

| Skill | Description | Requirements |
|-------|-------------|-------------|
| **calendar** | Google Calendar integration | `gcalcli` or `gog` |
| **gog** | Gmail, Calendar, Drive via `gog` CLI | `gog` CLI |
| **notion** | Notion pages, databases, and blocks | `NOTION_API_KEY` |
| **trello** | Trello boards, lists, and cards | `TRELLO_API_KEY` + `TRELLO_TOKEN` |
| **todoist** | Todoist task management | `TODOIST_API_TOKEN` |

## Installation

### Via DevClaw CLI

```bash
# Install a specific skill
copilot skill install github:jholhewres/devclaw-skills/skills/web-search

# Or clone and install all
git clone git@github.com:jholhewres/devclaw-skills.git
cp -r devclaw-skills/skills/* ./skills/
```

### During Setup

When running `copilot setup`, you'll be offered a selection of default skills to install.

### Manual

Copy the skill directory to your DevClaw `skills/` folder:

```bash
cp -r /path/to/devclaw-skills/skills/weather ./skills/weather
```

## Skill Format

Each skill is a directory containing a `SKILL.md` file:

```
skills/
├── web-search/
│   └── SKILL.md
├── weather/
│   └── SKILL.md
└── github/
    └── SKILL.md
```

### SKILL.md Structure

```markdown
---
name: my-skill
version: 0.1.0
author: your-name
description: "What this skill does"
category: data
tags: [tag1, tag2]
requires:
  bins: [required-cli-tool]
  any_env: [OPTIONAL_API_KEY]
---
# Skill Name

Instructions for the agent on how to use this skill.
Include bash examples, tips, and trigger phrases.
```

### Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `name` | yes | Unique identifier (`lowercase-with-dashes`) |
| `version` | no | Semver version |
| `author` | no | Author name |
| `description` | yes | Brief description |
| `category` | no | One of: `builtin`, `data`, `development`, `productivity`, `communication`, `automation` |
| `tags` | no | Tags for search and filtering |
| `requires.bins` | no | Required CLI tools (all must be present) |
| `requires.any_bins` | no | CLI tools (at least one must be present) |
| `requires.env` | no | Required environment variables |
| `requires.any_env` | no | Environment variables (at least one must be set) |

## Creating a Skill

1. Create a directory: `mkdir skills/my-skill`
2. Create `SKILL.md` with frontmatter and instructions
3. Test: `copilot skill info my-skill`
4. The skill is loaded automatically on next `copilot serve` or `copilot chat`

You can also ask DevClaw to create skills via chat:

> "Create a skill that checks Docker container status"

## Contributing

1. Fork the repository
2. Create your skill in `skills/<name>/SKILL.md`
3. Test with `copilot skill list` and `copilot skill info <name>`
4. Submit a pull request

## License

MIT
