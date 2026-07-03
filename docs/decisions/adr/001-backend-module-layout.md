---
title: ADR-001 Backend Module Layout
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# ADR-001: Backend Module Layout

Status: accepted

Date: 2026-07-03

## Context

Affiliate SaaS needs a backend that can support identity, marketplace programs, product catalogs, affiliate links, campaign generation, click tracking, imports, analytics, and billing without turning the MVP into a distributed system too early.

The product roadmap favors a fast, inspectable MVP that can run in Docker with PostgreSQL and, later, Redis Streams. The architecture overview already points to a modular monolith, so the repository layout should make module boundaries explicit while keeping local development and deployment simple.

## Decision

Use a Go modular monolith under `backend/`.

Initial layout:

```text
backend/
  cmd/
    api/
    worker/
  internal/
    app/
    config/
    http/
    modules/
      identity/
      marketplace/
      product/
      affiliate/
      linktracking/
      campaign/
      analytics/
      compliance/
      billing/
    platform/
      postgres/
      redis/
      storage/
      ai/
  migrations/
  tests/
```

Boundary rules:

- `cmd/api` and `cmd/worker` are thin composition roots.
- `internal/modules/*` owns domain behavior and application use cases.
- `internal/platform/*` owns adapters for PostgreSQL, Redis, object storage, AI providers, and external services.
- Cross-module calls should go through explicit service interfaces or application use cases, not shared global state.
- Database migrations live in `backend/migrations`.

## Alternatives Considered

- Single flat Go package: faster for the first endpoint, but it would make domain boundaries fuzzy exactly where this product needs clarity.
- Multiple Go services: aligns with future scale, but adds deployment, observability, auth, and data consistency costs before the MVP proves demand.
- Clean Architecture with many layers per module from day one: rigorous, but likely too ceremony-heavy before the domain model stabilizes.

## Consequences

Positive:

- Keeps MVP deployment simple while preserving extraction paths.
- Gives agents and humans predictable places to put code.
- Makes domain ownership visible before the codebase grows.
- Supports API and worker binaries without duplicating core logic.

Negative:

- Requires discipline to avoid modules reaching across boundaries casually.
- Some shared concerns may feel verbose until enough behavior exists.
- Service extraction later will still require explicit data and contract work.

## Verification

- New backend code lands under the documented layout.
- Domain code does not import web framework types directly unless it is explicitly HTTP delivery code.
- `cmd/api` and `cmd/worker` remain composition entrypoints, not business logic containers.
- Future domain docs map cleanly to `internal/modules/*`.

## Links

- Related docs:
  - `docs/architecture/system-overview.md`
  - `docs/architecture/context-map.md`
  - `docs/product/roadmap.md`
- Related issues:
  - None yet.
