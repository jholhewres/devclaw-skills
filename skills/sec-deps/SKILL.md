---
name: sec-deps
description: "Security audit for dependencies across languages"
metadata: {"openclaw":{"always":false,"emoji":"🛡️"}}
---
# Security Dependencies Audit

Verify vulnerabilities in project dependencies.

## Setup

**Audit tools** (install per language as needed):
- **Node.js**: `brew install node` (macOS) or `sudo apt install nodejs npm` (Ubuntu) — `npm audit` included
- **Go**: `go install golang.org/x/vuln/cmd/govulncheck@latest`
- **Python**: `pip install pip-audit` (or `pip install safety` for alternative)
- **Rust**: `cargo install cargo-audit`

## Node.js (npm)

```bash
# Audit
npm audit
npm audit --json | jq '.metadata.vulnerabilities'

# Fix automatically
npm audit fix
npm audit fix --force   # may update major versions

# Check outdated
npm outdated
```

## Go

```bash
# Install govulncheck (if needed)
go install golang.org/x/vuln/cmd/govulncheck@latest

# Check vulnerabilities
govulncheck ./...

# Check specific modules
govulncheck -show verbose ./...
```

## Python (pip)

```bash
# Install pip-audit (if needed)
pip install pip-audit

# Audit
pip-audit
pip-audit --format=json
pip-audit -r requirements.txt

# Safety (alternative)
pip install safety
safety check
safety check --json
```

## Rust (cargo)

```bash
# Install cargo-audit
cargo install cargo-audit

# Audit
cargo audit
cargo audit --json
```

## Multi-language

```bash
# Detect and audit automatically
audit_all() {
  echo "=== Checking for vulnerabilities ==="
  
  if [ -f "package.json" ]; then
    echo "--- Node.js ---"
    npm audit 2>/dev/null || echo "npm audit failed"
  fi
  
  if [ -f "go.mod" ]; then
    echo "--- Go ---"
    govulncheck ./... 2>/dev/null || echo "govulncheck not installed"
  fi
  
  if [ -f "requirements.txt" ]; then
    echo "--- Python ---"
    pip-audit -r requirements.txt 2>/dev/null || echo "pip-audit not installed"
  fi
  
  if [ -f "Cargo.toml" ]; then
    echo "--- Rust ---"
    cargo audit 2>/dev/null || echo "cargo-audit not installed"
  fi
}
audit_all
```

## Severity

| Level | Action |
|-------|--------|
| Critical | Fix immediately |
| High | Fix within 24h |
| Moderate | Fix in next sprint |
| Low | Evaluate and plan |

## Tips

- Run audit in CI/CD to block PRs with critical vulnerabilities
- Use `npm audit --production` to ignore devDependencies
- Keep lockfiles up to date
- Configure Dependabot or Renovate for automatic PRs
