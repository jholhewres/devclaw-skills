---
name: meeting-notes
description: "Structured meeting notes with action items extraction"
---
# Meeting Notes

Use **write_file**, **read_file**, **edit_file**, and **memory_save** for meeting notes.

## Create Meeting Note
Use **write_file** to create `~/.devclaw/meetings/YYYY-MM-DD-<title>.md`:

```markdown
# Meeting: <Title>
**Date:** YYYY-MM-DD HH:MM
**Attendees:** Name1, Name2, Name3
**Type:** standup | planning | retro | 1:1 | brainstorm | review

## Agenda
1. ...
2. ...

## Notes
- ...

## Decisions
- [ ] Decision 1

## Action Items
- [ ] @Name1 — Task description (due: YYYY-MM-DD)
- [ ] @Name2 — Task description (due: YYYY-MM-DD)

## Follow-up
- Next meeting: YYYY-MM-DD
```

## Search Meetings
```bash
ls -lt ~/.devclaw/meetings/ | head -20
grep -rl "TERM" ~/.devclaw/meetings/ --include="*.md"
grep -rl "@PersonName" ~/.devclaw/meetings/ --include="*.md"
```

## Extract Action Items
```bash
grep -h "\- \[ \]" ~/.devclaw/meetings/*.md | sort
grep -h "\- \[ \] @Name" ~/.devclaw/meetings/*.md   # by person
```

## Tips
- Create dir on first use: mkdir -p ~/.devclaw/meetings
- Save key decisions with memory_save for long-term recall
- Extract action items and offer to create reminders (cron_add)
- When user pastes raw meeting notes, restructure into template
