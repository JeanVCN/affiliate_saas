---
title: Project Resume Brief
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - documentation-phases.md
  - phase-0-readiness.md
  - ../../_meta/source-of-truth-map.md
---

# Project Resume Brief

Use this file when resuming the project in a new chat or with a fresh agent.

## Current State

Phase 0 is complete. Phase 1 foundational ADRs are accepted. Phase 2 MVP domain docs are complete. Phase 3 API/database baseline docs are complete. Phase 4 quality/security/local-dev docs are complete. Phase 5 AI/MCP/code-index docs are complete.

The repository now contains documentation plus a modular Go/Gin backend. The first backend slice has repository-backed modules for auth/session, workspace setup, marketplace/program setup, products/offers, affiliate links, short-link redirects, click recording, and click analytics. Auth includes signup, login, logout, `me`, Argon2id password hashing, HttpOnly session cookies, workspace RBAC, and OAuth readiness tables without provider token storage. There is intentionally no frontend app or AI/product automation yet.

## Product Direction

Build an AI-assisted affiliate commerce operating system, not a generic video clipping product.

Initial focus:

- TikTok Shop Affiliate.
- Shopee Affiliates.
- Mercado Livre.
- Amazon Associates/Creators.

MVP thesis:

```text
product -> affiliate link -> AI campaign -> channel package -> click tracking -> conversion import -> insight
```

## Where To Start

Always load:

1. `AGENTS.md`
2. `docs/INDEX.md`
3. `docs/workflows/development/documentation-phases.md`
4. `docs/workflows/development/phase-0-readiness.md`
5. `docs/workflows/development/documentation-completion-checklist.md`

For product context, load:

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/product/business-plan-affiliate-multimarketplace.md` only if deeper business context is needed.

For architecture context, load:

- `docs/architecture/system-overview.md`
- `docs/architecture/context-map.md`

## Current Architecture Decisions

- `docs/decisions/decision-log.md`
- `docs/decisions/adr/001-backend-module-layout.md`
- `docs/decisions/adr/002-http-router-framework.md`
- `docs/decisions/adr/003-database-migration-tool.md`
- `docs/decisions/adr/004-auth-session-strategy.md`
- `docs/decisions/adr/005-short-link-tracking-strategy.md`

## Current Domain Docs

- `docs/domains/README.md`
- `docs/domains/identity/README.md`
- `docs/domains/marketplace/README.md`
- `docs/domains/product/README.md`
- `docs/domains/affiliate/README.md`
- `docs/domains/link-tracking/README.md`
- `docs/domains/campaign/README.md`
- `docs/domains/analytics/README.md`
- `docs/domains/compliance/README.md`

## Current Contract Docs

- `docs/api/README.md`
- `docs/api/rest/mvp-endpoints.md`
- `docs/database/schema-overview.md`
- `docs/database/entity-catalog.md`
- `docs/database/migrations.md`

## Current Quality And Safety Docs

- `docs/quality/testing-strategy.md`
- `docs/quality/quality-gates.md`
- `docs/quality/evals.md`
- `docs/security/secrets-policy.md`
- `docs/security/threat-model.md`
- `docs/deployment/docker.md`
- `docs/onboarding/local-dev.md`
- `docs/onboarding/agent-onboarding.md`

## Current AI Tooling Docs

- `docs/mcp/mcp-strategy.md`
- `docs/mcp/project-mcp-spec.md`
- `docs/code-index/indexing-strategy.md`
- `docs/prompts/prompt-library.md`

## Next Phase

The next phase is continuing backend product features after auth hardening.

Continue with:

- PostgreSQL-backed validation for the authenticated first vertical slice:

```text
signup -> session cookie -> workspace RBAC -> marketplace program -> product -> offer -> affiliate link -> short redirect -> click event -> analytics query
```

Run:

```bash
cd backend
GOCACHE=/tmp/affiliate-saas-go-cache go test ./...
AFFILIATE_TEST_DATABASE_URL='postgres://affiliate:affiliate@localhost:55432/affiliate_saas?sslmode=disable' GOCACHE=/tmp/affiliate-saas-go-cache go test ./tests/integration
```

## Keep Out Of First Slice

- AI generation.
- Marketplace API integrations.
- OAuth provider token storage.
- Browser automation.
- Frontend scaffold unless explicitly requested.

## Suggested Next Commit

Commit the auth hardening slice as:

```text
feat: add secure auth sessions and workspace rbac
```

## Validation For Phase 0

Run from repository root:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

Expected result: command exits with code 0.
