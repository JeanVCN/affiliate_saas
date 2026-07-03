---
title: Initial Implementation Plan
status: active
owner: system-architect
last_verified_at: 2026-07-03
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
- Add health endpoint.
- Add config loading.
- Add PostgreSQL connection layer.
- Add Redis Streams abstraction only when first async job appears.

## Phase C: MVP Domain Slice

Do not start this phase until Phase 2 domain docs and Phase 3 API/database baselines exist.

Build the smallest vertical slice:

1. Workspace/user placeholder model.
2. Marketplace program registry.
3. Product registry.
4. Affiliate link registry.
5. Short-link redirect and click event.
6. Campaign draft generated manually first, AI later.
7. Dashboard-ready query for clicks by link/product.

## Phase D: AI And Imports

- Add provider abstraction.
- Add prompt templates for TikTok Shop/Shopee/Mercado Livre/Amazon playbooks.
- Add CSV import contract for conversions/commissions.
- Add compliance checklist data model.

## Phase E: Frontend Foundation

- Create Next.js app.
- Add app shell.
- Add product/link/campaign flows.
- Add dashboard charts after backend contract stabilizes.

## Current Next Action

Start the first implementation slice:

1. Implement repository-backed workspace endpoints.
2. Implement marketplace program setup endpoints.
3. Implement product and offer endpoints.
4. Implement affiliate link and short-link endpoints.
5. Implement `/r/{slug}` click recording and redirect.
6. Implement dashboard query for clicks by product/link.
7. Keep AI generation, marketplace integrations, OAuth, and frontend scaffold out until this slice is stable.
