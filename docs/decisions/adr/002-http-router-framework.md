---
title: ADR-002 HTTP Router Framework
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# ADR-002: HTTP Router Framework

Status: accepted

Date: 2026-07-03

## Context

The MVP API needs conventional REST endpoints for auth, workspaces, marketplace programs, products, affiliate links, campaigns, click redirects, imports, and dashboard queries.

The project should keep backend development fast, familiar, and easy to onboard. The team prefers Gin for the initial Go API because it offers a common routing/middleware model, practical request binding, and a large ecosystem for SaaS-style REST APIs.

## Decision

Use Gin as the Go HTTP framework/router for the MVP API.

Guidelines:

- Keep route registration close to module delivery code.
- Keep Gin types at the HTTP delivery boundary; domain and application code must not depend on `gin.Context`.
- Use middleware for request IDs, logging, recovery, auth, rate limiting, and CORS.
- Keep OpenAPI or endpoint documentation in `docs/api` before or alongside implementation.
- Do not place domain rules in HTTP handlers.
- Use explicit request/response DTOs instead of passing framework context through module services.

## Alternatives Considered

- Standard library `http.ServeMux`: minimal and increasingly capable, but route grouping and middleware composition remain less ergonomic for a multi-module SaaS API.
- `chi`: small, idiomatic, and compatible with `net/http`, but the team preference is Gin for faster onboarding and a more batteries-included REST API surface.
- Echo or Fiber: polished developer ergonomics, but more framework-specific surface area than the MVP needs.

## Consequences

Positive:

- Familiar framework for many Go backend developers.
- Productive routing, middleware, binding, and JSON response ergonomics.
- Large ecosystem and plenty of operational examples.
- Fits a modular monolith when Gin is kept at the delivery boundary.

Negative:

- Larger framework surface than `chi` or `http.ServeMux`.
- `gin.Context` can leak into application/domain code if boundaries are not enforced.
- Handler conventions, validation, and error response shape still need project-level choices.

## Verification

- API handlers can be tested through Gin's HTTP engine and `httptest`.
- Module route registration does not import unrelated modules directly.
- No `internal/modules/*` package imports Gin.
- Error response, auth, and request validation conventions are documented before broad endpoint implementation.

## Revision Notes

- 2026-07-03: Updated accepted router/framework from `chi` to Gin by project direction.

## Links

- Related docs:
  - `docs/architecture/system-overview.md`
  - `docs/workflows/development/implementation-plan.md`
- Related issues:
  - None yet.
