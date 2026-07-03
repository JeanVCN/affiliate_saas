---
title: MCP Strategy
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../security/secrets-policy.md
  - ../security/threat-model.md
  - ../context/context-engineering-guide.md
---

# MCP Strategy

MCP access for Affiliate SaaS is deny-by-default. Tools should improve retrieval, verification, and local productivity without exposing secrets, customer data, marketplace credentials, or private reports.

## Goals

- Give agents precise project context without loading the whole repository.
- Support safe local development and verification.
- Keep tool permissions narrow, auditable, and revocable.
- Avoid connecting to external marketplaces or customer systems before explicit approval.

## Default Policy

- No write access unless a task explicitly requires it.
- No access to secrets or credential stores.
- No marketplace account access in the MVP.
- No browser automation for posting, engagement, login, scraping, or marketplace workflows.
- No customer report ingestion through MCP without a documented import contract and user approval.

## Allowed Tool Categories

| Category | Default | Purpose |
|---|---|---|
| Repository file read | allowed | Load docs, source files, migrations, tests. |
| Repository file write | task-scoped | Edit project files when implementing an approved task. |
| Shell commands | task-scoped | Run documented verification commands. |
| Git status/diff | allowed | Inspect worktree state and produce summaries. |
| Local database | later, read/write in dev only | Run migrations and integration tests against disposable data. |
| Code index/search | allowed | Resolve symbols, routes, docs, and ownership quickly. |
| External web | explicit approval | Research official docs or current platform policies. |

## Disallowed By Default

- Production database access.
- Secret manager access.
- Marketplace credentials or OAuth tokens.
- Customer CSV/report folders.
- Browser automation against third-party platforms.
- Social engagement automation.
- Sending unpublished customer data to external AI tools.

## Approval Triggers

Require explicit approval before enabling a tool that can:

- access network resources;
- write outside the repository;
- access non-local databases;
- read private reports or exports;
- use paid AI/provider APIs;
- authenticate to marketplaces, social networks, or commerce platforms.

## Audit Expectations

Any future project MCP spec must document:

- tool name;
- purpose;
- read/write scope;
- data sensitivity;
- approval requirement;
- verification command or expected output;
- owner.

## Links

- `docs/mcp/project-mcp-spec.md`
- `docs/security/secrets-policy.md`
- `docs/security/threat-model.md`
- `docs/context/context-engineering-guide.md`
