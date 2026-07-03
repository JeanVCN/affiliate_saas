---
title: ADR-003 Database Migration Tool
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# ADR-003: Database Migration Tool

Status: accepted

Date: 2026-07-03

## Context

The MVP depends on PostgreSQL for users, workspaces, marketplace programs, products, affiliate links, click events, imported conversions, campaigns, and dashboard queries.

Migrations must be reviewable, deterministic, CI-friendly, and usable from Docker without requiring application startup.

## Decision

Use SQL-first migrations managed by `golang-migrate/migrate`.

Guidelines:

- Store migrations in `backend/migrations`.
- Use paired `up` and `down` SQL files while the product is pre-production.
- Prefer explicit SQL over ORM-generated migrations.
- Run migrations as a separate local/CI/deploy step.
- Treat migration filenames as append-only once merged.

## Alternatives Considered

- Goose: strong Go-native option with simple SQL migrations, but `golang-migrate/migrate` has broader CLI/deployment familiarity and explicit multi-driver support.
- Atlas: powerful schema management, but more capability than needed for an MVP baseline.
- ORM-managed migrations: convenient initially, but risks hiding SQL details that matter for analytics, tracking, and data integrity.
- Manual SQL with no migration tool: simple until environments diverge; too risky for even a small SaaS.

## Consequences

Positive:

- Migrations are plain SQL and easy to review.
- Works well in CI, Docker, and production deploy scripts.
- Keeps schema ownership separate from application boot.
- Avoids committing early to an ORM.

Negative:

- Developers must keep SQL files ordered and reversible.
- No automatic model-to-schema sync.
- Down migrations may become best-effort after production hardening.

## Verification

- A fresh database can be migrated from zero using only committed migration files.
- CI can run migrations before backend tests once the backend exists.
- Schema docs in `docs/database` reference migration policy and entity ownership.

## Links

- Related docs:
  - `docs/workflows/development/implementation-plan.md`
  - `docs/architecture/system-overview.md`
- Related issues:
  - None yet.
