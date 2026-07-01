---
title: Handoff Protocol
status: active
owner: system-architect
last_verified_at: 2026-07-01
source_of_truth: true
---

# Handoff Protocol

Use a handoff when another agent or future session must continue work without reading the full conversation.

## Required Fields

```text
handoff_id:
from_agent:
to_agent:
task:
current_state:
source_of_truth_read:
files_or_docs_changed:
decisions_made:
open_questions:
risks:
verification_done:
next_actions:
context_to_load_next:
do_not_repeat:
```

## Rules

- Keep handoff to 1-2 pages.
- Link sources instead of copying large excerpts.
- List exact files changed.
- Include verification evidence or state what was not verified.
- If a decision is architectural, create or reference an ADR.

