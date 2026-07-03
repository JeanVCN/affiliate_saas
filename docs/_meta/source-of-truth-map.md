---
title: Source Of Truth Map
status: active
owner: system-architect
last_verified_at: 2026-07-01
source_of_truth: true
---

# Source Of Truth Map

When documents conflict, use this precedence:

1. Code, tests, generated contracts, migrations.
2. Accepted ADRs.
3. Source-of-truth docs below.
4. Derived docs.
5. Memories and session summaries.
6. Old chat context.

| Concern | Source of truth | Notes |
|---|---|---|
| Product thesis | `docs/product/vision.md` | Supported by `docs/product/business-plan-affiliate-multimarketplace.md` |
| Personas | `docs/product/personas.md` | Drives UX, prompts, onboarding |
| Roadmap | `docs/product/roadmap.md` | Use for sequencing, not implementation detail |
| Architecture decisions | `docs/decisions/decision-log.md` and `docs/decisions/adr/` | Accepted ADRs supersede broader architecture docs |
| System architecture | `docs/architecture/system-overview.md` | Superseded only by ADR |
| Domain boundaries | `docs/architecture/context-map.md` | Domain docs will refine this |
| Domain rules | `docs/domains/*/README.md` | Source of truth for MVP business invariants and lifecycle |
| API contracts | `docs/api/README.md` and `docs/api/rest/mvp-endpoints.md` | Superseded by generated contracts once code exists |
| Database schema | `docs/database/schema-overview.md`, `docs/database/entity-catalog.md`, migrations | Migrations become highest source of truth once implemented |
| Quality gates | `docs/quality/quality-gates.md` | Source of truth for verification expectations |
| Testing strategy | `docs/quality/testing-strategy.md` | Source of truth for test coverage direction |
| Secrets and threat model | `docs/security/secrets-policy.md`, `docs/security/threat-model.md` | Source of truth for MVP safety constraints |
| Local development | `docs/onboarding/local-dev.md` and `docs/deployment/docker.md` | Source of truth until executable scripts exist |
| MCP/tool access | `docs/mcp/mcp-strategy.md`, `docs/mcp/project-mcp-spec.md` | Deny-by-default until approved |
| Code indexing | `docs/code-index/indexing-strategy.md` | Generated indexes never supersede code or migrations |
| Prompt library and evals | `docs/prompts/prompt-library.md`, `docs/quality/evals.md` | Source of truth for AI generation governance |
| Agent rules | `AGENTS.md` | Always-on, short |
| Context engineering | `docs/context/context-engineering-guide.md` | Controls token strategy |
| AI memory | `docs/memory/memory-architecture.md` | Durable memory policy |
| Agent responsibilities | `docs/agents/agent-registry.md` | Multi-agent routing |
| Handoff | `docs/agents/handoff-protocol.md` | Required for multi-agent transfer |
| Documentation architecture | `docs/_meta/engineering-documentation-blueprint.md` | Master blueprint |

## Update Policy

- If code changes domain behavior, update the domain source of truth.
- If architecture changes, create or update an ADR before editing derived docs.
- If a source-of-truth doc gets too large, split by domain or concern and update this map.
