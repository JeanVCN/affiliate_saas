---
title: Documentation Phases
status: active
owner: technical-writer
last_verified_at: 2026-07-01
source_of_truth: true
depends_on:
  - ../../_meta/engineering-documentation-blueprint.md
  - documentation-completion-checklist.md
---

# Documentation Phases

This plan divides the documentation foundation into small commits. The goal is to keep momentum without starting product implementation before the project has a stable operating base.

## Phase 0: Commit Today - Documentation Operating Base

Goal: make the repository navigable and safe for future AI-assisted work.

Scope:

- Always-on agent rules.
- Documentation index.
- Source-of-truth map.
- Product vision, personas, and roadmap.
- System overview and context map.
- Context engineering and memory baseline.
- Agent registry and handoff protocol.
- Base templates.
- Planning artifacts moved into `docs/`.

Out of scope:

- Technical ADR decisions.
- Domain modeling detail.
- API contracts.
- Database schema.
- Product implementation.

Definition of Done:

- `docs/INDEX.md` points to all current docs.
- `.docs-index.json` and `.context-manifest.json` are valid JSON.
- Root planning artifacts are organized under `docs/`.
- No application code or backend module is introduced in this phase.
- `git status` shows only intentional files.

Suggested commit:

```text
docs: establish AI-assisted development foundation
```

## Phase 1: Architecture Decision Base

Goal: make the first technical choices explicit before coding.

Scope:

- `docs/decisions/decision-log.md`.
- ADR for backend module layout.
- ADR for HTTP router/framework.
- ADR for database migration tool.
- ADR for auth/session direction.
- ADR for short-link tracking strategy.

Definition of Done:

- Each ADR has context, decision, alternatives, consequences, and verification.
- Decision log links all ADRs.
- `docs/architecture/system-overview.md` references accepted decisions.

Suggested commit:

```text
docs: record foundational architecture decisions
```

## Phase 2: MVP Domain Base

Goal: document the business model of the first vertical slice.

Scope:

- `docs/domains/affiliate/README.md`
- `docs/domains/marketplace/README.md`
- `docs/domains/campaign/README.md`
- `docs/domains/link-tracking/README.md`
- `docs/domains/analytics/README.md`
- `docs/domains/compliance/README.md`
- `docs/domains/identity/README.md`

Definition of Done:

- Each domain defines purpose, entities, invariants, lifecycle, risks, and open questions.
- Domains link to product roadmap and architecture context map.
- First vertical slice is explicitly described end-to-end.

Suggested commit:

```text
docs: define MVP domain model
```

## Phase 3: Contracts And Data Base

Goal: document API and database direction before implementation.

Scope:

- `docs/api/README.md`
- `docs/api/rest/mvp-endpoints.md`
- `docs/database/schema-overview.md`
- `docs/database/entity-catalog.md`
- `docs/database/migrations.md`

Definition of Done:

- MVP endpoints are listed with purpose and domain owner.
- Initial entities are mapped to domains.
- Migration policy is defined.
- No implementation is blocked by unknown contracts.

Suggested commit:

```text
docs: define initial API and data contracts
```

## Phase 4: Quality, Security, And Local Dev Base

Goal: make future coding verifiable and safe.

Scope:

- `docs/quality/testing-strategy.md`
- `docs/quality/quality-gates.md`
- `docs/security/secrets-policy.md`
- `docs/security/threat-model.md`
- `docs/deployment/docker.md`
- `docs/onboarding/local-dev.md`
- `docs/onboarding/agent-onboarding.md`

Definition of Done:

- Local commands are documented.
- Test strategy covers backend, frontend future, database, and AI prompts.
- Secrets policy is clear.
- Threat model covers affiliate links, OAuth, marketplace reports, AI outputs, and tracking data.

Suggested commit:

```text
docs: add quality security and local development base
```

## Phase 5: AI, MCP, And Code Index Base

Goal: prepare advanced agent productivity after core docs are stable.

Scope:

- `docs/mcp/mcp-strategy.md`
- `docs/mcp/project-mcp-spec.md`
- `docs/code-index/indexing-strategy.md`
- `docs/prompts/prompt-library.md`
- `docs/quality/evals.md`

Definition of Done:

- MCP allowlist and security model are documented.
- Project MCP tool list is defined but not necessarily implemented.
- Code indexing strategy defines exact search, symbol search, and semantic search roles.
- Prompt library has versioning and eval expectations.

Suggested commit:

```text
docs: define AI tooling and context retrieval strategy
```

## Phase 6: Ready For Product Development

Start implementation only when:

- Phase 0 is committed.
- Phase 1 ADRs are accepted.
- Phase 2 domains exist for the first vertical slice.
- Phase 3 API/data baseline exists.
- Phase 4 local verification and secrets policy exist.

The first implementation slice should be:

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```
