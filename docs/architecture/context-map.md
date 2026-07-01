---
title: Context Map
status: draft
owner: system-architect
last_verified_at: 2026-07-01
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

