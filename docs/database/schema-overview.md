---
title: Database Schema Overview
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../domains/README.md
  - ../decisions/adr/003-database-migration-tool.md
---

# Database Schema Overview

Affiliate SaaS starts on PostgreSQL with SQL-first migrations.

## Goals

- Support the first vertical slice without marketplace API dependency.
- Preserve workspace boundaries in every business table.
- Keep analytics queryable from normalized click/link/product/campaign relationships.
- Avoid premature service-specific databases.

## Initial Entity Groups

- Identity: users, workspaces, memberships, sessions.
- Marketplace: marketplace catalog, program catalog, workspace programs.
- Product: products, offers, product angles.
- Affiliate: affiliate links, link variants.
- Link Tracking: short links, click events.
- Campaign: campaigns, channel packages, publishing tasks.
- Compliance: compliance checks and findings.
- Analytics: conversion import batches, conversion import rows, aggregate query views later.

## Workspace Ownership

Every workspace-owned table must include `workspace_id` unless it is:

- a global catalog table such as `marketplaces` or `programs`;
- a pure join table whose parents already enforce workspace ownership;
- an auth/session table tied directly to users.

## Common Columns

Use these where practical:

- `id`
- `workspace_id` for workspace-owned records
- `created_at`
- `updated_at`
- `archived_at` for archiveable user-owned records
- `status` for lifecycle state

## Deletion Policy

- Prefer archive/soft-delete for user-owned business records.
- Hard-delete sessions when revoked or expired according to retention rules.
- Do not hard-delete click events or import rows until retention policy exists.

## Data Integrity

- Use foreign keys for core relationships.
- Use unique constraints for slugs, emails, and workspace-specific names where needed.
- Use check constraints for small enums where it improves safety.
- Keep generated analytics summaries derivable from raw events/imports during MVP.

## Next Step

The first migration should create identity, marketplace, product, affiliate link, short-link, click event, campaign draft, compliance finding, and conversion import tables needed by `docs/api/rest/mvp-endpoints.md`.
