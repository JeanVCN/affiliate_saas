---
title: Testing Strategy
status: active
owner: qa-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../api/README.md
  - ../database/schema-overview.md
  - ../domains/README.md
---

# Testing Strategy

This strategy defines how future implementation should be verified. The repository currently has documentation only, so executable test commands become active as code is added.

## Test Pyramid

- Unit tests: domain rules, validators, slug generation, URL handling, auth helpers, compliance checks.
- Integration tests: PostgreSQL repositories, migrations, session storage, redirect/click recording, analytics queries.
- HTTP tests: Gin route wiring, request validation, auth middleware, workspace authorization, error shape.
- End-to-end smoke tests: first vertical slice through API and frontend once both exist.
- Prompt/eval tests: AI campaign generation after prompt/provider docs exist.

## Backend

Expected command once backend exists:

```bash
go test ./...
```

Backend coverage priorities:

- Identity: auth, sessions, workspace authorization.
- Marketplace/Product/Affiliate: workspace scoping and relationship integrity.
- Link Tracking: slug uniqueness, redirect behavior, click event creation, recoverable write failure.
- Analytics: correct grouping by product, link, campaign, channel, and time range.
- Compliance: blocker/warning classification and claim/disclosure checks.

## Database

Expected checks once migrations exist:

```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
migrate -path backend/migrations -database "$DATABASE_URL" down 1
```

Database tests should verify:

- fresh database migrates from zero;
- core foreign keys and unique constraints work;
- workspace-owned tables cannot be queried across workspace boundaries by repository methods;
- click and import queries match documented analytics behavior.

## Frontend

Expected command once Next.js exists:

```bash
npm test
```

Frontend testing priorities:

- auth/session states;
- workspace selection;
- product/link/campaign flows;
- user-visible error feedback;
- dashboard empty/loading/error states.

## AI Prompts

AI tests start after prompt/provider docs exist. Minimum expectations:

- prompt templates are versioned;
- generated copy includes affiliate disclosure requirements when relevant;
- outputs do not invent price, availability, guarantee, health, or income claims;
- compliance findings can block risky outputs.

## Manual Smoke Path

Once API exists, the first smoke path is:

1. Create account and workspace.
2. Enable marketplace program manually.
3. Create product and offer.
4. Create affiliate link.
5. Create short link.
6. Visit `/r/{slug}` and confirm redirect.
7. Query analytics clicks by product/link.

## Current Verification

For the documentation-only repository:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```
