---
title: Initial Implementation Plan
status: active
owner: system-architect
last_verified_at: 2026-07-06
source_of_truth: true
depends_on:
  - ../../product/roadmap.md
  - ../../architecture/system-overview.md
  - ../../_meta/source-of-truth-map.md
---

# Initial Implementation Plan

This plan sequences work after the documentation foundation has started. Product implementation must not start until the documentation phases define the relevant decisions, domain rules, contracts, and verification strategy.

## Phase A: Finish Documentation Architecture

- Create initial ADRs for foundational technical choices. Done in `docs/decisions/`.
- Create domain docs for identity, marketplace, product, affiliate, link tracking, campaign, analytics, and compliance. Done in `docs/domains/`.
- Create API and database overview docs before implementing contracts. Done in `docs/api/` and `docs/database/`.
- Create quality, security, and deployment baseline docs. Done in `docs/quality/`, `docs/security/`, and `docs/deployment/`.
- Create onboarding docs for humans and agents. Done in `docs/onboarding/`.

## Phase B: Backend Foundation

Do not start this phase until Phase 1 ADRs are accepted.

- Choose Go HTTP router/framework via ADR. Done in `docs/decisions/adr/002-http-router-framework.md`.
- Choose migration tool via ADR. Done in `docs/decisions/adr/003-database-migration-tool.md`.
- Define project layout under `backend/`. Done in `docs/decisions/adr/001-backend-module-layout.md`.
- Add health endpoint. Done in `backend/internal/http`.
- Add config loading. Done in `backend/internal/config`.
- Add PostgreSQL connection layer. Done in `backend/internal/platform/postgres`.
- Add Redis Streams abstraction only when first async job appears.

## Phase C: MVP Domain Slice

Do not start this phase until Phase 2 domain docs and Phase 3 API/database baselines exist.

Build the smallest vertical slice:

1. Workspace/user placeholder model. Implemented for workspace endpoints.
2. Marketplace program registry. Implemented for manual marketplace/program setup.
3. Product registry. Implemented for products and offers.
4. Affiliate link registry. Implemented for destination links.
5. Short-link redirect and click event. Implemented for `/r/{slug}`.
6. Auth/session hardening. Implemented with signup, login, logout, me, Argon2id, session cookies, and workspace RBAC.
7. Campaign draft generated manually first, AI later. Implemented for manual drafts and channel packages.
8. Dashboard-ready queries. Implemented for click metrics, overview, and top products.
9. Manual conversion import batches and rows. Implemented for API-backed manual imports.
10. Compliance checklist data model and campaign check execution. Implemented for basic MVP rules.
11. Manual publishing tasks. Implemented for campaign task creation, listing, scheduling, and completion.
12. CSV conversion import parsing. Implemented for CSV text payloads into import rows.
13. Program policy notes for compliance context. Implemented as manual notes attached to campaign checks.

## Phase D: AI And Imports

- Add provider abstraction.
- Add prompt templates for TikTok Shop/Shopee/Mercado Livre/Amazon playbooks.
- Add richer import mapping rules after manual reconciliation is stable. Started with deterministic product mapping from known affiliate links.
- Expand compliance rules after manual policy notes and review workflow are stable.

## Phase E: Frontend Foundation

- Create Next.js app. Implemented in `frontend/`.
- Add app shell. Implemented with a two-column operational workspace.
- Add auth/session workflow. Implemented for signup, login, session restore, and logout.
- Add product/link starter flows. Implemented for marketplace program, product, offer, affiliate link, short link, and analytics overview.
- Add campaign/import/compliance flows after the first frontend foundation is committed.
- Add dashboard charts after analytics contracts and frontend layout stabilize.

## Current Next Action

Stabilize and commit the first frontend slice before AI work:

1. Run frontend checks: `cd frontend && npm run lint && npm run build`.
2. Keep backend checks available before committing or when API contracts change:

```bash
cd backend
GOCACHE=/tmp/affiliate-saas-go-cache go test ./...
AFFILIATE_TEST_DATABASE_URL='postgres://affiliate:affiliate@localhost:55432/affiliate_saas?sslmode=disable' GOCACHE=/tmp/affiliate-saas-go-cache go test ./tests/integration
```

3. Verify the frontend first slice through the Next proxy:

```text
signup/login -> session restore -> marketplace program -> product -> offer -> affiliate link -> short link -> analytics overview
```

4. Keep AI generation, marketplace integrations, provider OAuth token storage, scraping, and browser automation out until the manual frontend MVP is more complete.
5. Next candidates after this frontend checkpoint: campaign/compliance/import screens, richer analytics UI, or AI/provider abstraction once approved.
