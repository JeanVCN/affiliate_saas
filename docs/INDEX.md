---
title: Documentation Index
status: active
owner: technical-writer
last_verified_at: 2026-07-01
source_of_truth: true
---

# Documentation Index

## Meta

- `_meta/documentation-blueprint.md`: pointer to the current master blueprint.
- `_meta/engineering-documentation-blueprint.md`: master documentation and AI-development blueprint.
- `_meta/source-of-truth-map.md`: precedence and document ownership.
- `_meta/documentation-governance.md`: how docs are created, reviewed, split, and archived.

## Product

- `product/vision.md`: product thesis and scope.
- `product/personas.md`: target users and pains.
- `product/roadmap.md`: MVP to future roadmap.

## Architecture

- `architecture/system-overview.md`: system architecture baseline.
- `architecture/context-map.md`: bounded contexts and domain relationships.

## Decisions

- `decisions/decision-log.md`: accepted architecture decisions.
- `decisions/adr/`: architecture decision records.

## Domains

- `domains/README.md`: MVP domain model and first vertical slice.
- `domains/identity/README.md`: users, workspaces, membership, sessions, and authorization.
- `domains/marketplace/README.md`: marketplace and affiliate program setup.
- `domains/product/README.md`: workspace product and offer catalog.
- `domains/affiliate/README.md`: affiliate link registry and variants.
- `domains/link-tracking/README.md`: short links, redirects, and click events.
- `domains/campaign/README.md`: campaign drafts, channel packages, and publishing tasks.
- `domains/analytics/README.md`: click/import aggregates and dashboard insights.
- `domains/compliance/README.md`: affiliate disclosure, claims, policy, and AI output checks.

## API And Database

- `api/README.md`: API conventions and contract overview.
- `api/rest/mvp-endpoints.md`: MVP REST endpoint inventory.
- `database/schema-overview.md`: initial PostgreSQL schema direction.
- `database/entity-catalog.md`: initial entity/table catalog.
- `database/migrations.md`: SQL migration policy.

## Quality, Security, Deployment, Onboarding

- `quality/testing-strategy.md`: testing strategy for backend, database, frontend, and AI.
- `quality/quality-gates.md`: verification gates before merging or presenting work.
- `quality/evals.md`: AI prompt evaluation strategy.
- `security/secrets-policy.md`: secrets and sensitive data policy.
- `security/threat-model.md`: MVP threat model.
- `deployment/docker.md`: Docker-first deployment baseline.
- `onboarding/local-dev.md`: local development commands and workflow.
- `onboarding/agent-onboarding.md`: startup guide for AI agents.

## AI Tooling And Retrieval

- `mcp/mcp-strategy.md`: MCP access strategy and deny-by-default policy.
- `mcp/project-mcp-spec.md`: planned project MCP/tool surface.
- `code-index/indexing-strategy.md`: exact search, symbol search, generated index, and semantic search roles.
- `prompts/prompt-library.md`: AI campaign/compliance prompt families and versioning.

## Context, Memory, Agents

- `context/context-engineering-guide.md`: context loading, retrieval, compression, and budgets.
- `memory/memory-architecture.md`: durable memory model.
- `agents/agent-registry.md`: agent roles.
- `agents/handoff-protocol.md`: how agents transfer work.

## Development Workflows

- `workflows/development/implementation-plan.md`: immediate implementation sequence.
- `workflows/development/documentation-phases.md`: small documentation phases and commit boundaries.
- `workflows/development/documentation-completion-checklist.md`: remaining documentation foundation before feature work.
- `workflows/development/phase-0-readiness.md`: pre-commit readiness checklist for the current foundation.
- `workflows/development/project-resume-brief.md`: exact resume point for a new chat or fresh agent.

## Templates

- `templates/adr.md`
- `templates/handoff.md`
- `templates/session-summary.md`
- `templates/feature.md`
- `templates/task.md`

## Research And Planning Artifacts

Initial research and planning artifacts:

- `product/business-plan-affiliate-multimarketplace.md`: current business plan and product positioning.
- `research/market/tiktok-shop-content-saas-analysis.md`: earlier broader TikTok Shop/content-cutting analysis kept as research history.

Durable content should be promoted from planning artifacts into source-of-truth docs when it becomes operational.
