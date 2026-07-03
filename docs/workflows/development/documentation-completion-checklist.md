---
title: Documentation Completion Checklist
status: active
owner: technical-writer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../_meta/engineering-documentation-blueprint.md
---

# Documentation Completion Checklist

## Commit Today: Phase 0

- [x] `AGENTS.md`
- [x] Documentation index.
- [x] Source-of-truth map.
- [x] Product vision, personas, and roadmap.
- [x] System overview and context map.
- [x] Context engineering and memory baseline.
- [x] Agent registry and handoff protocol.
- [x] Base templates.
- [x] Planning artifacts moved into `docs/`.
- [x] No application code is introduced in Phase 0.

Commit message recommendation:

```text
docs: establish AI-assisted development foundation
```

## Already Started

- [x] `AGENTS.md`
- [x] `docs/README.md`
- [x] `docs/INDEX.md`
- [x] `docs/_meta/source-of-truth-map.md`
- [x] `docs/_meta/documentation-governance.md`
- [x] `docs/product/vision.md`
- [x] `docs/product/personas.md`
- [x] `docs/product/roadmap.md`
- [x] `docs/architecture/system-overview.md`
- [x] `docs/architecture/context-map.md`
- [x] `docs/context/context-engineering-guide.md`
- [x] `docs/memory/memory-architecture.md`
- [x] `docs/agents/agent-registry.md`
- [x] `docs/agents/handoff-protocol.md`
- [x] base templates

## Phase 1: Architecture Decision Base

- [x] `docs/decisions/decision-log.md`
- [x] First ADRs for module layout, router, migrations, auth, and short links.

## Phase 2: MVP Domain Base

- [x] `docs/domains/affiliate/README.md`
- [x] `docs/domains/marketplace/README.md`
- [x] `docs/domains/product/README.md`
- [x] `docs/domains/campaign/README.md`
- [x] `docs/domains/link-tracking/README.md`
- [x] `docs/domains/analytics/README.md`
- [x] `docs/domains/compliance/README.md`
- [x] `docs/domains/identity/README.md`

## Phase 3: Contracts And Data Base

- [x] `docs/api/README.md`
- [x] `docs/api/rest/mvp-endpoints.md`
- [x] `docs/database/schema-overview.md`
- [x] `docs/database/entity-catalog.md`
- [x] `docs/database/migrations.md`

## Phase 4: Quality Security And Local Dev Base

- [x] `docs/quality/testing-strategy.md`
- [x] `docs/quality/quality-gates.md`
- [x] `docs/security/secrets-policy.md`
- [x] `docs/security/threat-model.md`
- [x] `docs/deployment/docker.md`
- [x] `docs/onboarding/agent-onboarding.md`
- [x] `docs/onboarding/local-dev.md`

## P1 Before Beta

- [x] `docs/mcp/mcp-strategy.md`
- [x] `docs/mcp/project-mcp-spec.md`
- [x] `docs/code-index/indexing-strategy.md`
- [x] `docs/prompts/prompt-library.md`
- [x] `docs/quality/evals.md`
- [x] `docs/quality/quality-gates.md`
- [x] `docs/security/threat-model.md`
- [ ] `docs/product/metrics.md`
- [ ] platform research docs for TikTok Shop, Shopee, Mercado Livre, Amazon.

## Ready For Product Development

- [x] Phase 0 documentation operating base exists.
- [x] Phase 1 ADRs are accepted.
- [x] Phase 2 MVP domains exist.
- [x] Phase 3 API/data baseline exists.
- [x] Phase 4 local verification and secrets policy exist.
- [x] Phase 5 AI/MCP/code-index baseline exists.

## Ready To Start Coding When

- [x] P0 ADRs are accepted.
- [x] Initial domain docs exist for the first vertical slice.
- [x] API/database baseline docs exist.
- [x] Backend module layout ADR is accepted.
- [x] Local verification command is documented.
