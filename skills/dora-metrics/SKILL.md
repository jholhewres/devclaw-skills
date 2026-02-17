---
name: dora-metrics
description: "Calculate DORA metrics from git + deploy logs"
---
# DORA Metrics

Use **bash** (git) to calculate DORA engineering metrics.

## Deployment Frequency
```bash
# Tags/releases in last 30 days
git tag --sort=-creatordate --format='%(creatordate:short) %(refname:short)' | head -30
git log --since="30 days ago" --oneline --merges --first-parent main | wc -l
```

## Lead Time for Changes
```bash
# Time from first commit to merge (for recent PRs)
gh pr list -R OWNER/REPO --state merged --limit 20 --json number,createdAt,mergedAt,title
```

## Change Failure Rate
```bash
# Reverts and hotfixes in last 30 days
git log --since="30 days ago" --oneline --grep="revert\|hotfix\|rollback" -i | wc -l
# vs total deploys
git tag --sort=-creatordate --after="30 days ago" | wc -l
```

## Mean Time to Recovery (MTTR)
```bash
# Time between bug report and fix merge
gh issue list -R OWNER/REPO --state closed --label "bug" --search "closed:>=$(date -d '30 days ago' +%Y-%m-%d)" --json number,createdAt,closedAt --limit 20
```

## DORA Performance Levels
| Metric | Elite | High | Medium | Low |
|--------|-------|------|--------|-----|
| Deploy Freq | On-demand | Weekly | Monthly | Monthly+ |
| Lead Time | < 1 day | 1 week | 1 month | 6 months |
| Change Fail | 0-15% | 16-30% | 16-30% | 46-60% |
| MTTR | < 1 hour | < 1 day | < 1 week | 6 months |

## Tips
- Track over time using memory_save for historical data
- Compare across sprints for improvement trends
- Use git tags consistently for accurate deployment frequency
