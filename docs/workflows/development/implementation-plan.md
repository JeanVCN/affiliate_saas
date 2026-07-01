---
title: Initial Implementation Plan
status: active
owner: system-architect
last_verified_at: 2026-07-01
source_of_truth: true
depends_on:
  - ../../product/roadmap.md
  - ../../architecture/system-overview.md
  - ../../_meta/source-of-truth-map.md
---

# Initial Implementation Plan

This plan sequences work after the documentation foundation has started. Product implementation must not start until the documentation phases define the relevant decisions, domain rules, contracts, and verification strategy.

## Phase A: Finish Documentation Architecture

- Create initial ADRs for foundational technical choices.
- Create domain docs for affiliate, marketplace, campaign, link tracking, analytics, AI generation, compliance, identity, and billing.
- Create API and database overview docs before implementing contracts.
- Create quality, security, and deployment baseline docs.
- Create onboarding docs for humans and agents.

## Phase B: Backend Foundation

Do not start this phase until Phase 1 ADRs are accepted.

- Choose Go HTTP router/framework via ADR.
- Choose migration tool via ADR.
- Define project layout under `backend/`.
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

Create the P0 ADRs:

1. Go backend module layout.
2. HTTP router/framework.
3. Database migration tool.
4. Auth/session strategy.
5. Short link tracking strategy.
