---
name: sec-deps
description: "Auditoria de segurança de dependências: npm audit, govulncheck, pip audit"
metadata: {"openclaw":{"always":false,"emoji":"🛡️"}}
---
# Security Dependencies Audit

Verificar vulnerabilidades em dependências de projetos.

## Node.js (npm)

```bash
# Auditoria
npm audit
npm audit --json | jq '.metadata.vulnerabilities'

# Corrigir automaticamente
npm audit fix
npm audit fix --force   # pode atualizar major versions

# Verificar outdated
npm outdated
```

## Go

```bash
# Instalar govulncheck (se necessário)
go install golang.org/x/vuln/cmd/govulncheck@latest

# Verificar vulnerabilidades
govulncheck ./...

# Verificar módulos específicos
govulncheck -show verbose ./...
```

## Python (pip)

```bash
# Instalar pip-audit (se necessário)
pip install pip-audit

# Auditoria
pip-audit
pip-audit --format=json
pip-audit -r requirements.txt

# Safety (alternativa)
pip install safety
safety check
safety check --json
```

## Rust (cargo)

```bash
# Instalar cargo-audit
cargo install cargo-audit

# Auditoria
cargo audit
cargo audit --json
```

## Multi-linguagem

```bash
# Detectar e auditar automaticamente
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

## Severidade

| Nível | Ação |
|-------|------|
| Critical | Corrigir imediatamente |
| High | Corrigir em 24h |
| Moderate | Corrigir no próximo sprint |
| Low | Avaliar e planejar |

## Tips

- Rode audit em CI/CD para bloquear PRs com vulnerabilidades críticas
- Use `npm audit --production` para ignorar devDependencies
- Mantenha lockfiles atualizados
- Configure Dependabot ou Renovate para PRs automáticos
