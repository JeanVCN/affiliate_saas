---
title: Architecture Decision Log
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# Architecture Decision Log

This log tracks accepted architecture decisions for Affiliate SaaS. Accepted ADRs supersede broader architecture docs when they conflict.

| ADR | Status | Date | Decision |
|---|---|---:|---|
| [ADR-001](adr/001-backend-module-layout.md) | accepted | 2026-07-03 | Use a Go modular monolith under `backend/` with domain modules in `internal/modules`. |
| [ADR-002](adr/002-http-router-framework.md) | accepted | 2026-07-03 | Use Gin as the Go HTTP framework/router for the MVP API. |
| [ADR-003](adr/003-database-migration-tool.md) | accepted | 2026-07-03 | Use SQL-first migrations managed by `golang-migrate/migrate`. |
| [ADR-004](adr/004-auth-session-strategy.md) | accepted | 2026-07-03 | Use first-party email/password auth with secure cookie sessions for the MVP. |
| [ADR-005](adr/005-short-link-tracking-strategy.md) | accepted | 2026-07-03 | Use first-party short links and click events handled by the Go API. |

## Status Meanings

- `proposed`: under review and not yet binding.
- `accepted`: binding source of truth until superseded.
- `superseded`: replaced by a newer ADR.
- `rejected`: considered and explicitly not chosen.
