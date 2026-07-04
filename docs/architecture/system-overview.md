---
title: System Overview
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
---

# System Overview

## Architecture Style

Start as a modular monolith with clear domain boundaries and async workers. Do not start with microservices. Extract services only when operational load, team ownership, or scaling requirements justify it.

Backend domains live under `backend/internal/modules/*`. Modules are domain-first, not ceremony-first: use handlers, services, repositories, and models when they carry real behavior, but avoid creating extra files only to satisfy symmetry. Route registration should point to named handler methods or functions, and Go files should stay at or below 400 lines; split large handlers, services, or repositories by responsibility before they become catch-all files.

## Preferred Stack

- Backend: Go.
- Database: PostgreSQL.
- Queue: Redis Streams for MVP.
- Storage: S3-compatible object storage.
- Frontend: Next.js + React.
- Deploy: Docker first.

## Accepted Foundational Decisions

- Backend module layout: `docs/decisions/adr/001-backend-module-layout.md`
- HTTP router/framework: `docs/decisions/adr/002-http-router-framework.md`
- Database migration tool: `docs/decisions/adr/003-database-migration-tool.md`
- Auth/session strategy: `docs/decisions/adr/004-auth-session-strategy.md`
- Short-link tracking strategy: `docs/decisions/adr/005-short-link-tracking-strategy.md`
- Token-efficient assisted development: `docs/decisions/adr/006-token-efficient-assisted-development.md`

## Major Modules

- Identity and workspace.
- Marketplace/program catalog.
- Product and offer catalog.
- Affiliate link and click tracking.
- Campaign studio.
- AI generation.
- Compliance checks.
- Publishing tasks.
- Conversion imports.
- Analytics dashboard.
- Billing and usage.

## Initial Runtime

```mermaid
flowchart LR
  User[User] --> Web[Next.js App]
  Web --> API[Go API]
  API --> PG[(PostgreSQL)]
  API --> Redis[(Redis Streams)]
  API --> S3[(S3 Storage)]
  Redis --> Worker[Go Workers]
  Worker --> PG
  Worker --> AI[AI Providers]
  Worker --> Integrations[Marketplace/Social APIs]
```

## Key Decisions Pending ADRs

- AI provider abstraction.
- Redis Streams consumer group conventions.
