---
title: Documentation Completion Checklist
status: active
owner: technical-writer
last_verified_at: 2026-07-01
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

- [ ] `docs/decisions/decision-log.md`
- [ ] First ADRs for module layout, router, migrations, auth, and short links.

## Phase 2: MVP Domain Base

- [ ] `docs/domains/affiliate/README.md`
- [ ] `docs/domains/marketplace/README.md`
- [ ] `docs/domains/campaign/README.md`
- [ ] `docs/domains/link-tracking/README.md`
- [ ] `docs/domains/analytics/README.md`
- [ ] `docs/domains/compliance/README.md`
- [ ] `docs/domains/identity/README.md`

## Phase 3: Contracts And Data Base

- [ ] `docs/api/README.md`
- [ ] `docs/database/schema-overview.md`

## Phase 4: Quality Security And Local Dev Base

- [ ] `docs/quality/testing-strategy.md`
- [ ] `docs/security/secrets-policy.md`
- [ ] `docs/deployment/docker.md`
- [ ] `docs/onboarding/agent-onboarding.md`
- [ ] `docs/onboarding/local-dev.md`

## P1 Before Beta

- [ ] `docs/mcp/mcp-strategy.md`
- [ ] `docs/code-index/indexing-strategy.md`
- [ ] `docs/quality/quality-gates.md`
- [ ] `docs/security/threat-model.md`
- [ ] `docs/product/metrics.md`
- [ ] platform research docs for TikTok Shop, Shopee, Mercado Livre, Amazon.

## Ready To Start Coding When

- [ ] P0 ADRs are accepted.
- [ ] Initial domain docs exist for the first vertical slice.
- [ ] API/database baseline docs exist.
- [ ] Backend module layout ADR is accepted.
- [ ] Local verification command is documented.
