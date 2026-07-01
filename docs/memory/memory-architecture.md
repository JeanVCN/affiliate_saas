---
title: Memory Architecture
status: active
owner: context-engineering-specialist
last_verified_at: 2026-07-01
source_of_truth: true
---

# Memory Architecture

## Memory Layers

- Conversation context: temporary and not source of truth.
- Session summary: compact record of completed work.
- Agent memory: role-specific recurring learnings.
- Domain memory: durable business/domain facts.
- Decision memory: index of accepted ADRs.
- Source-of-truth docs: final durable knowledge.

## Rules

- Durable decisions must be documented in source-of-truth docs or ADRs.
- Memory cannot contain secrets or private tokens.
- Memory entries must identify scope, source, confidence, and freshness.
- Stale memory must be marked stale, not silently reused.

## Session Summary Minimum

```text
date:
objective:
files_changed:
docs_changed:
decisions:
verification:
open_questions:
next_context_to_load:
```

