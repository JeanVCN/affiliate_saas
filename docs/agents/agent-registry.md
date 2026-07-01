---
title: Agent Registry
status: active
owner: system-architect
last_verified_at: 2026-07-01
source_of_truth: true
---

# Agent Registry

| Agent | Responsibility | Reads | Updates | Done when |
|---|---|---|---|---|
| System Architect | Architecture, ADRs, boundaries | architecture, decisions, product | ADRs, architecture | decision and impact are documented |
| Backend Engineer | Go API, workers, integrations | domains, services, api, database | service/API docs | tests/build pass or technical plan is ready |
| Frontend Engineer | Next.js app and UX workflows | product, api, frontend docs | frontend docs | UI flow is implemented or specified |
| AI Engineer | Prompts, providers, evals | ai architecture, prompts, quality | prompts/evals | prompt behavior is versioned/testable |
| Database Engineer | Schema and migrations | database, domains | database docs | schema impact is documented |
| QA Engineer | Test strategy and gates | quality, product, domains | quality docs | acceptance checks are clear |
| Security Engineer | Threat model and policy risks | security, architecture | security docs | risks and mitigations are listed |
| Researcher | Official docs and market evidence | research index | research docs | sources and uncertainty are documented |
| Product Manager | Scope, personas, roadmap | product, research | product docs | success criteria are clear |
| Technical Writer | Docs structure and consistency | _meta, templates | docs/indexes | docs are linked and non-duplicative |

## Routing Rule

Use the smallest capable role. Add a second agent for review when work affects architecture, security, compliance, billing, or data integrity.

