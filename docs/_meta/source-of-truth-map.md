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
| System architecture | `docs/architecture/system-overview.md` | Superseded only by ADR |
| Domain boundaries | `docs/architecture/context-map.md` | Domain docs will refine this |
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
