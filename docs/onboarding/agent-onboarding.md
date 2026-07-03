---
title: Agent Onboarding
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../AGENTS.md
  - ../INDEX.md
  - ../workflows/development/project-resume-brief.md
---

# Agent Onboarding

Use this guide when an AI agent starts work in this repository.

## Always Load

1. `AGENTS.md`
2. `docs/INDEX.md`
3. `docs/_meta/source-of-truth-map.md`
4. `docs/workflows/development/project-resume-brief.md`

## Load By Task

- Architecture: `docs/architecture/`, `docs/decisions/`.
- Domain behavior: `docs/domains/`.
- API work: `docs/api/`, relevant domains.
- Database work: `docs/database/`, relevant domains.
- Security/auth/tracking: `docs/security/`, `docs/domains/identity/`, `docs/domains/link-tracking/`.
- Local/dev workflow: `docs/onboarding/local-dev.md`, `docs/deployment/docker.md`, `docs/quality/quality-gates.md`.

## Operating Rules

- Explore before editing.
- Keep changes scoped.
- Update docs when architecture, domain behavior, API contracts, database schema, prompts, deployment, or workflows change.
- Create or update an ADR for architecturally significant decisions.
- Keep `gin.Context` out of domain/application modules.
- Avoid scraping, browser automation, artificial engagement, and copyright-risk workflows.
- Do not commit secrets, credentials, customer reports, or private exports.

## Current Verification

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

When backend/frontend code exists, also run the relevant quality gates in `docs/quality/quality-gates.md`.
