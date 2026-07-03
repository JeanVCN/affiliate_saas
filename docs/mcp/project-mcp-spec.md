---
title: Project MCP Spec
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - mcp-strategy.md
  - ../onboarding/agent-onboarding.md
---

# Project MCP Spec

This document defines the desired project-specific MCP/tool surface. It is a specification, not a statement that these tools are already implemented.

## Tool Inventory

| Tool | Status | Access | Purpose |
|---|---|---|---|
| `repo.search` | planned | read-only | Search files with exact text/symbol queries. |
| `repo.read` | planned | read-only | Read focused file snippets. |
| `repo.diff` | planned | read-only | Summarize current worktree changes. |
| `docs.lookup` | planned | read-only | Resolve source-of-truth docs by concern. |
| `quality.verify_docs` | planned | read-only command | Parse `.docs-index.json` and `.context-manifest.json`. |
| `backend.test` | future | task-scoped command | Run `go test ./...` once backend exists. |
| `db.migrate_dev` | future | local dev database only | Run migrations against disposable local database. |
| `frontend.test` | future | task-scoped command | Run frontend tests/build once frontend exists. |

## Required Constraints

- Tools must not expose `.env`, `.env.local`, secrets, credentials, tokens, or private reports.
- Tools must be scoped to the repository or disposable local services.
- Tools must not authenticate to marketplace/social platforms.
- Tools that write files or run commands require task-specific intent.
- Tools should return concise results with file paths and line references where possible.

## Current Verification Tool

The current documentation-only verification command is:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

## Future Local Database Tool

When backend and migrations exist, a local database tool may support:

- create disposable dev/test database;
- run migrations;
- reset disposable data;
- run integration tests.

It must not connect to production or shared customer databases.

## Future Code Index Tool

When source code exists, a code index tool may expose:

- route inventory;
- Go package/module map;
- database migration/entity map;
- frontend route/component map;
- test coverage inventory.

It should use repository files as source of truth and avoid stale generated claims.
