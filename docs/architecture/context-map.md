---
title: Context Map
status: draft
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# Context Map

```mermaid
flowchart TD
  Identity[Identity/Workspace]
  Marketplace[Marketplace Programs]
  Product[Products/Offers]
  Link[Affiliate Links]
  Campaign[Campaigns]
  AI[AI Generation]
  Compliance[Compliance]
  Publishing[Publishing Tasks]
  Analytics[Analytics]
  Billing[Billing/Usage]

  Identity --> Marketplace
  Marketplace --> Product
  Product --> Link
  Product --> Campaign
  Campaign --> AI
  Campaign --> Compliance
  Campaign --> Publishing
  Link --> Analytics
  Analytics --> Campaign
  AI --> Billing
  Publishing --> Analytics
```

## Boundary Rules

- Marketplace integrations normalize external programs into internal products, offers, links, and reports.
- Campaign generation must not own product truth.
- Analytics reads click and conversion events; it does not mutate campaign content.
- Compliance checks annotate and block risky outputs; they do not generate copy.
- Billing consumes usage events; it does not call AI providers directly.

## Domain Docs

- Identity/Workspace: `docs/domains/identity/README.md`
- Marketplace Programs: `docs/domains/marketplace/README.md`
- Products/Offers: `docs/domains/product/README.md`
- Affiliate Links: `docs/domains/affiliate/README.md`
- Link Tracking: `docs/domains/link-tracking/README.md`
- Campaigns: `docs/domains/campaign/README.md`
- Analytics: `docs/domains/analytics/README.md`
- Compliance: `docs/domains/compliance/README.md`
