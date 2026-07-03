---
title: Threat Model
status: active
owner: security-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - secrets-policy.md
  - ../domains/README.md
  - ../api/README.md
---

# Threat Model

This threat model covers the MVP surface before implementation.

## Assets

- User accounts and password hashes.
- Workspace business data.
- Affiliate links and destination URLs.
- Click events and tracking metadata.
- Conversion imports and commission data.
- Campaign drafts and AI-generated commercial copy.
- Future marketplace/OAuth credentials.

## Trust Boundaries

- Browser to Go API.
- Public redirect route `/r/{slug}` to backend.
- Backend to PostgreSQL.
- Backend/worker to Redis Streams when added.
- Backend/worker to AI providers when added.
- User-provided CSV/import data to analytics.

## Key Threats And Mitigations

| Threat | Risk | Mitigation |
|---|---|---|
| Workspace data leak | Users access another workspace's data. | Server-side membership checks on every workspace-scoped endpoint. |
| Open redirect abuse | Short links redirect to malicious destinations. | Validate URL scheme/host rules and allow pausing/archiving links. |
| Click fraud or bot traffic | Analytics becomes misleading. | Record enough metadata for later bot heuristics; label metrics honestly. |
| Secret exposure | Credentials leak from Git/logs. | Secrets policy, `.env.example`, no real secrets in docs or repo. |
| Session theft | Account takeover. | HttpOnly/Secure/SameSite cookies, server-side revocation, password hashing. |
| CSV/import poisoning | Malformed imports corrupt analytics. | Validate file/row shape, preserve raw rows, mark unmatched rows clearly. |
| AI claim fabrication | Risky commercial copy. | Compliance checks, product-grounded prompts, user review before publishing. |
| Marketplace policy drift | Workflows violate platform rules. | Dated policy notes and no scraping/browser automation in MVP. |

## MVP Non-Goals

- No automated social engagement.
- No scraping marketplaces or social networks.
- No marketplace OAuth storage until integrations are approved.
- No public API tokens.
- No automated publishing through unofficial routes.

## Review Triggers

Update this document when:

- auth/session behavior changes;
- short-link tracking stores new request metadata;
- CSV imports are implemented;
- AI generation is implemented;
- official marketplace integrations are added;
- production deployment begins.
