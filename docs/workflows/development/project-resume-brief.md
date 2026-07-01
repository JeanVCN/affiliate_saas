---
title: Project Resume Brief
status: active
owner: system-architect
last_verified_at: 2026-07-01
source_of_truth: true
depends_on:
  - documentation-phases.md
  - phase-0-readiness.md
  - ../../_meta/source-of-truth-map.md
---

# Project Resume Brief

Use this file when resuming the project in a new chat or with a fresh agent.

## Current State

Phase 0 is complete and ready to commit.

The repository currently contains documentation and planning only. There is intentionally no backend module, frontend app, database schema, or product implementation yet.

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

## Next Phase

The next phase is Phase 1: Architecture Decision Base.

Create:

- `docs/decisions/decision-log.md`
- ADR for backend module layout.
- ADR for HTTP router/framework.
- ADR for database migration tool.
- ADR for auth/session direction.
- ADR for short-link tracking strategy.

## Do Not Start Yet

Do not implement backend, frontend, database migrations, API endpoints, AI prompts, or infrastructure before Phase 1 ADRs are accepted and Phase 2/3 docs define the first vertical slice.

## Suggested Next Commit

Commit current Phase 0 as:

```text
docs: establish AI-assisted development foundation
```

## Validation For Phase 0

Run from repository root:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

Expected result: command exits with code 0.

