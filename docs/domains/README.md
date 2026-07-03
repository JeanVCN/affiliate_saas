---
title: MVP Domain Model
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../product/vision.md
  - ../product/roadmap.md
  - ../architecture/context-map.md
---

# MVP Domain Model

This directory defines the business domains for the first Affiliate SaaS vertical slice.

## First Vertical Slice

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```

The slice proves the MVP thesis without depending on closed marketplace APIs:

1. Identity creates a user/workspace boundary.
2. Marketplace records a manually configured affiliate program.
3. Product stores the product/offer the workspace wants to promote.
4. Affiliate stores the user-provided destination URL for that product/program.
5. Link Tracking creates a short link, redirects visitors, and records click events.
6. Analytics aggregates clicks by product/link/campaign context.
7. Campaign uses product and link context for drafts and channel packages.
8. Compliance checks disclosures, unsupported claims, prohibited automation, and AI output risk.

## Domain Docs

- `identity/README.md`
- `marketplace/README.md`
- `product/README.md`
- `affiliate/README.md`
- `link-tracking/README.md`
- `campaign/README.md`
- `analytics/README.md`
- `compliance/README.md`

## Boundary Rule

When domain docs conflict with derived docs, use these domain docs until executable code, schemas, migrations, tests, or accepted ADRs supersede them.
