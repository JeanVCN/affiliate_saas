---
title: Context Engineering Guide
status: active
owner: context-engineering-specialist
last_verified_at: 2026-07-01
source_of_truth: true
---

# Context Engineering Guide

## Goal

Give agents the smallest context that preserves correctness.

## Loading Order

1. Read `AGENTS.md`.
2. Read `docs/INDEX.md`.
3. Read the source-of-truth doc for the task area.
4. Use `rg` or symbol search for exact files.
5. Open focused snippets, not whole directories.
6. Compress findings before implementation.

## Context Layers

| Layer | Example | Load when |
|---|---|---|
| Global | `AGENTS.md` | Always |
| Index | `docs/INDEX.md` | Any non-trivial task |
| Product | `docs/product/*` | Product/feature work |
| Architecture | `docs/architecture/*`, ADRs | Cross-module work |
| Domain | `docs/domains/*` | Business logic |
| Module | service/API/database docs | Implementation |
| Task | issue/spec/handoff | Always |
| Evidence | logs/tests/diffs | Verification |

## Budget Heuristic

- Keep always-on context under 1,200 tokens.
- Prefer 3-5 precise files over 20 broad reads.
- Summarize before switching agents.
- Put source links in summaries so the next agent can reload evidence.

## Anti-Patterns

- Reading the whole repo.
- Duplicating product and architecture text in every prompt.
- Using memory as a source of truth.
- Relying on semantic search when an exact symbol or route is known.

