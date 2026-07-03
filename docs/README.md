---
title: Documentation Home
status: active
owner: project
last_verified_at: 2026-07-01
source_of_truth: true
---

# Documentation Home

This directory is the development operating system for the Affiliate SaaS project. It is designed for long-running work by humans and AI agents with low token waste, clear sources of truth, and documented handoffs.

Start here:

- `INDEX.md` maps the documentation tree.
- `_meta/source-of-truth-map.md` defines which documents own each kind of decision.
- `product/vision.md` defines what the product is.
- `architecture/system-overview.md` defines the initial technical direction.
- `decisions/decision-log.md` records accepted architecture decisions.
- `domains/*/README.md` defines MVP domain rules and boundaries.
- `api/README.md` and `database/schema-overview.md` define initial contracts before implementation.
- `quality/quality-gates.md`, `security/secrets-policy.md`, and `onboarding/local-dev.md` define verification and safety baselines.
- `mcp/mcp-strategy.md`, `code-index/indexing-strategy.md`, and `prompts/prompt-library.md` define AI tooling and retrieval baselines.
- `context/context-engineering-guide.md` defines how agents should retrieve and compress context.
- `agents/agent-registry.md` defines agent responsibilities.

Rules:

- Update docs with code when contracts, architecture, domain rules, prompts, deployment, or operations change.
- Prefer concise source-of-truth docs plus links over repeated explanations.
- Use templates from `templates/` for ADRs, handoffs, feature specs, and session summaries.
