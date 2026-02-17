# DevClaw Skills Catalog

Official skill collection for [DevClaw](https://github.com/jholhewres/devclaw) — the AI agent for developers.

Skills teach the agent how to use CLI tools, APIs, and workflows by injecting contextual instructions into the LLM prompt. They do **not** register new tools — they leverage existing native tools (`bash`, `ssh`, `scp`, `read_file`, `write_file`, `memory_save`, `cron_add`, etc.).

> **Starter Pack** (docker, git, npm, go, ssh, github) is embedded in the DevClaw binary and pre-selected during setup. Everything else lives here.

## Skills (80)

### Development (25)

| Skill | Description | Requires |
|-------|-------------|----------|
| **github** | Issues, PRs, releases, CI via `gh` CLI | `gh` |
| **docker** | Containers, images, volumes, networks | `docker` |
| **docker-compose** | Multi-service orchestration, profiles, scaling | `docker` |
| **kubernetes** | Pods, services, deployments via kubectl | `kubectl` |
| **terraform** | Infrastructure as Code — plan, apply, state | `terraform` |
| **aws-cli** | S3, EC2, Lambda, CloudWatch, RDS, ECS | `aws` |
| **gcloud** | Compute, Cloud Run, GCS, BigQuery | `gcloud` |
| **azure-cli** | VMs, App Service, Blob Storage, SQL | `az` |
| **cargo** | Rust — build, test, clippy, publish | `cargo` |
| **pip-poetry** | Python packages — pip, poetry, pipenv | `python3` |
| **tmux** | Terminal multiplexing sessions | `tmux` |
| **claude-code** | AI coding via Claude Code CLI | `claude` |
| **image-gen** | Generate images from text (DALL-E/GPT) | — |
| **project-manager** | Register and manage dev projects | — |
| **filesystem** | Advanced file ops, search, rsync | — |
| **browser-automation** | Playwright/Puppeteer automation | — |
| **lighthouse** | Web performance audit (Lighthouse) | `npx` |
| **sec-deps** | Security audit across languages | — |
| **a11y-axe** | Accessibility audit (axe-core) | `npx` |
| **bundler-analyze** | JS bundle size analysis (Vite/Webpack) | — |
| **color-contrast** | WCAG color contrast ratios | — |
| **dora-metrics** | DORA engineering metrics from git | — |
| **changelog-gen** | CHANGELOG from conventional commits | — |
| **release-notes** | Release notes from issues + commits | — |
| **database** | PostgreSQL, MySQL, MongoDB, Redis, SQLite | — |

### Infrastructure (7)

| Skill | Description | Requires |
|-------|-------------|----------|
| **nginx** | Config, reverse proxy, SSL | `nginx` |
| **redis-cli** | Cache, queues, in-memory data | `redis-cli` |
| **systemd** | systemctl + journalctl services | — |
| **pm2** | Process management | `pm2` |
| **railway-deploy** | Railway projects and services | `railway` |
| **backup** | Backup scripts — rsync, tar, dumps | — |
| **health** | Service health checks, uptime, SSL | — |

### Productivity (17)

| Skill | Description | Requires |
|-------|-------------|----------|
| **obsidian** | Obsidian vaults — daily notes, wikilinks, tags | — |
| **journal** | Daily journaling with reflections | — |
| **pomodoro** | Focus sessions, breaks, tracking | — |
| **time-tracker** | Track time per task/project | — |
| **habits** | Daily habit tracking with streaks | — |
| **meeting-notes** | Structured notes + action items | — |
| **bookmarks** | Save and organize links | — |
| **standup-summary** | Daily standup from git + issues | — |
| **sprint-report** | Sprint reports from git + tracker | — |
| **notes** | Quick notes in local markdown | — |
| **timer** | Quick timers and alarms | — |
| **reminders** | Scheduled reminders | — |
| **translate** | Translate between any languages | — |
| **calendar** | Google Calendar | `gcalcli` or `gog` |
| **gog** | Google Workspace (Gmail, Calendar, Drive) | `gog` |
| **todoist** | Todoist task management | `TODOIST_API_TOKEN` |
| **notion** | Notion pages, databases, blocks | `NOTION_API_KEY` |

### Data & Web (6)

| Skill | Description | Requires |
|-------|-------------|----------|
| **web-search** | Brave Search / DuckDuckGo | `BRAVE_API_KEY` (optional) |
| **web-fetch** | Fetch readable content from URLs | — |
| **web-scraper** | Scraping with selectors, Playwright | — |
| **summarize** | Summarize URLs, videos, podcasts | — |
| **ocr** | Text extraction from images/PDFs | `tesseract` |
| **stock-prices** | Stocks, crypto, market data | — |

### Communication (5)

| Skill | Description | Requires |
|-------|-------------|----------|
| **gmail** | Gmail — list, read, send, labels | — |
| **email** | Send via Resend, SendGrid, Mailgun | API key |
| **slack** | Messages, channels, threads | `SLACK_BOT_TOKEN` |
| **discord** | Webhooks and bot API | — |
| **telegram** | Bot API — messages, files | `TELEGRAM_BOT_TOKEN` |

### Media & Files (8)

| Skill | Description | Requires |
|-------|-------------|----------|
| **ffmpeg** | Video/audio processing | `ffmpeg` |
| **imagemagick** | Image processing | `magick` or `convert` |
| **pandoc** | Document format conversion | `pandoc` |
| **image-tools** | Resize, convert, compress images | — |
| **video-tools** | Video processing via ffmpeg | `ffmpeg` |
| **audio-tts** | Text-to-speech (OpenAI, ElevenLabs) | API key |
| **pdf-tools** | PDF merge, split, compress, extract | — |
| **screenshot** | Web page screenshots | — |

### Builtin (4)

| Skill | Description |
|-------|-------------|
| **weather** | Forecasts via wttr.in / Open-Meteo |
| **calculator** | Math calculations and conversions |
| **location** | Geocoding, IP lookup, timezone |
| **qr-code** | Generate and decode QR codes |

### Integrations (8)

| Skill | Description | Requires |
|-------|-------------|----------|
| **jira** | Issues, sprints, transitions, worklogs | API credentials |
| **linear** | Issues, teams, projects | `LINEAR_API_KEY` |
| **trello** | Boards, lists, cards | API credentials |
| **asana** | Tasks, projects, workspaces | `ASANA_ACCESS_TOKEN` |
| **clickup** | Tasks, comments, time tracking | `CLICKUP_API_TOKEN` |
| **airtable** | Records, tables, batch operations | `AIRTABLE_API_KEY` |
| **payments** | Stripe, PayPal, Lightning | API credentials |
| **crypto-wallet** | Crypto wallets, balances, market data | — |

## Installation

### Via Setup Wizard

During `devclaw setup`, skills from this catalog are shown alongside the embedded Starter Pack. Select what you need.

### Via CLI

```bash
devclaw skill install github:jholhewres/devclaw-skills/skills/<name>
```

### Manual

```bash
git clone https://github.com/jholhewres/devclaw-skills.git
cp -r devclaw-skills/skills/<name> ./skills/
```

## Creating a Skill

Each skill is a directory with a `SKILL.md` file:

```
skills/my-skill/
└── SKILL.md
```

### SKILL.md Format

```markdown
---
name: my-skill
description: "What this skill teaches the agent"
---
# My Skill

Instructions for the agent. Tell it which **native tools** to use
(bash, ssh, read_file, etc.) and provide concrete command examples.
```

### Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `name` | yes | Unique identifier (`lowercase-dashes`) |
| `description` | yes | Brief description |
| `version` | no | Semver version |
| `category` | no | `builtin`, `data`, `development`, `productivity`, `infra`, `media`, `communication` |
| `tags` | no | Tags for search and filtering |
| `requires.bins` | no | Required CLI tools (all must be present) |
| `requires.any_bins` | no | CLI tools (at least one required) |
| `requires.env` | no | Required environment variables |
| `requires.any_env` | no | Env vars (at least one required) |

### Key Principle

Skills do **not** register new tools. They teach the LLM to use existing native tools:

| Native Tool | What it does |
|-------------|-------------|
| `bash` | Execute any shell command |
| `ssh` / `scp` | Remote commands and file transfer |
| `read_file` / `write_file` / `edit_file` | File operations |
| `web_search` / `web_fetch` | Web search and fetch |
| `memory_save` / `memory_search` | Long-term memory |
| `vault_get` / `vault_save` | Encrypted secrets |
| `cron_add` / `cron_list` | Scheduling |

## Contributing

1. Fork the repository
2. Create your skill in `skills/<name>/SKILL.md`
3. Add an entry to `index.yaml`
4. Submit a pull request

## License

[MIT](LICENSE)
