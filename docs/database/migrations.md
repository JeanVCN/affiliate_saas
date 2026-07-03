---
title: Database Migrations
status: active
owner: database-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../decisions/adr/003-database-migration-tool.md
  - schema-overview.md
---

# Database Migrations

Affiliate SaaS uses SQL-first PostgreSQL migrations managed by `golang-migrate/migrate`.

## Location

Migrations live in:

```text
backend/migrations/
```

## Naming

Use ordered filenames:

```text
000001_create_identity.up.sql
000001_create_identity.down.sql
000002_create_marketplace_product_links.up.sql
000002_create_marketplace_product_links.down.sql
```

## Rules

- Every merged migration is append-only.
- Use paired `up` and `down` migrations while pre-production.
- Keep migrations deterministic and reviewable.
- Prefer explicit constraints over application-only integrity.
- Do not generate migrations from an ORM.
- Do not run migrations from API process startup.

## Expected Commands

Exact commands will be finalized when backend tooling exists. The intended shape is:

```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
migrate -path backend/migrations -database "$DATABASE_URL" down 1
```

## First Migration Plan

Split initial migrations by dependency order:

1. Identity: users, workspaces, memberships, sessions.
2. Marketplace and product: marketplaces, programs, workspace programs, products, offers.
3. Affiliate and tracking: affiliate links, link variants, short links, click events.
4. Campaign and compliance: campaigns, channel packages, publishing tasks, compliance checks/findings.
5. Imports: conversion imports and rows.

## Verification

- A fresh database migrates from zero to latest.
- Down migration works for the latest pre-production migration.
- Backend tests can create a migrated test database once code exists.
- Schema docs are updated when migrations add or change entities.
