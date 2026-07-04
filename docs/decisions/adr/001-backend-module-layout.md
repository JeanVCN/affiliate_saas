---
title: ADR-001 Backend Module Layout
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
---

# ADR-001: Backend Module Layout

Status: accepted

Date: 2026-07-03

## Context

Affiliate SaaS needs a backend that can support identity, marketplace programs, product catalogs, affiliate links, campaign generation, click tracking, imports, analytics, and billing without turning the MVP into a distributed system too early.

The product roadmap favors a fast, inspectable first release that can run in Docker with PostgreSQL and, later, Redis Streams. The architecture overview already points to a modular monolith, so the repository layout should make module boundaries explicit while keeping local development and deployment simple.

The codebase should not be shaped as throwaway MVP scaffolding. It should start with an architecture that can evolve naturally as domains gain behavior, without requiring a later folder-level rewrite just to introduce handlers, services, and repositories.

## Decision

Use a Go modular monolith under `backend/`.

Initial layout:

```text
backend/
  cmd/
    api/
    worker/
  internal/
    config/
    http/
    modules/
      identity/
        handler.go
        service.go
        repository.go
        models.go
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
- `internal/http` owns process-level HTTP setup, middleware, health checks, and route composition.
- `internal/modules/*` owns domain HTTP handlers, services, repositories, models, and application behavior.
- `handler.go` is the only module layer that should import Gin or HTTP request/response types.
- Module handlers should be methods on a `Handler` type. `RegisterRoutes` should only instantiate the handler and map routes to named methods; it should not contain inline route lambdas except for trivial temporary experiments.
- `service.go` owns use-case orchestration, validation, and cross-repository coordination inside the module.
- `repository.go` owns persistence contracts and PostgreSQL implementations for that module's tables and read models.
- `internal/platform/*` owns shared infrastructure setup such as opening PostgreSQL pools, Redis clients, object storage clients, AI clients, and external service clients.
- Cross-module calls should go through explicit service interfaces or application use cases, not shared global state or direct table coupling in handlers.
- Database migrations live in `backend/migrations`.
- Keep Go files at or below 400 lines. When a module's `repository.go`, `service.go`, or `handler.go` approaches that limit or starts mixing unrelated responsibilities, split it by responsibility, for example `product_repository.go`, `offer_repository.go`, `click_repository.go`, or `dashboard_service.go`.

## Alternatives Considered

- Single flat Go package: faster for the first endpoint, but it would make domain boundaries fuzzy exactly where this product needs clarity.
- Multiple Go services: aligns with future scale, but adds deployment, observability, auth, and data consistency costs before the MVP proves demand.
- A single cross-domain `mvp` application/store package: fast for the first vertical slice, but it would make the initial implementation look disposable and would require restructuring exactly when the product starts to gain real behavior.
- Full Clean Architecture with many nested packages per module from day one: rigorous, but likely too ceremony-heavy before the domain model stabilizes.

## Consequences

Positive:

- Keeps MVP deployment simple while preserving extraction paths.
- Gives agents and humans predictable places to put code.
- Makes domain ownership visible before the codebase grows.
- Lets the first product slice become production code instead of temporary MVP code.
- Supports API and worker binaries without duplicating core logic.

Negative:

- Requires discipline to avoid modules reaching across boundaries casually.
- Some handler/service/repository separation may feel verbose while modules are still small.
- Service extraction later will still require explicit data and contract work.

## Verification

- New backend code lands under the documented layout.
- Domain code imports web framework types only from module `handler.go` files or process-level HTTP setup.
- `RegisterRoutes` functions map URLs to named handler methods instead of embedding business logic in anonymous functions.
- No regular Go source file exceeds 400 lines without an intentional split plan.
- `cmd/api` and `cmd/worker` remain composition entrypoints, not business logic containers.
- Domain docs map cleanly to `internal/modules/*`.

## Links

- Related docs:
  - `docs/architecture/system-overview.md`
  - `docs/architecture/context-map.md`
  - `docs/product/roadmap.md`
- Related issues:
  - None yet.
